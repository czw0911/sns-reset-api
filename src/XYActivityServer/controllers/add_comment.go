//添加评论
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	"XYAPIServer/XYActivityServer/libs"
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
	comDB.TalkID = u.GetString("TalkID")
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
	
	aDB := new(models.TalkList)
	aDB.TalkID = comDB.TalkID
	isG := aDB.ParseTalkID()
	if !isG {
		resp = XYLibs.RespStateCode["activity_not_find"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	comDB.ActivityID = aDB.ActivityID
	comDB.YearAndMonth =  aDB.YearAndMonth
	comDB.PostTime = time.Now().Unix()
	
	_,err := comDB.Add()
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["activity_tail_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
	avatar,_ := userAvatar.Get(comDB.UID)
	
	aDB.LastComment = comDB.Contents
	aDB.CommentUser = avatar.NickName
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

