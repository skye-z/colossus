package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// 引入静态资源
// go:embed all:frontend/dist
var assets embed.FS

func main() {
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
		// 静态资源服务
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// 背景颜色
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
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
		println("Error:", err.Error())
	}
}
