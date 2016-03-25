//注册
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}

func (u *LoginController) Post() {

	resp := XYLibs.RespStateCode["nil"]

	regType, _ := u.GetInt8("RegType")
	account := u.GetString("Account")
	pwd := u.GetString("PassWord")
	sign := u.GetString("Sign")

	
	auth := XYLibs.CheckSign(u.Ctx, sign,[]string{"Sign"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	LoginDB := new(models.UserBase)
	DetailInfo := new(models.UserDetailInfo)

	LoginDB.Account = account
	LoginDB.UID = XYLibs.ConvertAccountToUID(account)
	LoginDB.PassWord = XYLibs.HashLoginPassword(account, pwd)
	LoginDB.RegType = regType
	isExist,_ := LoginDB.IsUIDExist()
	if !isExist {
		resp = XYLibs.RespStateCode["user_account_not_exist"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	res, data, err := LoginDB.Login()
	resp = XYLibs.RespStateCode["login_fail"]
	if err != nil {
		beego.Error(err)
	}
	if res {
		
		objUserAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		avatar, err := objUserAvatar.Get(LoginDB.UID)
		if err == nil {
			if avatar.LastLoginTime == 0 && avatar.HomeVoice != "" {
				//初次登陆机器人认证乡音
				go func(){
					actDB := XYLibs.NewVoiceAuthOK(libs.RedisDBUser)
					actDB.UID = LoginDB.UID
					actDB.SetRecvNum()
					//获取机器人昵称
					nikeName := ""
					robotUID := XYLibs.RobotRandom()
					acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
					objAca,_ := acatar.Get(robotUID)
					nikeName = objAca.NickName
					
					if nikeName != "" {
						dbRemindMsg := new(models.UserRemindMsg)
						dbRemindMsg.UID = LoginDB.UID
						dbRemindMsg.MsgTypeID = XYLibs.REMIND_MESSAGE_TYPE_B
						dbRemindMsg.LastMsg = nikeName + ",认证了你的乡音。"
						dbRemindMsg.LastTime = time.Now().Unix()
						dbRemindMsg.Add()					
					}
					
				}()
			}
			avatar.SetRedisConnect(libs.RedisDBUser)
			avatar.LastLoginTime = time.Now().Unix()
			//fmt.Printf("%#v",avatar)
			avatar.Set()
			
		}

		allInfo := make(map[string]interface{})
		token := XYLibs.GenerateToken()
		allInfo["AccessToken"] = token

		//user detail info
		DetailInfo.UID = LoginDB.UID
		DetailInfo.HomeProvinceID, _ = strconv.Atoi(data[0]["HomeProvinceID"].(string))
		Detail := DetailInfo.GeteDetailInfoUID()
		delete(data[0], "HomeProvinceID")
		allInfo["Base"] = data

		//fmt.Printf("%#v\n", Detail)
		if len(Detail) > 0 {

			if Detail[0]["TagID"] != nil {
				Detail[0]["TagID"] = strings.Split(Detail[0]["TagID"].(string), ",")
			}
		}else{
			resp = XYLibs.RespStateCode["user_detail_error"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
		allInfo["Detail"] = Detail

		fielIP := beego.AppConfig.String("file_server_ip")
		for _, v := range allInfo["Detail"].([]orm.Params) {
			if v["Avatar"] != "" && v["Avatar"] != nil {
				v["Avatar"] = fmt.Sprintf("%s=%s", fielIP, v["Avatar"])
			}
			if v["Thumbnail"] != "" && v["Thumbnail"] != nil {
				v["Thumbnail"] = fmt.Sprintf("%s=%s", fielIP, v["Thumbnail"])
			}
			if v["HomeVoice"] != "" && v["HomeVoice"] != nil  {
				v["HomeVoice"] = fmt.Sprintf("%s=%s", fielIP, v["HomeVoice"])
			}
		}

		noSQLKey := fmt.Sprintf("%s:%d", XYLibs.NO_SQL_USER_LOGIN_TOKEN, LoginDB.UID)

		err = libs.RedisDBUser.SETEX(noSQLKey, 3600*24, token)
		if err != nil {
			beego.Error(err)
		}
		resp = XYLibs.RespStateCode["ok"]
		resp.Info = allInfo
		
		//登陆日志
		go func(regType int8,avatar *XYLibs.UserAvatar){
			params := &models.LogsLogin{
				UID :avatar.UID,
				HomeProvinceID:avatar.HomeProvinceID,
				HomeCityID:avatar.HomeCityID,
				HomeDistrictID:avatar.HomeDistrictID,
				LivingProvinceID:avatar.LivingProvinceID,
				LivingCityID:avatar.LivingCityID,
				LivingDistrictID:avatar.LivingDistrictID,
				RegType:regType,
				LoginTime:time.Now().Unix(),
			}
			logs := new(models.Logs)
			err := logs.AddLogsLogin(params)
			if err != nil {
				beego.Error(err)
			}
		}(LoginDB.RegType,&avatar)
	}

	u.Data["json"] = resp
	u.ServeJson()
	return

}

func (u *LoginController) Get() {

	//	a := models.UserAvater{1111,"david","aaaa","bbbbb"}
	//	b := models.UserAvater{2222,"candy","2222aa","22bbbbb"}

	//	a.Set()
	//	b.Set()

	//	c := []string{"1111","2222"}
	//	_,f,_ := a.GetAll(c)
	//	println(f[0].UID,f[0].NickName)

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}
