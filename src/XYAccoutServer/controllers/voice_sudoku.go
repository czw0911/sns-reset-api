//乡音九宫格列表
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"fmt"
)



type VoiceSudokuController struct {
	BaseController
}

func (u *VoiceSudokuController) Get() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}


func (u *VoiceSudokuController) Post() {
	
	uid,_ := u.GetInt64("UID")
	sign := u.GetString("Sign")
	SetVoiceSudokuJoinNum()
	//println(sign)
	
	loginToken := GetLoginToken(uint32(uid))
	if loginToken == "" {
		resp := XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	auth := XYLibs.CheckLoginSign(u.Ctx,sign,loginToken,[]string{"Sign"})
	
	if !auth {
		resp := XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return 
	}
	
	arrRes := make([]XYLibs.UserAvatar,0,9)
	objVoiceAuthOK := XYLibs.NewVoiceAuthOK(libs.RedisDBUser)
	objVoiceAuthOK.UID = uint32(uid)
	arrUID := objVoiceAuthOK.RandomUID()

	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		fielIP := beego.AppConfig.String("file_server_ip")
		_,res,err := userAvatar.GetAll(arrUID,fielIP)
		//fmt.Printf("%v\n",arrRes)
		if err != nil {
			beego.Error(err)
		}
		for _,v := range res {
			arrRes = append(arrRes,v)
		}
	}
	
	
	resp :=   XYLibs.RespStateCode["ok"]
	resp.Info = arrRes
	u.Data["json"] = resp
	u.ServeJson()
	
}

//设置活动参与人数
func  SetVoiceSudokuJoinNum() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_JOIN_NUM,10000000)
	err := libs.RedisDBActivity.INCR(noSQLKey)
	if err != nil {
			return false,err
	}
	return true , nil
}
