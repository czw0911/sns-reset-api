//获取乡友信息
package controllers

import (
	"XYAPIServer/XYAccoutServer/libs"
	"XYAPIServer/XYAccoutServer/models"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type TowneeListController struct {
	BaseController
}

func (u *TowneeListController) Post() {

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}

func (u *TowneeListController) Get() {
	resp := XYLibs.RespStateCode["ok"]
	db := new(models.UserDetailInfo)
	uid, _ := u.GetInt64("UID")
	db.UID = uint32(uid)
	db.HomeProvinceID, _ = u.GetInt("HomeProvinceID")
	db.HomeCityID, _ = u.GetInt("HomeCityID")
	db.LivingProvinceID, _ = u.GetInt("LivingProvinceID")
	db.LivingCityID, _ = u.GetInt("LivingCityID")
	db.PageType, _ = u.GetInt8("PageType")
	db.MaxID = u.GetString("MaxID")
	sign := u.GetString("Sign")

	//println(sign)

	loginToken := GetLoginToken(db.UID)
	if loginToken == "" {
		resp = XYLibs.RespStateCode["login_token_expire"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	auth := XYLibs.CheckLoginSign(u.Ctx, sign, loginToken,[]string{"Sign"})
	if !auth {
		resp = XYLibs.RespStateCode["sign_error"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	paging := XYLibs.NewPaging()
	arrData, resp := paging.PageingSingleTable(db)

	if len(arrData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	arrRes := make([]XYLibs.UserAvatar, 0, XYLibs.TABLE_LIMIT_NUM)
	arrUID := make([]string, 0, XYLibs.TABLE_LIMIT_NUM)
	for _, v := range arrData["List"].([]orm.Params) {

		arrUID = append(arrUID, v["UID"].(string))
		//delete(v,"ID")
	}
	if len(arrUID) > 0 {
		userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
		userAvatar.UID = db.UID
		fielIP := beego.AppConfig.String("file_server_ip")
		_, arrAvatar, err := userAvatar.GetAll(arrUID, fielIP)
		//fmt.Printf("%v\n",arrAvater)
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
	resp = XYLibs.RespStateCode["ok"]
	resp.Info = arrData
	u.Data["json"] = resp
	u.ServeJson()

}
