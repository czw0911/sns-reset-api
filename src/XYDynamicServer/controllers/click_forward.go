//转发
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYLibs"
	//"fmt"
	"github.com/astaxie/beego"

)



type ClickForwardController struct {
	BaseController
}

func (u *ClickForwardController) Post() {
	actDB := new(models.DynamicList)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	actDB.DynamicID = u.GetString("DynamicID")
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
	

	res,err := actDB.ClickForward()
	if res {
		resp = XYLibs.RespStateCode["ok"]	
	}else{
		beego.Error(err.Error())
		resp = XYLibs.RespStateCode["dynamic_click_forward_fail"]
	}
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *ClickForwardController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

