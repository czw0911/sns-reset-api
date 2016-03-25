//注册
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYAccoutServer/libs"
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"strconv"
	"XYAPIServer/XYLibs"
)




type RegController struct {
	BaseController
}

func (u *RegController) Post() {
	
	resp := XYLibs.RespStateCode["nil"]
	
	regType,_ := u.GetInt8("RegType")
	account := u.GetString("Account")
	pwd := u.GetString("PassWord")
	code := u.GetString("Code")
	sign := u.GetString("Sign")
	

    if account == "" || pwd == ""{
		resp = XYLibs.RespStateCode["reg_user_pwd_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckSign(u.Ctx,sign,[]string{"Sign"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	if regType != 3 {
		
		noSQLKey := fmt.Sprintf("%s:%s",XYLibs.NO_SQL_PREFIX_KEY_SMS,account)
		verifyCode , err := libs.RedisDBUser.Get(noSQLKey)
		if verifyCode == nil || err != nil {
			resp = XYLibs.RespStateCode["sms_server_error"]
		
			u.Data["json"] = resp
			u.ServeJson()
			return 
		} 
		
		
		if string(verifyCode.([]uint8)) != code {
	
			resp = XYLibs.RespStateCode["sms_verify_code_error"]
			
			u.Data["json"] = resp
			u.ServeJson()
			return 
		}
	}
	
	RegDB := new(models.UserBase)
	detailDB := new(models.UserDetailInfo)
	
	RegDB.Account = account
	RegDB.UID = XYLibs.ConvertAccountToUID(account)
	RegDB.PassWord = XYLibs.HashLoginPassword(account,pwd)
	RegDB.RegType = regType
	RegDB.RegisterTime = time.Now().Unix()
	RegDB.BindPhone,_ = strconv.ParseUint(RegDB.Account,10,64)
	
	detailDB.UID = RegDB.UID
	detailDB.HomeProvinceID,_ = u.GetInt("HomeProvinceID")
	detailDB.HomeCityID,_ = u.GetInt("HomeCityID")
	detailDB.HomeDistrictID,_ = u.GetInt("HomeDistrictID")
	detailDB.LivingProvinceID,_ = u.GetInt("LivingProvinceID")
	detailDB.LivingCityID,_ = u.GetInt("LivingCityID")
	detailDB.LivingDistrictID,_ = u.GetInt("LivingDistrictID")
	
	
	RegDB.HomeProvinceID = detailDB.HomeProvinceID
	isExist,err := RegDB.IsUIDExist()	
	if isExist {
		beego.Error(err)
		resp = XYLibs.RespStateCode["reg_user_isexist"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	res,err := RegDB.Reg()
	resp = XYLibs.RespStateCode["reg_user_fail"]
	if err != nil {
		beego.Error(err)
	}
	
	if res {
		
		detailDB.UID = RegDB.UID
		detailDB.NickName = XYLibs.GenerateRandomNickName()
		detailDB.Avatar = XYLibs.GenerateRandomAvatar()
		detailDB.Thumbnail = detailDB.Avatar
		_,e := detailDB.Reg()
		if e != nil {
			beego.Error(e)
			RegDB.Delete()
			u.Data["json"] = resp
			u.ServeJson()
			return 	
		}
		
		acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		acatar.UID = detailDB.UID
		acatar.NickName = detailDB.NickName
		acatar.Avatar = detailDB.Avatar
		acatar.Thumbnail = detailDB.Thumbnail
		acatar.HomeProvinceID = detailDB.HomeProvinceID
		acatar.HomeCityID = detailDB.HomeCityID
		acatar.HomeDistrictID = detailDB.HomeDistrictID
		acatar.LivingProvinceID = detailDB.LivingProvinceID
		acatar.LivingCityID = detailDB.LivingCityID
		acatar.LivingDistrictID = detailDB.LivingDistrictID
		_,e = acatar.Set()
		if e != nil {
			beego.Error(e)
			RegDB.Delete()
			u.Data["json"] = resp
			u.ServeJson()
			return 	
		}
		_,e = acatar.SetAllRegisterUID()
		if e != nil {
			beego.Error(e)
		}
		
		//注册日志
		go func(regType int8,avatar *XYLibs.UserAvatar){
			params := &models.LogsReg{
				UID :avatar.UID,
				HomeProvinceID:avatar.HomeProvinceID,
				HomeCityID:avatar.HomeCityID,
				HomeDistrictID:avatar.HomeDistrictID,
				LivingProvinceID:avatar.LivingProvinceID,
				LivingCityID:avatar.LivingCityID,
				LivingDistrictID:avatar.LivingDistrictID,
				RegType:regType,
				RegisterTime:time.Now().Unix(),
			}
			logs := new(models.Logs)
			err := logs.AddLogsReg(params)
			if err != nil {
				beego.Error(err)
			}
		}(RegDB.RegType,acatar)
	}
	//推送
	go func(){
		allPushID := detailDB.GetAllPushID()
		if len(allPushID) == 0 {
			return
		}
		apns := libs.NewAPNS()
		for _,pid := range allPushID {
			apns.Alert = "都在同城吗？有新同乡进入乡音，快去瞅瞅！"
			apns.Badge = "1"
			apns.MsgType = XYLibs.REMIND_MESSAGE_TYPE_E
			apns.DeviceToken = pid["PushID"].(string)
			apns.UID = pid["UID"].(string)
			apns.Send()
		}
		
	}()
	u.Data["json"] = XYLibs.RespStateCode["ok"]
	u.ServeJson()
	return 

}


func (u *RegController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
	
	
}


