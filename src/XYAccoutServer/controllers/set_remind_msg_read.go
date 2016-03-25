//设置提醒消息已读
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"github.com/astaxie/beego"	
	"XYAPIServer/XYLibs"
)



type SetRemindMsgReadController struct {
	BaseController
}

func (u *SetRemindMsgReadController) Get() {
	db := new(models.UserRemindMsg)
	resp := XYLibs.RespStateCode["ok"]
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	db.MsgTypeID = u.GetString("MsgTypeID")
	db.ReadNum,_ = u.GetInt("ReadNum")
	sign := u.GetString("Sign")
	//println(sign)
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

	_,err := db.SetReadNum()
	
	if err != nil {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["user_remind_read_set_fail"]
		u.ServeJson()
		return 
	}
	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *SetRemindMsgReadController) Post() {
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}


