package main

import (
	_ "b0go/apps/docs"
	_ "b0go/apps/pass"
	"net/http"
	"os"
	"strings"
	"time"

	"b0go/core/engine"
	_ "b0go/core/gateway"

	"github.com/pelletier/go-toml"
)

var (
	configIni = "config.ini"
	configTxt = "[gateway]\nLive = false\nListenAddr = \"$ip$\"\n\n[pass]\nLive = false\nPath = \"$path$\""
)

type configData struct {
	Ip   interface{}
	Path interface{}
}

// 启动HTTP服务
func StartServer() {
	engine.Run("config.ini")
	select {}
}

// Load Config
func (a *App) LoadConfig() configData {
	config, err := toml.LoadFile(configIni)
	if err != nil {
		//configTxt := "[gateway]\nLive = false\nListenAddr = \":8899\"\n\n[pass]\nLive = false\nPath = \"files\""
		configTxt = strings.Replace(configTxt, "$ip$", ":8899", -1)
		configTxt = strings.Replace(configTxt, "$path$", "files", -1)
		os.WriteFile(configIni, []byte(configTxt), 0755)
		config, _ = toml.LoadFile(configIni)
	}
	ListenAddr := config.Get("gateway.ListenAddr")
	Path := config.Get("pass.Path")
	//res, _ := json.Marshal(configData{Ip: ListenAddr, Path: Path})
	return configData{Ip: ListenAddr, Path: Path}
}

// Submit Config
func (a *App) SubmitConfig(ip, path string) string {
	var urls string
	if !strings.Contains(ip, ":") {
		ip = ":" + ip
		urls = "http://127.0.0.1" + ip
	}
	path = strings.ReplaceAll(path, `\`, `\\`)
	configTxt = strings.Replace(configTxt, "$ip$", ip, -1)
	configTxt = strings.Replace(configTxt, "$path$", path, -1)
	//configTxt := "[gateway]\nLive = false\nListenAddr = \"" + ip + "\"\n\n[pass]\nLive = false\nPath = \"" + path + "\""
	os.WriteFile(configIni, []byte(configTxt), 0755)

	// Behind Server
	time.Sleep(1000)
	go func() {
		StartServer()
	}()
	time.Sleep(10000)

	// Check Server
	if a.CheckServer(urls) {
		return "OK"
	}
	return "ERR"
}

// CheckServer 检测服务是否启动成功
func (a *App) CheckServer(url string) bool {
	_, err := http.Get(url)
	return err != nil
}
