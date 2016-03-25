//获取活动谈论列表
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	"XYAPIServer/XYActivityServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)


type TalkListController struct {
	BaseController
}

func (u *TalkListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *TalkListController) Get() {
	
	resp := XYLibs.RespStateCode["nil"]
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	resp.Info = arrData
	actDB := new(models.TalkList)
	uidt,_ := u.GetInt64("UID")
	actDB.UID = uint32(uidt)
	actDB.PageType,_ = u.GetInt8("PageType")
	actDB.MaxID  = u.GetString("MaxID")
	actDB.ActivityID,_ = u.GetInt64("ActivityID")
	sign := u.GetString("Sign")
	
	//if beego.AppConfig.String("enable_authentication") != "false" {
	
		loginToken := GetLoginToken(actDB.UID)
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
	//}

	
	isEmpty,err := actDB.IsEmptyTableName()
	if isEmpty {
		beego.Error(err)
		resp = XYLibs.RespStateCode["activity_list_is_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	if actDB.MaxID == "" {
		//增加查看人数	
		list := models.NewActivityList()
		list.ActivityID = actDB.ActivityID
		list.SetJoinNum()
	}
	
	paging := XYLibs.NewPaging()
	arrData,resp =  paging.PaginMultiDateTable(actDB)
	
	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	datatLen := len(arrData["List"].([]orm.Params))

	
	fielIP := beego.AppConfig.String("file_server_ip")
	
	arrUID := make([]string,0,datatLen)
	for _, v := range arrData["List"].([]orm.Params) {
		if v["Images"] != "" {
			v["Images"] = fmt.Sprintf("%s=%s",fielIP,v["Images"])
		}
		
		if v["Voices"] != "" {
			v["Voices"] = fmt.Sprintf("%s=%s",fielIP,v["Voices"])
		}

		arrUID = append(arrUID,v["UID"].(string))
		delete(v,"ID")
		actDB.TalkID = v["TalkID"].(string)
		v["IsClickGoodOrBad"],_ = actDB.IsClickGoodOrBad()
	}
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		userAvatar.UID = actDB.UID
		_,arrAvatar,err :=userAvatar.GetAll(arrUID,fielIP)
		//fmt.Printf("%v\n",arrAvater)
		if err != nil {
			beego.Error(err)
		}
		
		for _, v := range arrData["List"].([]orm.Params) {
			postUID,_ := strconv.ParseInt(v["UID"].(string),10,64)
			
			v["PostUser"] =  arrAvatar[uint32(postUID)]
			delete(v,"UID")
		}
	}
	
	
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	
	u.Data["json"] = resp
	u.ServeJson()
		
}

