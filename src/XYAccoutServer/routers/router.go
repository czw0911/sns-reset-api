// @APIVersion 1.0.0
// @Title 乡音账号服务 API
// @Description 提供账号类的API服务
// @Contact czw@outlook.com

package routers

import (
	"XYAPIServer/XYAccoutServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/common",
			
			beego.NSRouter("/get_region_list",&controllers.RegionController{}),
			beego.NSRouter("/get_job_list",&controllers.JobListController{}),
			beego.NSRouter("/get_tags_list",&controllers.TagsListController{}),
			beego.NSRouter("/home_voice_example",&controllers.HomeVoiceExampleController{}),
			
		),
		beego.NSNamespace("/user",
			
			beego.NSRouter("/get_sms_verifycode",&controllers.SMSController{}),
			beego.NSRouter("/register",&controllers.RegController{}),
			beego.NSRouter("/login",&controllers.LoginController{}),
			beego.NSRouter("/reset_password",&controllers.ReconverPasswdController{}),
			beego.NSRouter("/add_pushid",&controllers.AddPushIDController{}),
			beego.NSRouter("/voice_sudoku_list",&controllers.VoiceSudokuController{}),
			beego.NSRouter("/update_home_voice",&controllers.UpdateHomeVoiceController{}),
			beego.NSRouter("/voice_record_ranking",&controllers.RecordRankingController{}),
			beego.NSRouter("/get_townee_list",&controllers.TowneeListController{}),
			beego.NSRouter("/click_follow",&controllers.ClickFollowController{}),
			beego.NSRouter("/auth_voice",&controllers.VoiceAuthController{}),
			beego.NSRouter("/update_user_info",&controllers.UpdateUserInfoController{}),
			beego.NSRouter("/follow_other_list",&controllers.FollowOtherController{}),
			beego.NSRouter("/follow_me_list",&controllers.FollowMeController{}),
			beego.NSRouter("/auth_voice_log",&controllers.VoiceAuthLogController{}),
			beego.NSRouter("/add_feedback",&controllers.AddFeedbackController{}),
			beego.NSRouter("/set_remind",&controllers.RemindSetController{}),
			beego.NSRouter("/request_auth_voice_list",&controllers.VoiceAuthRequestListController{}),
			beego.NSRouter("/request_auth_voice",&controllers.VoiceAuthRequestController{}),
			beego.NSRouter("/get_remind_msg_list",&controllers.RemindMsgListController{}),
			beego.NSRouter("/set_remind_read_num",&controllers.SetRemindMsgReadController{}),
			beego.NSRouter("/get_remind_set",&controllers.RemindSetInfoController{}),
			beego.NSRouter("/apns",&controllers.ApnsController{}),
			beego.NSRouter("/get_sys_msg_list",&controllers.SysMsgListController{}),
			beego.NSRouter("/get_nickname",&controllers.GetNickNameController{}),
			
		),
	)
	beego.AddNamespace(ns)
}

