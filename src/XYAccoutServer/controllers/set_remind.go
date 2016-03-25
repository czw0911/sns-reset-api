//提醒设置
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"

)



type RemindSetController struct {
	BaseController
}

func (u *RemindSetController) Post() {
	db := new(models.UserRemindSet)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	db.Comment,_ = u.GetInt8("Comment")
	db.Follow,_ = u.GetInt8("Follow")
	db.Activity,_ = u.GetInt8("Activity")
	db.Message,_ = u.GetInt8("Message")
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(db.UID)
	if loginToken == "" {
		resp = XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign"})
	
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}

	res,err := db.Set()
	
	if res {
		resp = XYLibs.RespStateCode["ok"]
	}else{
		resp = XYLibs.RespStateCode["user_remind_set_fail"]
		beego.Error(err)
	}

	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *RemindSetController) Get() {
	
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}


