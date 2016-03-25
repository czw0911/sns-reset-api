package controllers

import (
	"XYAPIServer/XYPushServer/libs"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"github.com/timehop/apns"
	"fmt"
	"time"
)

var iosPush apns.Client

type IOSController struct {
	beego.Controller
}


func (i *IOSController) Post() {
	resp := XYLibs.RespStateCode["method_not_find"]
	i.Data["json"] = resp
	i.ServeJson()
}


func (i *IOSController) Get() {
	
    alert := i.GetString("Alert")
	badge,_ := i.GetInt("Badge")
	deviceToken := i.GetString("DeviceToken")
	msgType := i.GetString("MsgType")
	
	
	p := apns.NewPayload()
	p.APS.Alert.Body = alert
	p.APS.Badge = &badge
	p.APS.Sound = "bingbong.aiff"
	p.SetCustomValue("msg_type",msgType)
	
	m := apns.NewNotification()
	m.Payload = p
	//m.DeviceToken = "6d356613103a308a037124dca3e5a69ae0b8cf50cc8868cf53ba163a25b1a045"
	m.DeviceToken = deviceToken
	fmt.Printf("msg_type:%s \t DeviceToken:%s alert:%s\t \n",msgType,m.DeviceToken,p.APS.Alert)
	err := iosPush.Send(m)
	if err != nil {
		beego.Error(err)
	}
	
}

func init(){
	var err error
	iosPush,err = libs.NewAPNS()
	if err != nil {
		panic(err.Error())
	}
	go func() {
		
	    for f := range iosPush.FailedNotifs {
	        fmt.Println("Notif", f.Notif.ID, "failed with", f.Err.Error())
	    }
		println("fend")
	}()
	go func(){
		f, err :=libs.NewFeedback()
		if err != nil {
			    beego.Error("Could not create feedback", err.Error())
		}
		for {
			for ft := range f.Receive() {
				    fmt.Println("Feedback for token:", ft.DeviceToken)
			}
			time.Sleep(30 * time.Second)
		}
	}()
}
