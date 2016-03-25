//动态列表
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYDynamicServer/libs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs"
	"fmt"
	"strconv"
	"strings"
)


type DynamicListController struct {
	BaseController
}

func (u *DynamicListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *DynamicListController) Get() {
	resp := XYLibs.RespStateCode["nil"]
	db := new(models.DynamicList)
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	uid,_ := u.GetInt64("UID")
	db.UID = uint32(uid)
	db.HomeProvinceID,_ = u.GetInt("HomeProvinceID")
	db.HomeCityID,_ = u.GetInt("HomeCityID")
	db.LivingProvinceID,_ = u.GetInt("LivingProvinceID")
	db.LivingCityID,_ = u.GetInt("LivingCityID")
	db.PageType,_ = u.GetInt8("PageType")
	db.MaxID  = u.GetString("MaxID")	
	sign := u.GetString("Sign")
	
	//if beego.AppConfig.String("enable_authentication") != "false" {	
		
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
	
	//}
	
	if db.HomeProvinceID == 0 {
		resp = XYLibs.RespStateCode["dynamic_provinceid_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
//	if db.LivingProvinceID == 0 && db.LivingCityID == 0 {
//		resp = XYLibs.RespStateCode["dynamic_cityid_null"]
//		u.Data["json"] = resp
//		u.ServeJson()
//		return 
//	}
	
	isEmpty,err := db.IsEmptyTableName()
	if isEmpty {
		beego.Error(err)
		resp = XYLibs.RespStateCode["dynamic_list_is_null"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	paging := XYLibs.NewPaging()
	arrData,resp =  paging.PaginMultiDateTable(db)
	
	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	datatLen := len(arrData["List"].([]orm.Params))

	
	fielIP := beego.AppConfig.String("file_server_ip")
	
	arrUID := make([]string,0,datatLen)
	for _, v := range arrData["List"].([]orm.Params) {
		
		v["Images"] = strings.TrimSpace(v["Images"].(string))
		v["Voices"] = strings.TrimSpace(v["Voices"].(string))
		
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


