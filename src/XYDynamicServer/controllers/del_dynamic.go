//删除动态
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
)



type DelDynamicController struct {
	BaseController
}

func (u *DelDynamicController) Post() {
	comDB := new(models.DynamicList)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	comDB.UID = uint32(uid)
	comDB.DynamicID = u.GetString("DynamicID")
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(comDB.UID)
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
	
	isG := comDB.ParseDynamicID()
	if !isG {
		resp = XYLibs.RespStateCode["dynamic_not_find"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	_,err := comDB.ClickDelete()
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["dynamic_delete_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	resp = XYLibs.RespStateCode["ok"]	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *DelDynamicController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

