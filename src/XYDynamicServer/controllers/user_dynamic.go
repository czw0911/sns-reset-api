//用户动态
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYDynamicServer/libs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs"
	"fmt"
	"strconv"
)


type UserDynamicListController struct {
	BaseController
}

func (u *UserDynamicListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *UserDynamicListController) Get() {
	resp := XYLibs.RespStateCode["nil"]
	db := new(models.UserDynamic)
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	vuid,_ := u.GetInt64("ViewUID")
	db.ViewUID = uint32(vuid)
	
	db.PageType,_ = u.GetInt8("PageType")
	db.MaxID  = u.GetString("MaxID")	
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
	
	

	isEmpty,err := db.IsEmptyTableName()
	if isEmpty {
		beego.Error(err)
		resp = XYLibs.RespStateCode["ok"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	paging := XYLibs.NewPaging()
	arrData,resp =  paging.PaginMultiDateTableByHashIDName(db,"DynamicID",19,25)
	
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
		if v["DynamicID"] != nil {
			//待优化查看数
			go func(id string){
				db.DynamicID = id
				db.ClickView()
			}(v["DynamicID"].(string))
		}
		db.DynamicID = v["DynamicID"].(string)
		v["IsClickGood"],_ = db.IsClickGood()
	}
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		userAvatar.UID = db.UID
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


