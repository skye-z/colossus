package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/colossus/backend/model"
	"github.com/skye-z/colossus/backend/service"
)

// 启动服务
func Start() (ok bool) {
	// 关闭调试
	gin.SetMode(gin.ReleaseMode)
	// 注册路由
	route := register()
	// 启动服务
	route.Run(":18703")
	return true
}

// 注册路由
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
	// 接口 查询用户目录
	route.POST("/file/home", fileService.GetHomePath)
	// 接口 查询文件列表
	route.POST("/file/list", fileService.GetFileList)
	// 接口 查询文件详情
	route.POST("/file/info", fileService.GetFileInfo)
	// 接口 下载文件
	route.POST("/file/down", fileService.DownloadFile)
	// 接口 上传文件
	route.POST("/file/up", fileService.UploadFile)
	// 接口 移动文件
	route.POST("/file/move", fileService.MoveFile)
	// 接口 删除文件
	route.POST("/file/remove", fileService.RemoveFile)
	// 接口 创建目录
	route.POST("/file/directory", fileService.CreateDirectory)

	// 创建命令模型
	codeModel := model.CodeModel{DB: engine}
	// 创建命令服务
	codeService := service.CodeService{CodeModel: codeModel}
	// 接口 获取命令列表
	route.GET("/code/list", codeService.GetList)
	// 接口 添加命令
	route.POST("/code/add", codeService.Add)
	// 接口 编辑命令
	route.POST("/code/:id", codeService.Edit)
	// 接口 删除命令
	route.DELETE("/code/:id", codeService.Del)

	// 创建凭证模型
	certModel := model.CertModel{DB: engine}
	// 创建凭证服务
	certService := service.CertService{CertModel: certModel}
	// 接口 获取凭证列表
	route.GET("/cert/list", certService.GetList)
	// 接口 获取简单凭证列表
	route.GET("/cert/list/all", certService.GetAllList)
	// 接口 获取初始化数据
	route.GET("/cert/add", certService.InitAdd)
	// 接口 添加凭证
	route.POST("/cert/add", certService.Add)
	// 接口 编辑凭证
	route.POST("/cert/:id", certService.Edit)
	// 接口 删除凭证
	route.DELETE("/cert/:id", certService.Del)

	// 接口 执行命令
	route.POST("/host/run", fileService.RunCMD)

	return route
}

// 配置跨域
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
