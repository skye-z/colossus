package main

import (
	"context"
	"fmt"

	"github.com/skye-z/colossus/backend"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogPrint(ctx, "colossus startup")
	backend.Start()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "关闭 Colossus",
		Message:       "确认要终止正在进行的操作并关闭 Colossus 吗?",
		Buttons:       []string{"取消", "确认"},
		DefaultButton: "取消",
		CancelButton:  "取消",
	})

	if err != nil {
		return false
	}
	return dialog != "确认"
}
