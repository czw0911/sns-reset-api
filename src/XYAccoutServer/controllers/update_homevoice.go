//修改乡音
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYAccoutServer/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"fmt"
	"os"
	"time"
)

type Size interface {
    Size() int64
}

type UpdateHomeVoiceController struct {
	BaseController
}

func (u *UpdateHomeVoiceController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *UpdateHomeVoiceController) Post() {
	localDB := new(models.UserDetailInfo)
	uid,_ := u.GetInt64("UID")
	localDB.UID = uint32(uid)
	localDB.VoiceLen,_  = u.GetInt("VoiceLen")
	voiceFile,voiceHeader, _ := u.Ctx.Request.FormFile("HomeVoice")	
	sign := u.GetString("Sign")
		
	loginToken := GetLoginToken(localDB.UID)
	if loginToken == "" {
		resp := XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign","HomeVoice"})
	
	if !auth {
		resp := XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	if voiceSizeInterface, ok := voiceFile.(Size); ok {
		 beego.Info("上传voice文件的大小为: %d mime: %s", voiceSizeInterface.Size(),voiceHeader.Header.Get("Content-Type"))
		
		if voiceSizeInterface.Size() > XYLibs.UPLOAD_FILE_MAX_SIZE {
			 resp := XYLibs.RespStateCode["upload_voice_max"]
			 u.Data["json"] = resp
			 u.ServeJson()
			 return
		}
		ym := time.Now().Format("200601")
		fName,fPath := XYLibs.GetUserUpLoadFileNameAndPath(localDB.UID,ym,"caf")
		_,err := os.Stat(fPath)
		if err != nil {
			err =os.MkdirAll(beego.AppConfig.String("upload_path")+"/"+fPath,0666)
			beego.Error(err)
		}
		saveName :=fmt.Sprintf("%s/%s/%s",beego.AppConfig.String("upload_path"),fPath,fName)
		saveRes := u.SaveToFile("HomeVoice",saveName)
		if saveRes != nil {
			beego.Error(saveRes)
			resp := XYLibs.RespStateCode["upload_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
		localDB.HomeVoice = fName
	}
	
	if localDB.HomeVoice == "" {
		u.Data["json"] = XYLibs.RespStateCode["upload_fail"]
		u.ServeJson()
			return
	}
	
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
	_,e := localDB.UpdateHomeVoice()
	if e != nil {
		beego.Error(e)
		resp = XYLibs.RespStateCode["user_update_homevoice_fail"]
	}

	data.SetRedisConnect(libs.RedisDBUser)
	data.HomeVoice = localDB.HomeVoice
	data.VoiceLen = localDB.VoiceLen
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


