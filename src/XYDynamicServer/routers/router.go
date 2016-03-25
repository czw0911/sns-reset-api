
package routers

import (
	"XYAPIServer/XYDynamicServer/controllers"
	"github.com/astaxie/beego"
)


func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/dynamic",			
			beego.NSRouter("/add_dynamic",&controllers.AddDynamicController{}),
			beego.NSRouter("/del_dynamic",&controllers.DelDynamicController{}),
			beego.NSRouter("/get_dynamic_list",&controllers.DynamicListController{}),
			beego.NSRouter("/click_good",&controllers.ClickGoodController{}),
			beego.NSRouter("/click_forward",&controllers.ClickForwardController{}),
			beego.NSRouter("/get_comment_list",&controllers.CommentListController{}),
			beego.NSRouter("/add_comment",&controllers.AddCommentController{}),
			beego.NSRouter("/del_comment",&controllers.DelCommentController{}),
			beego.NSRouter("/get_user_dynamic",&controllers.UserDynamicListController{}),
			beego.NSRouter("/share",&controllers.DynamicShareController{}),
		),
	)
	beego.AddNamespace(ns)
}
