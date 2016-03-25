//@Description 苹果推送
//@Contact czw@outlook.com

package libs

import (
	"github.com/astaxie/beego"
	"github.com/timehop/apns"
	"os"
)

var (
	APNSAddr string //推送服务地址
	FeedbackAddr string //推送反馈
	APNSCert string //连接证书文件地址
	APNSKey  string //连接密钥文件地址
)


func NewAPNS() (apns.Client,error) {

	return apns.NewClientWithFiles(APNSAddr,APNSCert,APNSKey)
}


func NewFeedback()(apns.Feedback,error){
	return apns.NewFeedbackWithFiles(FeedbackAddr,APNSCert,APNSKey)
}

func init() {
	
	APNSAddr =  beego.AppConfig.String("apns_url")
	if APNSAddr == ""{
			panic(" apns  url  not config" )
	}
	
	FeedbackAddr =  beego.AppConfig.String("feedback_url")
	if FeedbackAddr == ""{
			panic(" Feedback url  not config" )
	}

	APNSCert =  beego.AppConfig.String("apns_cert_path")
	if _, err := os.Stat(APNSCert); err != nil {
		if os.IsNotExist(err) {
			panic("apns cert file not find" + err.Error())
		}
	}
	
	APNSKey =  beego.AppConfig.String("apns_key_path")
	if _, err := os.Stat(APNSKey); err != nil {
		if os.IsNotExist(err) {
			panic("apns key file not find" + err.Error())
		}
	}
}