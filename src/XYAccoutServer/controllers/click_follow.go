//点关注
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
	//"fmt"
	"github.com/astaxie/beego"

)



type ClickFollowController struct {
	BaseController
}

func (u *ClickFollowController) Post() {
	actDB := XYLibs.NewFollow(libs.RedisDBUser)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	fid,_ := u.GetInt64("FollowUID")
	actDB.FollowUID = uint32(fid)
	sign := u.GetString("Sign")
	
	//println(sign)
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
	
	if actDB.UID == actDB.FollowUID {
		resp = XYLibs.RespStateCode["dynamic_follow_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	actDB.GetFollowState()
	if actDB.FollowState  == 1 ||  actDB.FollowState == 3 {
		resp = XYLibs.RespStateCode["dynamic_mulclick_follow"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	_,err := actDB.SetFollowYou()
	if err != nil {
		beego.Error(err.Error())
		resp = XYLibs.RespStateCode["dynamic_follow_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	
	uidBak := actDB.UID
	actDB.UID = actDB.FollowUID
	actDB.FollowUID = uidBak	
	_,err =actDB.SetFollowMe()
	if err != nil {
		actDB.FollowUID = actDB.UID
		actDB.UID = uidBak
		actDB.RemFollowYou()
		beego.Error(err.Error())
		resp = XYLibs.RespStateCode["dynamic_follow_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	resp = XYLibs.RespStateCode["ok"]	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *ClickFollowController) Get() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

