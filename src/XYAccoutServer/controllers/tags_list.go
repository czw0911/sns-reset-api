//标签列表
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"
)



type TagsListController struct {
	BaseController
}

func (u *TagsListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *TagsListController) Get() {
	
	tabDB := new(models.TagsList)
	resp :=   XYLibs.RespStateCode["ok"]
	resp.Info = tabDB.GetAll()
	u.Data["json"] = resp
	u.ServeJson()	
}


