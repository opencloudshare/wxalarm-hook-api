package main

import (
	_ "wxapiserver/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logs.Async()
	logs.SetLogger(logs.AdapterFile, `{"filename":"/root/go/src/wxapiserver/logs/GrafanaReq.log","level":7}`)
	beego.Run()
}
