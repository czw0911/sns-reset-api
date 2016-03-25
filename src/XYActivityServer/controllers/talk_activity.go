//谈论活动
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	//"XYAPIServer/XYGroupsServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"os"
	"net/url"
	
)

type Size interface {
    Size() int64
}

type EchoTalkID  struct {
		TalkID string
}

type TalkActivityController struct {
	BaseController
}

func (u *TalkActivityController) Post() {
	actDB := new(models.TalkList)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	actDB.ActivityID,_ = u.GetInt64("ActivityID")
	talkContent,e := url.QueryUnescape(u.GetString("TalkContent"))
	if e != nil {
		beego.Error(e)
		actDB.TalkContent = u.GetString("TalkContent")
	}else{
		actDB.TalkContent = talkContent
	}
	
	actDB.VoiceLen,_  = u.GetInt("VoiceLen")	
	imgFile, imgHeader, _ := u.Ctx.Request.FormFile("Images")
	voiceFile,voiceHeader, _ := u.Ctx.Request.FormFile("Voices")
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(actDB.UID)
	if loginToken == "" {
		resp = XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign","Voices","Images"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}	
	
	aDB := models.NewActivityList()
	aDB.ActivityID = actDB.ActivityID
	isG,_ := aDB.IsExist() 
	if !isG {
		resp = XYLibs.RespStateCode["activity_not_find"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	if imgSizeInterface, ok := imgFile.(Size); ok {
		 beego.Info("上传img文件的大小为: %d mime:%s", imgSizeInterface.Size(),imgHeader.Header.Get("Content-Type"))

		 if imgSizeInterface.Size() > XYLibs.UPLOAD_FILE_MAX_SIZE {
			resp = XYLibs.RespStateCode["upload_file_max_size_err"]
			u.Data["json"] = resp
			u.ServeJson()
			return 
		}
		
		ym := time.Now().Format("200601")
		fName,fPath := XYLibs.GetActivityUpLoadFileNameAndPath(actDB.ActivityID,actDB.UID,ym,"png")
		_,err := os.Stat(fPath)
		if err != nil {
			err =os.MkdirAll(beego.AppConfig.String("upload_path")+"/"+fPath,0666)
			if err != nil {
				beego.Error(err)
			}
			
		}
		saveName :=fmt.Sprintf("%s/%s/%s",beego.AppConfig.String("upload_path"),fPath,fName)
		saveRes := u.SaveToFile("Images",saveName)
		if saveRes != nil {
			beego.Error(saveRes)
			resp = XYLibs.RespStateCode["activity_upload_file_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
		actDB.Images = fName
	}
	
	if voiceSizeInterface, ok := voiceFile.(Size); ok {
		 beego.Info("上传voice文件的大小为: %d mime: %s", voiceSizeInterface.Size(),voiceHeader.Header.Get("Content-Type"))
		
		if voiceSizeInterface.Size() > XYLibs.UPLOAD_FILE_MAX_SIZE {
			 resp = XYLibs.RespStateCode["upload_file_max_size_err"]
			 u.Data["json"] = resp
			 u.ServeJson()
			 return
		}
		ym := time.Now().Format("200601")
		fName,fPath := XYLibs.GetActivityUpLoadFileNameAndPath(actDB.ActivityID,actDB.UID,ym,"caf")
		_,err := os.Stat(fPath)
		if err != nil {
			err =os.MkdirAll(beego.AppConfig.String("upload_path")+"/"+fPath,0666)
			beego.Error(err)
		}
		saveName :=fmt.Sprintf("%s/%s/%s",beego.AppConfig.String("upload_path"),fPath,fName)
		saveRes := u.SaveToFile("Voices",saveName)
		if saveRes != nil {
			beego.Error(saveRes)
			resp = XYLibs.RespStateCode["activity_upload_file_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
		actDB.Voices = fName
	}
	
	actDB.PostTime = time.Now().Unix()
	actDB.YearAndMonth = time.Now().Format("200601")
	
	_,err := actDB.TalkActivity()
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["activity_tail_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	
	resp = XYLibs.RespStateCode["ok"]	
	resp.Info = EchoTalkID{actDB.TalkID}
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *TalkActivityController) Get() {
	
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

