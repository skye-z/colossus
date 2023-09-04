package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/colossus/backend/model"
	"github.com/skye-z/colossus/backend/service"
)

func Start() (ok bool) {
	// 关闭调试
	gin.SetMode(gin.ReleaseMode)
	// 注册路由
	route := register()
	// 启动服务
	route.Run(":18703")
	return true
}

func register() *gin.Engine {
	// 获取数据库引擎
	engine := GetDBEngine()
	// 创建默认路由
	route := gin.Default()
	// 加载跨域服务
	route.Use(Cors())

	// 创建Socket服务
	socket := SocketService{DB: engine}
	// 挂载Socket服务
	route.GET("/ws", socket.Run)

	// 创建配置服务
	configService := service.ConfigService{}
	// 接口 获取所有配置
	route.GET("/config/all", configService.GetAll)
	// 接口 更新配置
	route.POST("/config/update", configService.Update)

	// 创建主机模型
	hostModel := model.HostModel{DB: engine}
	// 创建主机服务
	hostService := service.HostService{HostModel: hostModel}
	// 接口 添加主机
	route.POST("/host/add", hostService.Add)
	// 接口 编辑主机
	route.POST("/host/:id", hostService.Edit)
	// 接口 删除主机
	route.DELETE("/host/:id", hostService.Del)
	// 接口 获取主机列表
	route.GET("/host/list", hostService.GetList)
	// 接口 获取主机详情
	route.GET("/host/:id", hostService.GetItem)

	// 创建文件服务
	fileService := service.FileService{
		HostModel: hostModel,
	}
	// 接口 查询文件列表
	route.POST("/file/list", fileService.GetFileList)
	// 接口 查询文件详情
	route.POST("/file/info", fileService.GetFileInfo)
	// 接口 下载文件
	route.POST("/file/down", fileService.DownloadFile)
	// 接口 上传文件
	route.POST("/file/up", fileService.UploadFile)
	// 接口 重命名文件
	route.POST("/file/move", fileService.MoveFile)

	return route
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Credentials", "false")
		}
		c.Header("Access-Control-Allow-Headers", "content-type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,DELETE")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
