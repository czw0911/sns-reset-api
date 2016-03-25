//评论列表
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYDynamicServer/libs"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
)


type CommentListController struct {
	BaseController
}

func (u *CommentListController) Post() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *CommentListController) Get() {
	
	resp := XYLibs.RespStateCode["nil"]
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	resp.Info = arrData
	actDB := new(models.Comments)
	uidt,_ := u.GetInt64("UID")
	uid := uint32(uidt)
	actDB.DynamicID = u.GetString("DynamicID")
	actDB.PageType,_ = u.GetInt8("PageType")
	actDB.MaxID  = u.GetString("MaxID")
	sign := u.GetString("Sign")
	//if beego.AppConfig.String("enable_authentication") != "false" {	
		loginToken := GetLoginToken(uid)
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
	
	aDB := new(models.DynamicList)
	aDB.DynamicID = actDB.DynamicID
	isG := aDB.ParseDynamicID()
	if !isG {
		resp = XYLibs.RespStateCode["dynamic_not_find"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	actDB.HomeProvinceID = aDB.HomeProvinceID
	actDB.YearAndMonth =  aDB.YearAndMonth
	
	
	paging := XYLibs.NewPaging()
	arrData,resp =  paging.PageingSingleTable(actDB)
	
	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	

	
	arrUID := make([]string,0,XYLibs.TABLE_LIMIT_NUM)
	for _, v := range arrData["List"].([]orm.Params) {

		arrUID = append(arrUID,v["UID"].(string))
		//delete(v,"ID")
	}
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		userAvatar.UID = uid
		fielIP := beego.AppConfig.String("file_server_ip")
		_,arrAvatar,err :=userAvatar.GetAll(arrUID,fielIP)
		//fmt.Printf("%v\n",arrAvater)
		if err != nil {
			beego.Error(err)
		}
		
		for _, v := range arrData["List"].([]orm.Params) {
			postUID,_ := strconv.ParseInt(v["UID"].(string),10,64)
			v["PostUser"] =  arrAvatar[uint32(postUID)]
			v["CommentID"] = fmt.Sprintf("%s_%s",v["DynamicID"].(string),v["ID"].(string))
			delete(v,"UID")
			delete(v,"ID")
			delete(v,"DynamicID")
		}
	}
	
	

	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	
	u.Data["json"] = resp
	u.ServeJson()
		
}

