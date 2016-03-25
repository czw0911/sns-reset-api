//添加推送id
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYAccoutServer/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
)



type AddPushIDController struct {
	BaseController
}

func (u *AddPushIDController) Post() {
	AddPushIDBaseDB := new(models.UserDetailInfo)
	resp := XYLibs.RespStateCode["nil"]
	AddPushIDBaseDB.PushType,_ = u.GetInt8("PushType")
	uid,_ := u.GetInt64("UID")
	AddPushIDBaseDB.UID = uint32(uid)
	AddPushIDBaseDB.PushID = u.GetString("PushID")
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(AddPushIDBaseDB.UID)
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
	avatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
	data,_ := avatar.Get(AddPushIDBaseDB.UID)
	
	if data.HomeProvinceID == 0 {
		resp = XYLibs.RespStateCode["add_pushid_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	AddPushIDBaseDB.HomeProvinceID = data.HomeProvinceID
	res,err := AddPushIDBaseDB.UpdatePUSHID()
	
	if res {
		if AddPushIDBaseDB.PushType == 1 && AddPushIDBaseDB.PushID != "" {
			AddPushIDBaseDB.SetCacheIOSPUSHID()
		}
		resp = XYLibs.RespStateCode["ok"]
	}else{
		resp = XYLibs.RespStateCode["add_pushid_fail"]
		beego.Error(err)
	}

	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *AddPushIDController) Get() {
	
	//a := new(models.UserBase)
	//d,_ := a.GetCacheIOSPUSHID([]string{"4228047827","4228047828","4228047829"})
	//fmt.Printf("%v\n",d)
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}


