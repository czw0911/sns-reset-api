//修改乡音
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"XYAPIServer/XYRobot/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	//"fmt"
	//"os"
	"time"
)



type UpdateHomeVoiceController struct {
	BaseController
}

func (u *UpdateHomeVoiceController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *UpdateHomeVoiceController) Get() {
	localDB := new(models.UserDetailInfo)
	uid,_ := u.GetInt64("UID")
	localDB.UID = uint32(uid)
	localDB.VoiceLen,_  = u.GetInt("VoiceLen")	
	localDB.HomeVoice = u.GetString("HomeVoice")
		

	
	
	acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
	data,err := acatar.Get(localDB.UID)
	if err != nil {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["user_update_homevoice_fail"]
		u.ServeJson()
			return
	}
	
	resp := XYLibs.RespStateCode["ok"]
	localDB.HomeProvinceID = data.HomeProvinceID
	println(localDB.HomeProvinceID)
	_,e := localDB.UpdateHomeVoice()
	if e != nil {
		beego.Error(e)
		resp = XYLibs.RespStateCode["user_update_homevoice_fail"]
	}

	data.SetRedisConnect(libs.RedisDBUser)
	data.HomeVoice = localDB.HomeVoice
	data.VoiceLen = localDB.VoiceLen
	data.LastLoginTime = time.Now().Unix()
	_,err = data.Set()
	if err != nil {
		
  		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_homevoice_fail"]
	}
	
	_,err = localDB.SaveHomeVoiceToRevordList()
	if err != nil {
  		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_homevoice_fail"]
	}
	u.Data["json"] = resp
	u.ServeJson()
	
}


