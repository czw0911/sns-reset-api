
package routers

import (
	"XYAPIServer/XYRobot/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		
		beego.NSNamespace("/robot",		
			beego.NSRouter("/register",&controllers.RegController{}),
			beego.NSRouter("/new_reg",&controllers.CRegController{}),
			beego.NSRouter("/add_dynamic",&controllers.AddDynamicController{}),
			beego.NSRouter("/update_home_voice",&controllers.UpdateHomeVoiceController{}),
			beego.NSRouter("/apns",&controllers.APNSController{}),
			beego.NSRouter("/tmp_up_cache_userinfo",&controllers.UpdateCacheUserInfoController{}),
			beego.NSRouter("/update_userinfo",&controllers.UpdateUserInfoController{}),
		),
		
		beego.NSNamespace("/api",		
			beego.NSRouter("/add_dynamic",&controllers.AddDynamicController{}),
			beego.NSRouter("/talk_activity",&controllers.TalkActivityController{}),
			beego.NSRouter("/activity_list",&controllers.ActivityListController{}),
			beego.NSRouter("/reg_num",&controllers.RegNumController{}),
		),
		
		beego.NSNamespace("/open",		
			beego.NSRouter("/get_user_info",&controllers.GeUserInfoController{}),
			beego.NSRouter("/add_activity_num",&controllers.AddActivityViewNumController{}),
		),
	)
	beego.AddNamespace(ns)
}

