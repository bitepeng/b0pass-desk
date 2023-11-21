package main

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
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

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// 在窗口执行js代码
func (a *App) ExecJs(str string) {
	runtime.WindowExecJS(a.ctx, str)
}

// 选择目录对话框
func (a *App) OpenDirectoryDialog() string {
	res, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{DefaultDirectory: "./"})
	return res
}

// 选择文件对话框
func (a *App) OpenFileDialog() string {
	res, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{DefaultDirectory: "./"})
	return res
}

// 选择多个文件对话框
func (a *App) OpenMultipleFilesDialog() []string {
	res, _ := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{DefaultDirectory: "./"})
	return res
}

// 保存文件对话框
func (a *App) SaveFileDialog() string {
	res, _ := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{DefaultDirectory: "./"})
	return res
}

// 消息提示对话框
func (a *App) MessageDialog() string {
	res, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Question",
		Message:       "Do you want to continue?",
		DefaultButton: "No",
	})
	return res
}

// 添加菜单
func (a *App) AddMenu() {
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {

	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(a.ctx)
	})

	if os.Getenv("GOOS") == "darwin" {
		AppMenu.Append(menu.EditMenu())
	}
	runtime.MenuSetApplicationMenu(a.ctx, AppMenu)
}
