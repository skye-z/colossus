package backend

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/skye-z/colossus/backend/model"
	"github.com/skye-z/colossus/backend/service"
	"golang.org/x/crypto/ssh"
	"xorm.io/xorm"
)

// 连接升级程序
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 缓存池
type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

// 写入缓存
func (w *wsBufferWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

// 获取缓存
func (w *wsBufferWriter) Bytes() []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Bytes()
}

// 刷新缓存
func (w *wsBufferWriter) Reset() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.buffer.Reset()
}

// WebSocket服务
type SocketService struct {
	// 数据库引擎
	DB *xorm.Engine
	// 会话读写管道
	stdinPipe io.WriteCloser
	// 连接缓存
	comboOutput *wsBufferWriter
	// SSH会话
	session *ssh.Session
	// WebSocket连接
	wsConn *websocket.Conn
}

// 启动WebSocket服务
func (s *SocketService) Run(ctx *gin.Context) {
	// 获取主机编号
	queryId := ctx.DefaultQuery("id", "")
	if queryId == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	hostId, _ := strconv.ParseInt(queryId, 10, 64)
	// 获取主机信息
	hostModel := model.HostModel{DB: s.DB}
	host := &model.Host{Id: hostId}
	hostModel.GetItem(host)
	if len(host.Address) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	// 升级连接
	upgrade, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	defer upgrade.Close()
	// 组装连接配置
	config := service.SSHService{
		Address:  host.Address,
		Port:     host.Port,
		AuthType: host.AuthType,
		User:     host.User,
		Secret:   host.Secret,
	}
	// 创建客户端
	client, err2 := config.CreateClient()
	if err2 != nil {
		upgrade.WriteMessage(websocket.TextMessage, []byte("\n[1] 创建客户端失败"+err2.Error()))
		return
	}
	// 获取宽高
	cols, _ := strconv.Atoi(ctx.DefaultQuery("cols", "80"))
	rows, _ := strconv.Atoi(ctx.DefaultQuery("rows", "90"))
	// 连接主机
	session, err3 := config.Connect(client, cols, rows)
	if err3 != nil {
		upgrade.WriteMessage(websocket.TextMessage, []byte("\n[2] 创建会话失败"+err3.Error()))
		return
	}
	pipe, _ := session.StdinPipe()
	wsBuffer := new(wsBufferWriter)
	session.Stdout = wsBuffer
	session.Stderr = wsBuffer
	// 启动终端
	err = session.Shell()
	if err != nil {
		upgrade.WriteMessage(websocket.TextMessage, []byte("第三步:启动shell终端失败"+err.Error()))
		return
	}
	// 暂存连接信息
	var connect = &SocketService{
		stdinPipe:   pipe,
		comboOutput: wsBuffer,
		session:     session,
		wsConn:      upgrade,
	}
	// 转入协程
	quitChan := make(chan bool, 3)
	// 启动连接
	connect.start(ctx, quitChan)
	// 等待响应
	go connect.Wait(quitChan)
	<-quitChan
}

// 开始传输数据
func (s *SocketService) start(ctx *gin.Context, quitChan chan bool) {
	// 接收前端传入
	go s.receiveWsMsg(ctx, quitChan)
	// 发送后端输出
	go s.sendWsOutput(quitChan)
}

// 接受消息
func (s *SocketService) receiveWsMsg(ctx *gin.Context, quitChan chan bool) {
	wsConn := s.wsConn
	defer setQuit(quitChan)
	for {
		select {
		case <-quitChan:
			return
		default:
			_, data, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("连接断开")
				return
			}
			if data[0] == 33 && data[1] == 126 {
				cmd := string(data)
				rc := strings.Split(cmd[2:], ":")
				// 获取宽高
				cols, _ := strconv.Atoi(rc[0])
				rows, _ := strconv.Atoi(rc[1])
				s.session.WindowChange(rows, cols)
			} else {
				s.stdinPipe.Write(data)
			}
		}
	}
}

// 发送消息
func (s *SocketService) sendWsOutput(quitChan chan bool) {
	wsConn := s.wsConn
	defer setQuit(quitChan)
	ticker := time.NewTicker(time.Millisecond * time.Duration(60))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if s.comboOutput == nil {
				return
			}
			bytes := s.comboOutput.Bytes()
			if len(bytes) > 0 {
				wsConn.WriteMessage(websocket.TextMessage, bytes)
				s.comboOutput.buffer.Reset()
			}
		case <-quitChan:
			return
		}

	}
}

// 等待关闭
func (s *SocketService) Wait(quitChan chan bool) {
	s.session.Wait()
	log.Println("关闭连接")
	setQuit(quitChan)
}

// 关闭连接
func setQuit(quitChan chan bool) {
	quitChan <- true
}
