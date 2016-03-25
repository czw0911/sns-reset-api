//关注我的人
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYLibs"
	"strconv"
	"github.com/astaxie/beego"
	//"fmt"

)



type FollowMeController struct {
	BaseController
}

func (u *FollowMeController) Get() {
	actDB := XYLibs.NewFollow(libs.RedisDBUser)
	actDB.FollowType = 1
	resp := XYLibs.RespStateCode["ok"]
	uid,_ := u.GetInt64("UID")
	actDB.UID = uint32(uid)
	actDB.MaxID,_ = u.GetInt("MaxID")
	actDB.PageType,_ = u.GetInt8("PageType")
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
	

	arrData,err := actDB.Pageing()
	
	if err != nil || arrData["List"] == nil {
		beego.Error(err)
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}

	
	arrRes := make([]XYLibs.UserAvatar, 0, XYLibs.TABLE_LIMIT_NUM)
	arrUID := make([]string, 0, XYLibs.TABLE_LIMIT_NUM)
	for _, v := range arrData["List"].([]string) {

		arrUID = append(arrUID, v)
	}
	//println(arrUID[0])
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		userAvatar.UID = actDB.UID
		fielIP := beego.AppConfig.String("file_server_ip")
		_, arrAvatar, err := userAvatar.GetAll(arrUID, fielIP)
		//fmt.Printf("%#v\n",arrAvatar)
		if err != nil {
			beego.Error(err)
		}

		for _, v := range arrUID {
			id, _ := strconv.ParseInt(v, 10, 64)
			if x, ok := arrAvatar[uint32(id)]; ok {
				arrRes = append(arrRes, x)
			}
		}
	}

	arrData["List"] = arrRes
	//fmt.Printf("%#v\n",arrData["List"])
	resp.Info = arrData
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *FollowMeController) Post() {
	

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
		
}

