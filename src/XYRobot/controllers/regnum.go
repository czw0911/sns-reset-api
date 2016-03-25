//注册人数
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs"
	"time"
)




type RegNumController struct {
	BaseController
}

func (u *RegNumController) Get() {
	
	resp := XYLibs.RespStateCode["ok"]
	
	queryType,_ := u.GetInt8("Type")
	
	RegDB := new(models.RegNUM)
	RegDB.RegisterTime = int64(time.Now().Year())
	var res []orm.Params
	switch queryType {
		case 1 :
		  res = RegDB.GetEveryDay()
		default :
		  res = RegDB.GetAllByHome()
		
	}
	resp.Info = res
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *RegNumController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
	
	
}


