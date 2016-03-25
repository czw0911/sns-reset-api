//添加反馈
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"time"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"

)



type AddFeedbackController struct {
	BaseController
}

func (u *AddFeedbackController) Post() {
	db := new(models.Feedback)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	db.Contents = u.GetString("Contents")
	db.Contact = u.GetString("Contact")
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
	if db.Contents == "" {
		resp = XYLibs.RespStateCode["feedback_content_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	db.PostTime = time.Now().Unix()
	res,err := db.Add()
	
	if res {
		resp = XYLibs.RespStateCode["ok"]
	}else{
		resp = XYLibs.RespStateCode["feedback_content_fail"]
		beego.Error(err)
	}

	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *AddFeedbackController) Get() {
	
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}


