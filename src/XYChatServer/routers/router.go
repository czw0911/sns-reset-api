package routers

import (
	"XYAPIServer/XYChatServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/chat",
			beego.NSRouter("/get_rc_token",&controllers.TokenController{}),
			beego.NSRouter("/get_sht_token",&controllers.SHTController{}),
		),
		
	)
	beego.AddNamespace(ns)
}
