package main

import (
	_ "XYAPIServer/XYRobot/docs"
	_ "XYAPIServer/XYRobot/routers"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
)

func main() {
	XYLibs.PageErrorSet()
	beego.Run()
}
