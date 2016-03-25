//乡音认证日志
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
	//"fmt"
	"github.com/astaxie/beego"
	"XYAPIServer/XYAccoutServer/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	//"time"
)



type VoiceAuthLogController struct {
	BaseController
}

func (u *VoiceAuthLogController) Get() {
	actDB := new(models.VoiceAuthLog)
	resp := XYLibs.RespStateCode["ok"]
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
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
	
    paging := XYLibs.NewPaging()
	arrData,resp =  paging.PageingSingleTable(actDB)
	
	if arrData["List"] == nil {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	datatLen := len(arrData["List"].([]orm.Params))
	fielIP := beego.AppConfig.String("file_server_ip")
	
	arrUID := make([]string,0,datatLen)
	for _, v := range arrData["List"].([]orm.Params) {

		arrUID = append(arrUID,v["AuthUID"].(string))
		delete(v,"ID")

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
			postUID,_ := strconv.ParseInt(v["AuthUID"].(string),10,64)
			
			v["PostUser"] =  arrAvatar[uint32(postUID)]
			delete(v,"AuthUID")
		}
	}
	
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	
	u.Data["json"] = resp
	u.ServeJson()


}


func (u *VoiceAuthLogController) Post() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

