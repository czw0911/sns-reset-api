//发布动态
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"time"
	//"os"
	//"fmt"
	"strconv"
)



type AddDynamicController struct {
	BaseController
}

func (u *AddDynamicController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *AddDynamicController) Get() {
	
	db := new(models.DynamicList)
	
	account := u.GetString("Account")
	if account != "" {
		db.UID = XYLibs.ConvertAccountToUID(account)
	}
	RegDB := new(models.UserBase)
	RegDB.UID = db.UID
	isExist,err := RegDB.IsUIDExist()	
	if !isExist {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["reg_user_isexist"]
		u.ServeJson()
		return 
	}
	db.DynamicContent = u.GetString("DynamicContent")
	db.HomeProvinceID,_ = strconv.Atoi(u.GetString("HomeProvinceID"))
	db.Images = u.GetString("DynamicImages")
	db.Voices = u.GetString("Voices")
	db.VoiceLen,_ = strconv.Atoi(u.GetString("VoiceLen"))
	db.PostTime = time.Now().Unix()
	db.YearAndMonth = time.Now().Format("200601")
	res,err := db.Add()
	if err != nil {
		beego.Error(err)
	}
	if res {
		u.Data["json"] = XYLibs.RespStateCode["ok"]
	}else{
		u.Data["json"] = XYLibs.RespStateCode["fail"]
	}
	
	u.ServeJson()
	return 

	
}


