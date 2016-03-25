//动态分享
package controllers

import (
	"XYAPIServer/XYDynamicServer/models"
	"XYAPIServer/XYDynamicServer/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"fmt"
	"strconv"
)


type DynamicShareController struct {
	BaseController
}

func (u *DynamicShareController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *DynamicShareController) Get() {
	resp := XYLibs.RespStateCode["nil"]
	db := new(models.DynamicList)
	db.DynamicID =  u.GetString("DynamicID")
	
	isG := db.ParseDynamicID()
	if !isG {
		resp = XYLibs.RespStateCode["dynamic_not_find"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	arrData :=  db.GetDynamicInfo()
	
	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	datatLen := len(arrData)

	fielIP := beego.AppConfig.String("file_server_ip")
	
	arrUID := make([]string,0,datatLen)
	for _, v := range arrData {
		if v["Images"] != "" {
			v["Images"] = fmt.Sprintf("%s=%s",fielIP,v["Images"])
		}
		
		if v["Voices"] != "" {
			v["Voices"] = fmt.Sprintf("%s=%s",fielIP,v["Voices"])
		}

		arrUID = append(arrUID,v["UID"].(string))
		delete(v,"ID")
		db.DynamicID = v["DynamicID"].(string)
		v["IsClickGood"],_ = db.IsClickGood()
	}
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		_,arrAvatar,err :=userAvatar.GetAll(arrUID,fielIP)
		//fmt.Printf("%v\n",arrAvater)
		if err != nil {
			beego.Error(err)
		}
		
		for _, v := range arrData{
			postUID,_ := strconv.ParseInt(v["UID"].(string),10,64)
			
			v["PostUser"] =  arrAvatar[uint32(postUID)]
			delete(v,"UID")
		}
	}
	//获取评论信息
	actDB := new(models.Comments)
	actDB.DynamicID = db.DynamicID
	actDB.HomeProvinceID = db.HomeProvinceID
	actDB.YearAndMonth =  db.YearAndMonth
	arrComment := actDB.PageFirst()
	arrUID = make([]string,0,datatLen)
	for _, v := range arrComment {

		arrUID = append(arrUID,v["UID"].(string))
		//delete(v,"ID")
	}
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		fielIP := beego.AppConfig.String("file_server_ip")
		_,arrAvatar,err :=userAvatar.GetAll(arrUID,fielIP)
		//fmt.Printf("%v\n",arrAvater)
		if err != nil {
			beego.Error(err)
		}
		
		for _, v := range arrComment {
			postUID,_ := strconv.ParseInt(v["UID"].(string),10,64)
			v["PostUser"] =  arrAvatar[uint32(postUID)]
			v["CommentID"] = fmt.Sprintf("%s_%s",v["DynamicID"].(string),v["ID"].(string))
			delete(v,"UID")
			delete(v,"ID")
			delete(v,"DynamicID")
		}
	}
	result := make(map[string]interface{},2)
	result["dynamic"] = arrData[0]
	result["comment"] = arrComment
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = result
	
	u.Data["json"] = resp
	u.ServeJson()
	
	
	
	
	
}


