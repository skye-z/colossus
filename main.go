package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/skye-z/colossus/backend/common"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// 加载日志文件地址
	cacheDir, _ := os.UserCacheDir()
	logPath := fmt.Sprintf("%s/%s", cacheDir, "colossus.log")
	fmt.Println("Log path: " + logPath)
	fileLogger := NewFileLogger(logPath)
	// 初始化配置文件
	common.InitConfig()
	// 创建应用程序实例
	app := NewApp()
	// 运行应用程序
	err := wails.Run(&options.App{
		// 窗口标题
		Title: "Colossus",
		// 默认宽度
		Width: 1024,
		// 默认高度
		Height: 700,
		// 最小宽度
		MinWidth: 1024,
		// 最小高度
		MinHeight: 700,
		// 日志记录器
		Logger: fileLogger,
		// 开发日志级别
		LogLevel: logger.DEBUG,
		// 生产日志级别
		LogLevelProduction: logger.ERROR,
		// 静态资源服务
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// 背景颜色
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Windows: &windows.Options{
			DisableWindowIcon: true,
			Theme:             windows.Dark,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			TitleBar:             mac.TitleBarHiddenInset(),
			Appearance:           mac.NSAppearanceNameDarkAqua,
			About: &mac.AboutInfo{
				Title:   "Colossus",
				Message: "© 2023 Skye Zhang",
				Icon:    icon,
			},
		},
		// 应用启动事件
		OnStartup:     app.startup,
		OnBeforeClose: app.beforeClose,
		// 调试
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fileLogger.Print("Error:" + err.Error())
	}
}
