package controllers

import (
	"encoding/json"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYChatServer/models"
	"fmt"
)


type TokenController struct {
	beego.Controller
}


func (u *TokenController) Post() {
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}



func (u *TokenController) Get() {
	resp := XYLibs.RespStateCode["nil"]
	uid := u.GetString("UID")
	nickName := u.GetString("NickName")
    acatar := u.GetString("Avatar")
	key := beego.AppConfig.String("AppKey")
	secret := beego.AppConfig.String("AppSecret")
	rcServer, rcError := models.NewRCServer(key,secret, "json")
	if rcError != nil {
		beego.Error(rcError)
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	byteData, rcError := rcServer.UserGetToken(uid, nickName, acatar)
	if rcError != nil {
		beego.Error(rcError)
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	var foo interface{}
    json.Unmarshal(byteData, &foo)
	fmt.Printf("%#v\n",foo)
	
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = foo
	u.Data["json"] = resp
	u.ServeJson()
}


