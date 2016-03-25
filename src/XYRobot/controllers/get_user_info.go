//获取用户信息
package controllers

import (
	"XYAPIServer/XYRobot/libs"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"strings"
	//"fmt"
	//"encoding/json"
)



type GeUserInfoController struct {
	BaseController
}

func (u *GeUserInfoController) Get() {
	actDB := XYLibs.NewUserAvatar(libs.RedisDBUser)
	resp := XYLibs.RespStateCode["nil"]
	strUID := u.GetString("UID")
	
	arrUID := strings.Split(strUID,",")
	if len(arrUID) > 10 {
		arrUID = arrUID[:10]
	}
	fileServerAddr := beego.AppConfig.String("file_server_ip")
	//fmt.Printf("%#v\n",arrUID)
	res,data ,err := actDB.GetAll(arrUID,fileServerAddr)
	//fmt.Printf("%#v\n",data)
	if err != nil || !res {
		beego.Error(err.Error())
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	arr := make([]XYLibs.UserAvatar,0,10)
	for _,v := range data {
		arr = append(arr,v)
	}
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arr
	u.Data["json"] = resp
	u.ServeJson()
	return 

}


func (u *GeUserInfoController) Post() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

