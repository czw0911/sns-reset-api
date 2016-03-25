//@Description 短信息服务
//@Contact czw@outlook.com

package libs

import (
	"fmt"
	"net/url"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	//"strings"
	"time"
	//"io/ioutil"
)

var (
	smsPriority = "5"
	smsGametype  = "2"
	smsActtype	 = "97"
	smsIP 		 = ""
)


type smsServer struct {}

func NewSMSServer() *smsServer {
	return new(smsServer)
}


func (s *smsServer) SendSMS(phone,msg string) bool {
	p := url.Values{}
	p.Add("priority",smsPriority)
	p.Add("gametype",smsGametype)
	p.Add("acttype",smsActtype)
	p.Add("dest_mobile",phone)
	p.Add("msg_content",msg)
	url := fmt.Sprintf("%s/emaysendMsg?%s",smsIP,p.Encode())
	_,_ = httplib.Get(url).SetTimeout(100 * time.Second, 30 * time.Second).Response()
//	req,err := httplib.Get(url).SetTimeout(100 * time.Second, 30 * time.Second).Response()
//	fmt.Printf("\n%#v\n\n",req)
//	fmt.Printf("\n%#v\n\n",err)
//	if req == nil {
//		beego.Error("emaysendMsg connect error:",err) 
//		return false
//	}
//	defer req.Body.Close()
//	body, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		beego.Error("emaysendMsg error:",err) 
//		return false
//	}
//	res := string(body)
//	arrRes := strings.Split(res,"|")
//	if len(arrRes) == 0 {
//		beego.Error("emaysendMsg response is null:",res)
//		return false
//	}
//	if arrRes[0] != "0" {
//		beego.Error("emaysendMsg response error:",res)
//		return false
//	}
	return true
}

func init() {
	smsIP = beego.AppConfig.String("sms_server_ip")
	if smsIP == "" {
		panic("sms server not config")
	}
	smsPriority = beego.AppConfig.String("sms_priority")
	if smsPriority == "" {
		panic("sms server not config Priority")
	}
	smsGametype = beego.AppConfig.String("sms_gametype")
	if smsGametype == "" {
		panic("sms server not config Gametype")
	}
	smsActtype = beego.AppConfig.String("sms_acttype")
	if smsActtype == "" {
		panic("sms server not config Acttype")
	}
}