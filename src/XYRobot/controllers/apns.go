//修改乡音
package controllers

import (
	"XYAPIServer/XYRobot/models"
	"XYAPIServer/XYRobot/libs"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs"
	"strconv"
	//"fmt"
	//"os"
	//"time"
)



type APNSController struct {
	BaseController
}

func (u *APNSController) Post() {
	
	println(u.Ctx.Request.Header.Get("User-Agent"))
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *APNSController) Get() {
	msgType := u.GetString("type")
	msgContent := u.GetString("msg")
	city := new(models.RegionList)
	data := city.GetProvince()
	if len(data) == 0 {
		u.Data["json"] = "not city"
		u.ServeJson()
		return
	}
	detail := new(models.UserDetailInfo)
	for _,c := range data {
		detail.HomeProvinceID,_ = strconv.Atoi(c["RegionID"].(string))
		arr := detail.GetAllPushID()
		if len(arr) == 0 {
			continue
		}
		go func(mtype,msg string,arr []orm.Params){
			apns := libs.NewAPNS()
			for _,pid := range arr {
				apns.Alert = msg
				apns.Badge = "1"
				apns.MsgType = msgType
				apns.DeviceToken = pid["PushID"].(string)
				apns.UID = pid["UID"].(string)
				apns.Send()
			}
		}(msgType,msgContent,arr)
	}
	u.Data["json"] = "done"
	u.ServeJson()
}


