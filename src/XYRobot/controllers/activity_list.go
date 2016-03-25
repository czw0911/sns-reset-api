//活动列表
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"XYAPIServer/XYLibs"
)

type ActivityListController struct {
	BaseController
}

func (u *ActivityListController) Get() {
	actDB := models.NewActivityList()
	resp := XYLibs.RespStateCode["ok"]	
	resp.Info = actDB.GetAllActivity()
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *ActivityListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

