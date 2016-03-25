//求乡音认证
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"XYAPIServer/XYAccoutServer/models"
	"strconv"
	"time"
)



type VoiceAuthRequestController struct {
	BaseController
}

func (u *VoiceAuthRequestController) Get() {
	actDB := new(models.VoiceAuthRequest)
	resp := XYLibs.RespStateCode["ok"]
	uid,_ := u.GetInt64("UID")
	actDB.ReuestUID = uint32(uid)
	sign := u.GetString("Sign")
	
	//println(sign)
	loginToken := GetLoginToken(actDB.ReuestUID)
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
	
	acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
	arrUID,err := acatar.GetRandomUID()
	if err != nil {
		beego.Error(err)
	}
	//获取昵称
	nikeName := ""
	objAca,err := acatar.Get(actDB.ReuestUID)
	if err != nil {
		beego.Error(err)
	}
	nikeName = objAca.NickName
	fmt.Printf("%#v\n\n",arrUID)
	if len(arrUID) > 0 {
		for _,v := range arrUID {
			id ,e := strconv.ParseInt(v,10,64)
			if e == nil {
				actDB.UID = uint32(id)
				actDB.Add()
				if nikeName != "" {
					dbRemindMsg := new(models.UserRemindMsg)
					dbRemindMsg.UID = actDB.UID
					dbRemindMsg.MsgTypeID = XYLibs.REMIND_MESSAGE_TYPE_A
					dbRemindMsg.LastMsg = nikeName + ",请求你认证乡音。"
					dbRemindMsg.LastTime = time.Now().Unix()
					dbRemindMsg.Add()
					//发送apns
					go func(uid uint32,msg string){
						suid := fmt.Sprintf("%d",dbRemindMsg.UID)
						pid,_ := acatar.GetCacheIOSPUSHID(suid)
						if pid != "" {
							 apns := libs.NewAPNS()
							 apns.Alert = dbRemindMsg.LastMsg
							 apns.Badge = "1"
							 apns.MsgType = XYLibs.REMIND_MESSAGE_TYPE_A
							 apns.DeviceToken = pid
							 apns.UID = suid
							 go apns.Send()
						}
						
					}(dbRemindMsg.UID,dbRemindMsg.LastMsg)
				}
			}
		}
	}

   
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *VoiceAuthRequestController) Post() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

