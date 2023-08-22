package main

import (
	"context"

	local "runtime"

	"github.com/skye-z/colossus/backend"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogPrint(ctx, "colossus startup")
	backend.Start()
}

func (a *App) GetOSName() string {
	return local.GOOS
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

	return dialog != "确认" && dialog != "是" && dialog != "Yes"
}
