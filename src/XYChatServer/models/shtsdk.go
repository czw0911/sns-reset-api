// @Description 生活通im SDK的go版本.
// @Contact czw@outlook.com

package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"XYAPIServer/XYLibs"
	"github.com/astaxie/beego"
)

const (	
	SHT_REG  = "/add_user"
	SHT_UPDATE_TOKEN = "/upd_user_token"
	SHT_APP_TYPE = "5"
)

var (
 sht_im_server_ip string
 sht_app_type string
)

type SHTResponse struct {
	Retcode int  `json:"retcode"`
	Errormsg string `json:"errormsg"`
	Content interface{} `json:"content"`
}

type SHTIMServer struct {
	AppType string
	UID string
	Nickname string 
	UpdateType string 
}

//初始化SHTIMServer
func NewSHTIMServer() (server *SHTIMServer) {
	srv := new(SHTIMServer)
	srv.AppType = SHT_APP_TYPE
	srv.UpdateType = "1"
	return srv
}


func (sht *SHTIMServer) Reg()(SHTResponse,error) {
	site := fmt.Sprintf("%s%s",sht_im_server_ip,SHT_REG)
	var res SHTResponse
	p := url.Values{}
	p.Add("app_type",sht.AppType)
	p.Add("user_account",sht.UID)
	p.Add("user_name",sht.Nickname)
	url := fmt.Sprintf("%s?%s",site,p.Encode())
	d,err := XYLibs.HttpGet(url)
	if err != nil || len(d) == 0 {
		return res,err
	}
	
	err = json.Unmarshal(d,&res)
	if err != nil {
		return res,err
	}
	return res,nil
}


func  (sht *SHTIMServer) GetLoginToken()(SHTResponse,error) {
	site := fmt.Sprintf("%s%s",sht_im_server_ip,SHT_UPDATE_TOKEN)
	var res SHTResponse
	p := url.Values{}
	p.Add("app_type",sht.AppType)
	p.Add("user_account",sht.UID)
	p.Add("upd_cause",sht.UpdateType)
	url := fmt.Sprintf("%s?%s",site,p.Encode())
	d,err := XYLibs.HttpGet(url)
	if err != nil || len(d) == 0 {
		return res,err
	}
	
	err = json.Unmarshal(d,&res)
	if err != nil {
		return res,err
	}
	return res,nil
}

func init() {
	sht_im_server_ip = beego.AppConfig.String("sht_im_server_ip")
	if sht_im_server_ip == "" {
		panic("sht im server ip not config")
	}
	sht_app_type = beego.AppConfig.String("sht_app_type")
	if sht_app_type == "" {
		panic("sht app type not config")
	}
}