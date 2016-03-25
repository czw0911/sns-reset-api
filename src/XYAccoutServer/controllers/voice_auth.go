//乡音认证
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
	"XYAPIServer/XYAccoutServer/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"fmt"
)



type VoiceAuthController struct {
	BaseController
}

func (u *VoiceAuthController) Post() {
	actDB := XYLibs.NewVoiceAuthOK(libs.RedisDBUser)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	actDB.AuthUID = u.GetString("AuthUID")
	answer,_ := u.GetInt8("Answer")
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
	
	if actDB.AuthUID == ""{
		resp = XYLibs.RespStateCode["user_uid_homevoice_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	bakUID := actDB.UID
	intAuthUID,_ := strconv.ParseInt(actDB.AuthUID,10,64)
	
	if actDB.UID == uint32(intAuthUID) {
		resp = XYLibs.RespStateCode["user_self_homevoice_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	IsAuthVoice,_ := actDB.IsAuth()

	if IsAuthVoice != 0 {
		resp = XYLibs.RespStateCode["user_mulauth_homevoice_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	if answer == 0 {
		actDB.SetSendErrorNum()
	}else{
		db := new(models.VoiceAuthLog)
		db.UID = actDB.UID
		db.AuthUID = actDB.AuthUID
		db.AuthTime = time.Now().Unix()
		db.AuthType = XYLibs.AUTH_HOME_VOICE_TYPE_SEND
		_,err := db.Add()
		if err != nil {
			beego.Error(err)
		}
		db.UID =  uint32(intAuthUID)
		db.AuthUID = fmt.Sprintf("%d",actDB.UID)
		db.AuthType = XYLibs.AUTH_HOME_VOICE_TYPE_RECV
		_,err = db.Add()
		if err != nil {
			beego.Error(err)
		}
		
		_,err = actDB.SetSendIndex()
		if err != nil {
			beego.Error(err)
			resp = XYLibs.RespStateCode["user_save_homevoice_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return 
		}
		actDB.SetRanking()
		actDB.UID = uint32(intAuthUID)
		actDB.SetRecvNum()
		
		go func(){
			//获取昵称
			nikeName := ""
			acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
			objAca,_ := acatar.Get(bakUID)
			nikeName = objAca.NickName
			
			if nikeName != "" {
				dbRemindMsg := new(models.UserRemindMsg)
				dbRemindMsg.UID = uint32(intAuthUID)
				dbRemindMsg.MsgTypeID = XYLibs.REMIND_MESSAGE_TYPE_B
				dbRemindMsg.LastMsg = nikeName + ",认证了你的乡音。"
				dbRemindMsg.LastTime = time.Now().Unix()
				dbRemindMsg.Add()
				//发送apns
				suid := fmt.Sprintf("%d",dbRemindMsg.UID)
				pid,_ := acatar.GetCacheIOSPUSHID(suid)
				if pid != "" {
							apns := libs.NewAPNS()
							apns.Alert = dbRemindMsg.LastMsg
							apns.Badge = "1"
							apns.MsgType = XYLibs.REMIND_MESSAGE_TYPE_B
							apns.DeviceToken = pid
							apns.UID = suid
							go apns.Send()
				}
				
			}
			
		}()
	}
	
	//删除求乡音认证列表记录
	vAuth := new(models.VoiceAuthRequest)
	vAuth.UID = bakUID
	vAuth.ReuestUID = uint32(intAuthUID)
	_,err := vAuth.Del()
	if err != nil {
		beego.Error(err)
	}

	resp = XYLibs.RespStateCode["ok"]	
	u.Data["json"] = resp
	u.ServeJson()


}


func (u *VoiceAuthController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

