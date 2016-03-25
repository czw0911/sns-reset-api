//活动分享
package controllers

import (
	"XYAPIServer/XYActivityServer/models"
	"XYAPIServer/XYActivityServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"strconv"
)


type TalkListShareController struct {
	BaseController
}

func (u *TalkListShareController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *TalkListShareController) Get() {
	
	resp := XYLibs.RespStateCode["ok"]
	actDB := new(models.TalkList)
	actDB.TalkID = u.GetString("TalkID")
	
	if actDB.TalkID == ""{
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	pRes := actDB.ParseTalkID()
	if !pRes {
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	//println("---",pRes)
	
	db := models.NewActivityList()
	db.ActivityID = actDB.ActivityID
	activityData := db.GetActivityName()
	if len(activityData) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	data := actDB.GetDataByTalkID()
	if len(data) == 0 {
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	
	fielIP := beego.AppConfig.String("file_server_ip")
	//if data,ok := talkData[0]; ok {
		for _, v := range data {
			if img,ok := v["Images"].(string); ok && img != "" {
				v["Images"] = fmt.Sprintf("%s=%s",fielIP,img)
			}
			
			if voice,ok := v["Voices"].(string); ok && voice != ""{
				v["Voices"] = fmt.Sprintf("%s=%s",fielIP,voice)
			}
			if uid,ok := v["UID"].(string);ok {
				userAvatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
				postUID,_ := strconv.ParseInt(uid,10,64)
				objPostUser,_ := userAvatar.Get(uint32(postUID))
				v["PostUser"] = objPostUser.NickName
				v["Avatar"] =fmt.Sprintf("%s=%s",fielIP,objPostUser.Avatar)
			}
			v["ActivityName"] = activityData[0]["ActivityName"]			
		}
		resp.Info = data
	//}

	u.Data["json"] = resp
	u.ServeJson()
		
}

