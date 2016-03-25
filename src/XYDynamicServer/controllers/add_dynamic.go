//发布动态
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYDynamicServer/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"net/url"
	"time"
	"os"
	"fmt"
)

type Size interface {
    Size() int64
}

type EchoDynamicID  struct {
		DynamicID string
}

type AddDynamicController struct {
	BaseController
}

func (u *AddDynamicController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *AddDynamicController) Post() {
	resp := XYLibs.RespStateCode["nil"]
	db := new(models.DynamicList)
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	db.HomeProvinceID,_ = u.GetInt("HomeProvinceID")
	db.HomeCityID,_ = u.GetInt("HomeCityID")
	db.LivingProvinceID,_ = u.GetInt("LivingProvinceID")
	db.LivingCityID,_ = u.GetInt("LivingCityID")
	dynamicContent,e := url.QueryUnescape(u.GetString("DynamicContent"))
	if e != nil {
		beego.Error(e)
		db.DynamicContent = u.GetString("DynamicContent")
	}else{
		db.DynamicContent = dynamicContent
	}
	db.VoiceLen,_  = u.GetInt("VoiceLen")	
	imgFile, imgHeader, _ := u.Ctx.Request.FormFile("Images")
	voiceFile,voiceHeader, _ := u.Ctx.Request.FormFile("Voices")	
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(db.UID)
	if loginToken == "" {
		resp = XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}

	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign","Images","Voices"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	if db.DynamicContent == "" {
		resp = XYLibs.RespStateCode["dynamic_content_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	if db.HomeProvinceID == 0 {
		resp = XYLibs.RespStateCode["dynamic_provinceid_null"]
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
	fName,fPath := XYLibs.GetDynamicUpLoadFileNameAndPath(db.HomeProvinceID,db.UID,ym,"png")
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
		db.Images = fName
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
		fName,fPath := XYLibs.GetDynamicUpLoadFileNameAndPath(db.HomeProvinceID,db.UID,ym,"caf")
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
		db.Voices = fName
	}
	
	db.PostTime = time.Now().Unix()
	db.YearAndMonth = time.Now().Format("200601")
		
	_,err := db.Add()
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["dynamic_add_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	//推送
	go func(){	
		follow := XYLibs.NewFollow(libs.RedisDBUser)
		follow.UID = db.UID
		fdata,err := follow.GetAllFollowMe()
		if err != nil {
			beego.Error(err)
			return
		}
		if len(fdata) == 0 {
			return
		}
		avatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		apns := libs.NewAPNS()
		tstr := ""
		for _,v := range fdata {
			pushID,_ := avatar.GetCacheIOSPUSHID(v)
			if pushID != "" {
				tstr = XYLibs.Substr(db.DynamicContent,10)
				apns.Alert = fmt.Sprintf("【%s...】您关注的乡友发布了一条新动态，快去瞅瞅吧！",tstr)
				apns.Badge = "1"
				apns.MsgType = XYLibs.REMIND_MESSAGE_TYPE_F
				apns.DeviceToken = pushID
				apns.UID = v
				apns.Send()
			}
		}
	}()
	resp = XYLibs.RespStateCode["ok"]	
	resp.Info = EchoDynamicID{db.DynamicID}
	u.Data["json"] = resp
	u.ServeJson()
	return 
	
}


