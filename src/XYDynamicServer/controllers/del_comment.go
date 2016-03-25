//删除评论
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
)



type DelCommentController struct {
	BaseController
}

func (u *DelCommentController) Post() {
	comDB := new(models.Comments)
	resp := XYLibs.RespStateCode["nil"]
	uid,_ := u.GetInt64("UID")
	comDB.UID = uint32(uid)
	comDB.CommentID = u.GetString("CommentID")
	sign := u.GetString("Sign")
	arrID := strings.Split(comDB.CommentID,"_")
	if len(arrID) != 2 {
		resp = XYLibs.RespStateCode["dynamic_comment_id_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
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
	comDB.DynamicID = arrID[0]
	comDB.ID,_ = strconv.ParseInt(arrID[1],10,64)
	
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

	isExist := comDB.IsComments()
	if !isExist {
		u.Data["json"] = XYLibs.RespStateCode["dynamic_comment_delete_fail"]
		u.ServeJson()
		return
	}

	_,err := comDB.ClickDelete()
	if err != nil {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["dynamic_comment_delete_fail"]
		u.ServeJson()
		return
	}
	aDB.ClickDelComment()
	resp = XYLibs.RespStateCode["ok"]	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *DelCommentController) Get() {
	
	
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

