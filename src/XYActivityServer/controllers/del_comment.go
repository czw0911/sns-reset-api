//删除评论
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
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
	uid,_ := u.GetInt64("UID")
	comDB.UID = uint32(uid)
	comDB.CommentID = u.GetString("CommentID")
	sign := u.GetString("Sign")
	
	arrID := strings.Split(comDB.CommentID,"_")
	if len(arrID) != 2 {
		u.Data["json"] = XYLibs.RespStateCode["activity_comment_id_error"]
		u.ServeJson()
		return 
	}
	
	loginToken := GetLoginToken(comDB.UID)
	if loginToken == "" {
		u.Data["json"] = XYLibs.RespStateCode["login_token_expire"]
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign"})
	if !auth {
		u.Data["json"] = XYLibs.RespStateCode["sign_error"]
		u.ServeJson()
		return 
	}
	comDB.TalkID = arrID[0]
	comDB.ID,_ = strconv.ParseInt(arrID[1],10,64)
	
	aDB := new(models.TalkList)
	aDB.TalkID = comDB.TalkID
	isG := aDB.ParseTalkID()
	if !isG {
		u.Data["json"] = XYLibs.RespStateCode["activity_not_find"]
		u.ServeJson()
		return 
	}
	comDB.ActivityID = aDB.ActivityID
	comDB.YearAndMonth =  aDB.YearAndMonth
	
	isExist := comDB.IsComments()
	if !isExist {
		u.Data["json"] = XYLibs.RespStateCode["activity_comment_delete_fail"]
		u.ServeJson()
		return
	}
	
	_,err := comDB.ClickDelete()
	if err != nil {
		beego.Error(err)
		u.Data["json"] = XYLibs.RespStateCode["activity_comment_delete_fail"]
		u.ServeJson()
		return
	}
	aDB.ClickDelComment()
	u.Data["json"] = XYLibs.RespStateCode["ok"]
	u.ServeJson()
	return 

}


func (u *DelCommentController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

