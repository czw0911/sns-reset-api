//添加评论
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"time"
	"net/url"
)



type AddCommentController struct {
	BaseController
}

func (u *AddCommentController) Post() {
	comDB := new(models.Comments)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	comDB.UID = uint32(uid)
	comDB.DynamicID = u.GetString("DynamicID")
	contents,e := url.QueryUnescape(u.GetString("Contents"))
	if e != nil {
		beego.Error(e)
		comDB.Contents = u.GetString("Contents")
	}else{
		comDB.Contents = contents
	}
	sign := u.GetString("Sign")
	
	loginToken := GetLoginToken(comDB.UID)
	if loginToken == "" {
		resp = XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}	
	
	if comDB.Contents == "" {
		resp = XYLibs.RespStateCode["comment_content_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	
	aDB := new(models.DynamicList)
	aDB.DynamicID = comDB.DynamicID
	isG := aDB.ParseDynamicID()
	if !isG {
		resp = XYLibs.RespStateCode["dynamic_not_find"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	comDB.HomeProvinceID = aDB.HomeProvinceID
	comDB.YearAndMonth =  aDB.YearAndMonth
	comDB.PostTime = time.Now().Unix()
	
	_,err := comDB.Add()
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["dynamic_comment_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	aDB.LastComment = comDB.Contents
	aDB.ClickComment()
	resp = XYLibs.RespStateCode["ok"]	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *AddCommentController) Get() {
	
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

