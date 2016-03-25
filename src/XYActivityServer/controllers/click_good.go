//点赞
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	//"XYAPIServer/XYGroupsServer/libs"
	"XYAPIServer/XYLibs"
	//"fmt"
	"github.com/astaxie/beego"

)



type ClickGoodController struct {
	BaseController
}

func (u *ClickGoodController) Post() {
	actDB := new(models.TalkList)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	actDB.TalkID = u.GetString("TalkID")
	sign := u.GetString("Sign")
	
	
	loginToken := GetLoginToken(actDB.UID)
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
	
	
	res,err := actDB.ClickGood()
	if res {
		resp = XYLibs.RespStateCode["ok"]	
	}else{
		beego.Error(err.Error())
		resp = XYLibs.RespStateCode["activity_click_good_fail"]
	}
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *ClickGoodController) Get() {
	
//	a := models.ActivityTags{2001,"家乡事","255,255,255,1"}
//	b := models.ActivityTags{2002,"瞎扯淡","255,255,255,1"}
//	b.Set()
//	a.Set()
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

