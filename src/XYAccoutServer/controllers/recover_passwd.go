//找回密码
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYAccoutServer/libs"
	"fmt"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"

)



type ReconverPasswdController struct {
	BaseController
}

func (u *ReconverPasswdController) Post() {
	RecoverPwdBaseDB := new(models.UserBase)
	resp := XYLibs.RespStateCode["nil"]
	account := u.GetString("Account")
	RecoverPwdBaseDB.UID = XYLibs.ConvertAccountToUID(account)
	pwd := u.GetString("PassWord")
	code := u.GetString("Code")
	sign := u.GetString("Sign")
	
	auth := XYLibs.CheckSign(u.Ctx,sign,[]string{"Sign"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	noSQLKey := fmt.Sprintf("%s:%s",XYLibs.NO_SQL_PREFIX_KEY_SMS,account)
	redisDB := libs.RedisDBUser
	verifyCode , err := redisDB.Get(noSQLKey)
	if verifyCode == nil || err != nil {
		resp = XYLibs.RespStateCode["sms_server_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	} 
	
	if string(verifyCode.([]uint8)) != code {
		resp = XYLibs.RespStateCode["sms_verify_code_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	RecoverPwdBaseDB.PassWord = XYLibs.HashLoginPassword(account,pwd)
	
	res,err := RecoverPwdBaseDB.UpdatePassWD()
	
	if res {
		resp = XYLibs.RespStateCode["ok"]
	}else{
		resp = XYLibs.RespStateCode["recover_fail"]
		beego.Error(err)
	}

	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *ReconverPasswdController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

