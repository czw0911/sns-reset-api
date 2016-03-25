//@Description 日志
//@Contact czw@outlook.com


package models

import (
	"errors"
	"net/rpc"
	"net/rpc/jsonrpc"
	"github.com/astaxie/beego"
)


//注册日志
type LogsReg struct {
	
	
	UID uint32 //用户id
	
	HomeProvinceID int //  家乡省id

	HomeCityID int //  家乡城市id

	HomeDistrictID int //  家乡区县id

	LivingProvinceID int //  居住地省id

	LivingCityID int //  居住地城市id

	LivingDistrictID int //  居住地区县id
	
	RegType int8 // 注册类型,1:手机 2：微博 3:微信
	
	RegisterTime int64 //注册时间
}

//登陆日志
type LogsLogin struct {
	
	
	UID uint32 //用户id
	
	HomeProvinceID int //  家乡省id

	HomeCityID int //  家乡城市id

	HomeDistrictID int //  家乡区县id

	LivingProvinceID int //  居住地省id

	LivingCityID int //  居住地城市id

	LivingDistrictID int //  居住地区县id
	
	RegType int8 // 注册类型,1:手机 2：微博 3:微信
	
	LoginTime int64 //时间
}

type Logs int

func (self *Logs) connectRPC() (*rpc.Client, error) {
	ip :=  beego.AppConfig.String("rpc_server_ip")
	return jsonrpc.Dial("tcp",ip)
}

func (self *Logs) AddLogsReg(args *LogsReg) error {
	
	 client,err := self.connectRPC()
	 if err != nil {
			return err
	 }
	 defer client.Close()
	
	 var result bool
	 err = client.Call("LogsReg.Add",args,&result)
	 if err != nil {
		return err
	 }
	 if !result {
		return errors.New("call LogsReg.Add fail")
	}
	return nil	 
}

func (self *Logs) AddLogsLogin(args *LogsLogin) error {
	
	 client,err := self.connectRPC()
	 if err != nil {
			return err
	 }
	 defer client.Close()
	
	 var result bool
	 err = client.Call("LogsLogin.Add",args,&result)
	 if err != nil {
		return err
	 }
	 if !result {
		return errors.New("call LogsLogin.Add fail")
	}
	return nil	 
}