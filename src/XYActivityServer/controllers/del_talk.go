//删除谈论活动
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
)



type DelTalkController struct {
	BaseController
}

func (u *DelTalkController) Post() {
	actDB := new(models.TalkList)
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	actDB.TalkID = u.GetString("TalkID")
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(actDB.UID)
	if loginToken == "" {
		u.Data["json"] = XYLibs.RespStateCode["login_token_expire"]
		u.ServeJson()
		return 
	}
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign"})
	if !auth {
		u.Data["json"] = XYLibs.RespStateCode["sign_error"]
		u.ServeJson()
		return 
	}
		
	isG := actDB.ParseTalkID()
	if !isG {
		u.Data["json"] = XYLibs.RespStateCode["activity_not_find"]
		u.ServeJson()
		return 
	}
	
	_,err := actDB.ClickDelete()
	if err != nil {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["activity_delete_fail"]
		u.ServeJson()
		return
	}	
	u.Data["json"] = XYLibs.RespStateCode["ok"]
	u.ServeJson()
	return 
	
}


func (u *DelTalkController) Get() {
		
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

