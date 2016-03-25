//乡音认证排名
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
)


type RecordRankingController struct {
	BaseController
}

func (u *RecordRankingController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *RecordRankingController) Post() {
	
	uid,_ := u.GetInt64("UID")
	sign := u.GetString("Sign")
	
	//println(sign)
	
	loginToken := GetLoginToken(uint32(uid))
	if loginToken == "" {
		resp := XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign"})
	
	if !auth {
		resp := XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	

	objVoiceAuthOK := XYLibs.NewVoiceAuthOK(libs.RedisDBUser)
	objVoiceAuthOK.UID = uint32(uid)

	resp :=   XYLibs.RespStateCode["ok"]
	resp.Info = objVoiceAuthOK.GetRordRanking()
	u.Data["json"] = resp
	u.ServeJson()
	
}


