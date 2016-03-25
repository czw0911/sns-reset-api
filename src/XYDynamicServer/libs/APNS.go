//@Description 苹果推送
//@Contact czw@outlook.com

package libs

import (
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
	"fmt"
	"net/url"
)

var PushiOSAddr string

type APNS struct {
	
	Alert string //消息内容
	
	Badge string 
	
	DeviceToken string
	
	MsgType string
	
	UID string //用户id
	
}

func NewAPNS() *APNS {
	return new(APNS)	
}


func (s *APNS) Send(){
	noSQLKey := fmt.Sprintf("%s:%s",XYLibs.NO_SQL_USER_PUSH_IS_SEND,s.UID)
	res ,_ := RedisDBUser.Get(noSQLKey)
	if res != nil && s.MsgType != XYLibs.REMIND_MESSAGE_TYPE_C {
		//println("sss")
		//return
	}
	p := url.Values{}
	p.Add("Alert",s.Alert)
	p.Add("Badge",s.Badge)
	p.Add("MsgType",s.MsgType)
	p.Add("DeviceToken",s.DeviceToken)
	url := fmt.Sprintf("%s?%s",PushiOSAddr,p.Encode())
	XYLibs.HttpGet(url)
	RedisDBUser.SETEX(noSQLKey,21600,1)
	
}


func init() {
	
	PushiOSAddr =  beego.AppConfig.String("ios_push_url")
	if PushiOSAddr == ""{
			panic(" ios_push_url  not config" )
	}

}