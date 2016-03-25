//注册
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"XYAPIServer/XYRobot/libs"
	//"fmt"
	"github.com/astaxie/beego"
	"time"
	"math/rand"
	"XYAPIServer/XYLibs"
)


func randomRegDate() int64 {

	arr := make([]int64,6,6)
	arr[0] = 86400	
	arr[1] = 432000	
	arr[2] = 864000	
	arr[3] = 3456000
	arr[4] = 1728000	
	arr[5] = 777600
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := r.Int31n(6)
	if t < 6 {
	    return time.Now().Unix() -  arr[t]
	}
	return time.Now().Unix() -  arr[1]
}

type RegController struct {
	BaseController
}

func (u *RegController) Get() {
	
	resp := XYLibs.RespStateCode["nil"]
	
	regType,_ := u.GetInt8("RegType")
	account := u.GetString("Account")
	pwd := u.GetString("PassWord")
	
	

	
	RegDB := new(models.UserBase)
	detailDB := new(models.UserDetailInfo)
	
	RegDB.Account = account
	RegDB.UID = XYLibs.ConvertAccountToUID(account)
	RegDB.PassWord = XYLibs.HashLoginPassword(account,pwd)
	RegDB.RegType = regType
	RegDB.RegisterTime = randomRegDate()
	
	detailDB.UID = RegDB.UID
	detailDB.HomeProvinceID,_ = u.GetInt("HomeProvinceID")
	detailDB.HomeCityID,_ = u.GetInt("HomeCityID")
	detailDB.HomeDistrictID,_ = u.GetInt("HomeDistrictID")
	detailDB.LivingProvinceID,_ = u.GetInt("LivingProvinceID")
	detailDB.LivingCityID,_ = u.GetInt("LivingCityID")
	detailDB.LivingDistrictID,_ = u.GetInt("LivingDistrictID")
	detailDB.NickName = u.GetString("NickName")
	detailDB.ProfessionID, _ = u.GetInt("ProfessionID")
	detailDB.JobID, _ = u.GetInt("JobID")
	detailDB.Gender, _ = u.GetInt("Gender")
	detailDB.Birthday, _ = u.GetInt("Birthday")
	detailDB.TagID = u.GetString("TagID")
	detailDB.DiySign = u.GetString("DiySign")
	
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
		acatar.LastLoginTime = randomRegDate()
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
	}
	u.Data["json"] = XYLibs.RespStateCode["ok"]
	u.ServeJson()
	return 

}


func (u *RegController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
	
	
}


