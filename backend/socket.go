package backend

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

// 连接升级程序
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

func (w *wsBufferWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

func (w *wsBufferWriter) Bytes() []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Bytes()
}

func (w *wsBufferWriter) Reset() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.buffer.Reset()
}

type SocketService struct {
	stdinPipe   io.WriteCloser
	comboOutput *wsBufferWriter
	session     *ssh.Session
	wsConn      *websocket.Conn
}

func (s *SocketService) Run(context *gin.Context) {
	upgrade, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		context.AbortWithStatus(http.StatusOK)
		fmt.Println("http升级websocket失败")
		return
	}
	defer upgrade.Close()

	config := SSHService{
		Address:  "192.168.1.2",
		Port:     22,
		AuthType: AUTH_TYPE_PASSWORD,
		User:     "root",
		Secret:   "PKUsz123",
	}
	client, err2 := config.CreateClient()
	if err2 != nil {
		upgrade.WriteMessage(websocket.TextMessage, []byte("\n[1] 创建客户端失败"+err2.Error()))
		return
	}

	session, err3 := config.Connect(client)
	if err3 != nil {
		upgrade.WriteMessage(websocket.TextMessage, []byte("\n[2] 创建会话失败"+err3.Error()))
		return
	}

	pipe, _ := session.StdinPipe()
	wsBuffer := new(wsBufferWriter)
	session.Stdout = wsBuffer
	session.Stderr = wsBuffer

	err = session.Shell()
	if err != nil {
		upgrade.WriteMessage(websocket.TextMessage, []byte("第三步:启动shell终端失败"+err.Error()))
		return
	}
	var connect = &SocketService{
		stdinPipe:   pipe,
		comboOutput: wsBuffer,
		session:     session,
		wsConn:      upgrade,
	}

	quitChan := make(chan bool, 3)
	connect.start(quitChan)
	go connect.Wait(quitChan)
	<-quitChan
}

func (s *SocketService) start(quitChan chan bool) {
	go s.receiveWsMsg(quitChan)
	go s.sendWsOutput(quitChan)
}

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

func (s *SocketService) receiveWsMsg(quitChan chan bool) {
	wsConn := s.wsConn
	defer setQuit(quitChan)
	for {
		select {
		case <-quitChan:
			return
		default:
			_, data, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Println("receiveWsMsg=>读取ws信息失败", err)
				return
			}
			s.stdinPipe.Write(data)
		}
	}
}

func (s *SocketService) Wait(quitChan chan bool) {
	if err := s.session.Wait(); err != nil {
		setQuit(quitChan)
	}
}

func setQuit(quitChan chan bool) {
	quitChan <- true
}
