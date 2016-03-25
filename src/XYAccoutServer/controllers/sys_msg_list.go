//获取系统消息信息
package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SysMsgListController struct {
	BaseController
}

func (u *SysMsgListController) Post() {

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}

func (u *SysMsgListController) Get() {
	resp := XYLibs.RespStateCode["ok"]
	db := new(models.SysMsgLog)
	uidt, _ := u.GetInt64("UID")
	uid := uint32(uidt)
	db.PageType, _ = u.GetInt8("PageType")
	db.MaxID = u.GetString("MaxID")
	sign := u.GetString("Sign")

	//println(sign)

	loginToken := GetLoginToken(uid)
	if loginToken == "" {
		resp = XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	auth := XYLibs.CheckLoginSign(u.Ctx, sign, loginToken,[]string{"Sign"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	paging := XYLibs.NewPaging()
	arrData, resp := paging.PageingSingleTable(db)

	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	isRead := new(models.UserRemindMsg)
	isRead.UID = uid
	lastTime,_ := strconv.Atoi(isRead.GetLastReadSysMsg())
	if db.PageType == XYLibs.PAGE_TYPE_UP && db.MaxID == "" {
	
		isRead.MsgTypeID = XYLibs.REMIND_MESSAGE_TYPE_C
		go isRead.SetReadNum()
	}
	
	for _, v := range arrData["List"].([]orm.Params) {

		v["IsRead"] = "0"
		postTime ,_ := strconv.Atoi(v["PostTime"].(string))
		//println(postTime,"--",lastTime)
		if postTime <= lastTime {
			v["IsRead"] = "1"
		}
	}


	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	u.Data["json"] = resp
	u.ServeJson()

}
