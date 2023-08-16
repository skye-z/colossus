package backend

import (
	"github.com/gin-gonic/gin"
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
	route := gin.Default()

	socket := SocketService{}

	route.GET("/ws", socket.Run)

	return route
}
