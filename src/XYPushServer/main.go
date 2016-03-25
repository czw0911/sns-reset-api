package main

import (
	_ "XYAPIServer/XYPushServer/docs"
	_ "XYAPIServer/XYPushServer/routers"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	XYLibs.PageErrorSet()
	beego.Run()
}
