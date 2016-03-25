//获取提醒消息列表
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"
	"strconv"
)



type RemindMsgListController struct {
	BaseController
}

func (u *RemindMsgListController) Get() {
	db := new(models.UserRemindMsg)
	resp := XYLibs.RespStateCode["ok"]
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	sign := u.GetString("Sign")
	//println(sign)
	loginToken := GetLoginToken(db.UID)
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
	
	//系统消息（乡音团队)
	sysUnReadNum := db.GetSysMsgUnreadNum()
	if sysUnReadNum > 0 {
		sysMsg := db.GetNewsSysMsg()
		if len(sysMsg) > 0 {
			db.MsgTypeID = XYLibs.REMIND_MESSAGE_TYPE_C
			db.LastMsg = sysMsg[0]["Messages"].(string)
			db.LastTime ,_ =  strconv.ParseInt(sysMsg[0]["PostTime"].(string),10,64)
			db.UnreadNum , _ = strconv.Atoi(db.GetLastReadSysMsg())
			db.Add()
		}
	}
	
	res := db.Get()
	
	if len(res) > 0 {
		for _,v := range res {
			v["MsgTypeName"] = XYLibs.RemindMessageTypeDefine[v["MsgTypeID"].(string)]
			if v["MsgTypeID"].(string) == XYLibs.REMIND_MESSAGE_TYPE_C {
				v["UnreadNum"] = sysUnReadNum
			}
		}
	}
	
	resp.Info = res
	
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *RemindMsgListController) Post() {
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}


