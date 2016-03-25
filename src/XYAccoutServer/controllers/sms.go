package controllers

import (
	//"XYAPIServer/XYAccoutServer/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"math/rand"
	"time"
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
)



type SMSController struct {
	BaseController
}

func (u *SMSController) Get() {

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


//获取短信息验证码
func (u *SMSController) Post() {
	resp := XYLibs.RespStateCode["nil"]
	phone := u.GetString("Account")
	smsType := 1 //预留
	sign := u.GetString("Sign")
	if phone == "" {
		resp = XYLibs.RespStateCode["sms_phone_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckSign(u.Ctx,sign,[]string{"Sign"})

	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	resp = XYLibs.RespStateCode["sms_server_error"]
	r := sendVerifyCode(smsType,phone)
	if r {
		resp = XYLibs.RespStateCode["ok"]
	}
	u.Data["json"] = resp
	u.ServeJson()
	
	
}

func sendVerifyCode(smsType int,phone string) bool{
	noSQLKey := ""
	expire := 1800
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	verifyCode := fmt.Sprintf("%06d",r.Intn(999999))
	msg := fmt.Sprintf("验证码：%s，请在30分钟内使用。欢迎使用乡音，开启同乡交友之旅~~",verifyCode)
	switch smsType{
		
		case 1 :
		//register
			noSQLKey = fmt.Sprintf("%s:%s",XYLibs.NO_SQL_PREFIX_KEY_SMS,phone)
			
		case 2 :
		//recover password
			noSQLKey = fmt.Sprintf("%s:%d",XYLibs.NO_SQL_PREFIX_KEY_SMS_RECOVER_PWD,phone)
			
		default :
			return false
	}
	
	err := libs.RedisDBUser.SETEX(noSQLKey,expire,verifyCode)	
	if err != nil {
		beego.Error("sms  vefiry code write redis failed!",err)
		return false
	}
	iconvGBK := mahonia.NewEncoder("gbk")	
	gbk := iconvGBK.ConvertString(msg)	
	smsSRV := libs.NewSMSServer()
	b := smsSRV.SendSMS(phone,gbk)
	if b {
		return true
	}
	return false
}



