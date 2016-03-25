//乡音范例
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"
)



type HomeVoiceExampleController struct {
	BaseController
}

func (u *HomeVoiceExampleController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *HomeVoiceExampleController) Get() {
	
	db := new(models.HomeVoiceExample)
	resp :=   XYLibs.RespStateCode["ok"]
	resp.Info = db.GetAll()
	u.Data["json"] = resp
	u.ServeJson()
	
}


