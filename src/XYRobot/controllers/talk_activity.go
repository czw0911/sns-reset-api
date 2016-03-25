//谈论活动
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"time"
)

type TalkActivityController struct {
	BaseController
}

func (u *TalkActivityController) Get() {
	actDB := new(models.TalkList)
	resp := XYLibs.RespStateCode["nil"]
	account := u.GetString("Account")
	if account != "" {
		actDB.UID = XYLibs.ConvertAccountToUID(account)
	}
	
	RegDB := new(models.UserBase)
	RegDB.UID = actDB.UID
	isExist,err := RegDB.IsUIDExist()	
	if !isExist {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["reg_user_isexist"]
		u.ServeJson()
		return 
	}
	
	actDB.ActivityID,_ = u.GetInt64("ActivityID")
	actDB.TalkContent = u.GetString("TalkContent")
	actDB.VoiceLen,_  = u.GetInt("VoiceLen")
	actDB.Images = u.GetString("Images")
	actDB.Voices = u.GetString("Voices")
	actDB.PostTime = time.Now().Unix()
	actDB.YearAndMonth = time.Now().Format("200601")
	_,err = actDB.TalkActivity()
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["activity_tail_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	resp = XYLibs.RespStateCode["ok"]	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *TalkActivityController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

