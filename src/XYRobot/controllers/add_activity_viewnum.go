//增加活动查看次数
package controllers

import (
	"XYAPIServer/XYRobot/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"fmt"
)



type AddActivityViewNumController struct {
	BaseController
}

func (u *AddActivityViewNumController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *AddActivityViewNumController) Get() {
	
	activityID,_ := u.GetInt64("ActivityID")
	u.Data["json"] = XYLibs.RespStateCode["fail"]
	if activityID > 0 {
		noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_JOIN_NUM,activityID)

		err := libs.RedisDBActivity.INCR(noSQLKey)
		if err == nil {
				u.Data["json"] = XYLibs.RespStateCode["ok"]
		}else{
			beego.Error(err)
		}
	}
	u.ServeJson()
	return 
	
	
	
	
}


