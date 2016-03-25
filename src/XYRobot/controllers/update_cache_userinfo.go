//修改用户缓存信息（临时用)
package controllers

import (
	"XYAPIServer/XYRobot/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
)



type UpdateCacheUserInfoController struct {
	BaseController
}

func (u *UpdateCacheUserInfoController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *UpdateCacheUserInfoController) Get() {
	
	uidt,_ := u.GetInt64("UID")
	uid := uint32(uidt)
	ProfessionID, _ := u.GetInt("ProfessionID")
	JobID, _ := u.GetInt("JobID")
	Gender, _ := u.GetInt("Gender")
	Birthday, _ := u.GetInt("Birthday")

	
	
	acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
	data,err := acatar.Get(uid)
	if err != nil {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["user_update_homevoice_fail"]
		u.ServeJson()
			return
	}
	
	resp := XYLibs.RespStateCode["ok"]
	data.SetRedisConnect(libs.RedisDBUser)
	data.Gender = Gender
	data.ProfessionID = ProfessionID
	data.JobID = JobID
	data.Birthday = Birthday
	
	_,err = data.Set()
	if err != nil {
		
  		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_homevoice_fail"]
	}
	
	u.Data["json"] = resp
	u.ServeJson()
	
}


