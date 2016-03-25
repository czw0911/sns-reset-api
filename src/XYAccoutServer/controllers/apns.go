//apns
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
)




type ApnsController struct {
	BaseController
}

func (u *ApnsController) Get() {
	
	a := u.GetString("Alert")
	id := u.GetString("DeviceToken")
	typeMsg := u.GetString("MsgType")
	apns := libs.NewAPNS()
	apns.Alert = a
	apns.Badge = "1"
	apns.DeviceToken = id
	apns.MsgType = typeMsg
	go apns.Send()
}


func (u *ApnsController) Post() {
	
	
	u.Data["json"] = "OK"
	u.ServeJson()
		
}


