//获取昵称信息
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"

)



type GetNickNameController struct {
	BaseController
}

func (u *GetNickNameController) Post() {
	actDB := XYLibs.NewUserAvatar(libs.RedisDBUser)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	vid,_ := u.GetInt64("ViewUID")
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
	
	data ,err := actDB.Get(uint32(vid))

	if err != nil {
		beego.Error(err.Error())
		resp = XYLibs.RespStateCode["nil"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	fileServerAddr := beego.AppConfig.String("file_server_ip")
	res := make(map[string]string)
	res["NickName"] = data.NickName
	res["Avatar"] = fmt.Sprintf("%s=%s", fileServerAddr, data.Avatar)
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = res	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *GetNickNameController) Get() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

