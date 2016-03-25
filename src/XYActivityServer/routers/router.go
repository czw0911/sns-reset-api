
package routers

import (
	"XYAPIServer/XYActivityServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/activity",
			
			beego.NSRouter("/get_talk_list",&controllers.TalkListController{}),
			beego.NSRouter("/add_talk",&controllers.TalkActivityController{}),
			beego.NSRouter("/del_talk",&controllers.DelTalkController{}),
			beego.NSRouter("/click_good",&controllers.ClickGoodController{}),
			beego.NSRouter("/click_bad",&controllers.ClickBadController{}),
			beego.NSRouter("/get_activity_list",&controllers.ActivityListController{}),
			beego.NSRouter("/get_comment_list",&controllers.CommentListController{}),
			beego.NSRouter("/add_comment",&controllers.AddCommentController{}),
			beego.NSRouter("/del_comment",&controllers.DelCommentController{}),
			beego.NSRouter("/share",&controllers.TalkListShareController{}),
			beego.NSRouter("/user_talk_list",&controllers.UserTalkListController{}),	
		),
		beego.NSNamespace("/file",
			beego.NSRouter("/view_file",&controllers.FileViewController{}),			
		),
	)
	beego.AddNamespace(ns)
}
