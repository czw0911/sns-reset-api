//获取我的活动谈论列表
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


type UserTalkListController struct {
	BaseController
}

func (u *UserTalkListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *UserTalkListController) Get() {
	
	resp := XYLibs.RespStateCode["ok"]
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	resp.Info = arrData
	actDB := new(models.UserTalkList)
	uidt,_ := u.GetInt64("UID")
	actDB.UID = uint32(uidt)
	vuid,_ := u.GetInt64("ViewUID")
	actDB.PageType,_ = u.GetInt8("PageType")
	actDB.MaxID  = u.GetString("MaxID")
	sign := u.GetString("Sign")
	
	//println(sign)
	
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
	
	if vuid > 0 {
		actDB.UID = uint32(vuid)
	}
	
	isEmpty,err := actDB.IsEmptyTableName()
	if isEmpty {
		beego.Error(err)
		resp = XYLibs.RespStateCode["activity_list_is_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	paging := XYLibs.NewPaging()
	arrData,resp =  paging.PaginMultiDateTableByHashIDName(actDB,"TalkID",19,27)
	
	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	datatLen := len(arrData["List"].([]orm.Params))
	
	objActName := models.NewActivityList()
	arrActivityName := objActName.GetAllActivity()
	activityName := make(map[string]interface{})
	for _, v := range arrActivityName {		
		if id,ok := v["ActivityID"].(string);ok{
			activityName[id] = v["ActivityName"].(string)
		}
	}
	
	
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
	//fmt.Printf("%#v\n",activityName)
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
			actDB.TalkID = v["TalkID"].(string)
			actDB.ParseTalkID()
			v["ActivityID"] = actDB.ActivityID
			//println(fmt.Sprintf("%s",actDB.ActivityID))
			v["ActivityName"],_ = activityName[fmt.Sprintf("%d",actDB.ActivityID)]
		}
	}
	
	
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	
	u.Data["json"] = resp
	u.ServeJson()
		
}

