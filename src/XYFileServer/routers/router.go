package routers

import (
	"XYAPIServer/XYFileServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/file",
			beego.NSRouter("/view_file",&controllers.FileViewController{}),			
		),
	)
	beego.AddNamespace(ns)
}
