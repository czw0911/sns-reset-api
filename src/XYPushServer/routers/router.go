package routers

import (
	"XYAPIServer/XYPushServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/push",
			beego.NSRouter("/ios",&controllers.IOSController{}),	
	)
	beego.AddNamespace(ns)
}
