package controllers

import (
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYChatServer/models"
	"XYAPIServer/XYChatServer/libs"
	"fmt"
	"strconv"
)


type SHTController struct {
	beego.Controller
}


func (u *SHTController) Post() {
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}



func (u *SHTController) Get() {
	resp := XYLibs.RespStateCode["nil"]
	uid := u.GetString("UID")
	nickName := u.GetString("NickName")
    sign := u.GetString("Sign")
	intUID,_ := strconv.Atoi(uid)
	
	loginToken := GetLoginToken(uint32(intUID))
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
	
	rcServer := models.NewSHTIMServer()
	rcServer.UID = uid
	rcServer.Nickname = nickName
	reg,err := rcServer.Reg()
	if err != nil {
		beego.Error(err)
	}
	tk,e := rcServer.GetLoginToken()
	if e != nil {
		beego.Error(e)
	}
	if tk.Retcode == 0 {
		resp = XYLibs.RespStateCode["ok"]
		resp.Info = tk.Content
	}
	fmt.Printf("%#v\n",reg)
	fmt.Printf("%#v\n",tk)
	u.Data["json"] = resp
	u.ServeJson()
}


func GetLoginToken(uid uint32) string {
	//baseLock.Lock()
	//defer baseLock.Unlock()
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_USER_LOGIN_TOKEN,uid)
	
	token , err := libs.RedisDBUser.Get(noSQLKey)
	if err != nil {
		beego.Error(err)
		return ""
	}
	if token != nil {
		if t,ok := token.([]uint8);ok{
			return string(t)
		}else{
			println("-----",token.(string))
		}
	}
	return ""
	
}