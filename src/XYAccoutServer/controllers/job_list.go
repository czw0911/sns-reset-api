//职业列表
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"

)



type JobListController struct {
	BaseController
}

func (u *JobListController) Post() {
//	fmt.Printf("%#v\n",u.Ctx.Request.Form)
//	fmt.Printf("%#v\n",u.Ctx.Request.URL.RawQuery)
//	println(u.Ctx.Request.Form.Get("vvv"))
//	println(u.Ctx.Request.Form.Encode())
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *JobListController) Get() {
//	fmt.Printf("%#v\n",u.Ctx.Request.URL.RawQuery)
//	println(u.Ctx.Request.Method)
//	println(u.Ctx.Request.URL.Query())
//	t := u.Ctx.Request.Form
//	fmt.Printf("%#v\n",t.Encode())
	jobDB := new(models.JobList)
	resp :=   XYLibs.RespStateCode["ok"]
	resp.Info = jobDB.GetAll()
	u.Data["json"] = resp
	u.ServeJson()
	
}


