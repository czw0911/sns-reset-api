//提醒设置信息
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"

)



type RemindSetInfoController struct {
	BaseController
}

func (u *RemindSetInfoController) Get() {
	db := new(models.UserRemindSet)
	resp := XYLibs.RespStateCode["ok"]
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
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

	res := db.Get()
	
	if len(res) > 0 {
		resp.Info = res[0]
	}else{
		resp = XYLibs.RespStateCode["user_remind_set_get_fail"]
	}

	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *RemindSetInfoController) Post() {
	
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}


