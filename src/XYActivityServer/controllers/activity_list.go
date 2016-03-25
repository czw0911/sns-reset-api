//获取活动主题列表
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	//"XYAPIServer/XYGroupsServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"time"
	//"strconv"
)


type ActivityListController struct {
	BaseController
}

func (u *ActivityListController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *ActivityListController) Get() {
	
	resp := XYLibs.RespStateCode["nil"]
	arrData := make(map[string]interface{})
	arrData["MaxID"] = nil
	arrData["List"] = nil
	resp.Info = arrData
	actDB := models.NewActivityList()
	uidt,_ := u.GetInt64("UID")
	uid := uint32(uidt)
	actDB.PageType,_ = u.GetInt8("PageType")
	actDB.MaxID  = u.GetString("MaxID")
	sign := u.GetString("Sign")
	
	//if beego.AppConfig.Bool("enable_authentication") != "false" {
	
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
	//dt := make([]orm.Params,0,models.TABLE_LIMIT_NUM)
	paging := XYLibs.NewPaging()
	arrData,resp =  paging.PageingSingleTable(actDB)
	
	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	
	fielIP := beego.AppConfig.String("file_server_ip")
	
	arrThemeID := make([]string,0,XYLibs.TABLE_LIMIT_NUM)
	for _, v := range arrData["List"].([]orm.Params) {
		//fmt.Printf("%#v\n",v)
		arrThemeID = append(arrThemeID,v["ActivityID"].(string))
		v["JoinNum"] = "0"
		if v["ActivityDesImg"] != "" {
			v["ActivityDesImg"] = fmt.Sprintf("%s=%s",fielIP,v["ActivityDesImg"])
		}
		delete(v,"ID")
	}
	if len(arrThemeID) > 0 {
		_,arrNum,err :=actDB.GetAllJoinNum(arrThemeID)
		//fmt.Printf("%v\n",arrNum)
		if err != nil {
			beego.Error(err)
		}
		l := len(arrNum)
		for k, v := range arrData["List"].([]orm.Params) {
				if k < l {
					v["JoinNum"] =  arrNum[k]
				}
		}
	}
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	
	u.Data["json"] = resp
	u.ServeJson()
		
}

