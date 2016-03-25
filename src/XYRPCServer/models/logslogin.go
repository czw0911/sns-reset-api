// 登陆日志
//@Contact czw@outlook.com

package models

import (
	"fmt"
	"time"
)

func NewLogsLogin() *LogsLogin {
	return new(LogsLogin)
}

type LogsLogin struct {
	
	tableName string //hash到的表名
	
	UID uint32 //用户id
	
	HomeProvinceID int //  家乡省id

	HomeCityID int //  家乡城市id

	HomeDistrictID int //  家乡区县id

	LivingProvinceID int //  居住地省id

	LivingCityID int //  居住地城市id

	LivingDistrictID int //  居住地区县id
	
	RegType int8 // 注册类型,1:手机 2：微博 3:微信
	
	LoginTime int64 //登陆时间
}

func (self *LogsLogin) initTableName() {
	 self.tableName =  fmt.Sprintf("login_logs_%s",time.Now().Format("200601"))
}

func (self *LogsLogin) Add(args *LogsLogin, reply *bool) error {
	
	self.initTableName()
	
	sqlStr := ` CREATE  TABLE  IF NOT EXISTS ` + self.tableName + `
				 (
					  ID bigint(20) unsigned NOT NULL AUTO_INCREMENT,
					  UID int(11) unsigned NOT NULL,
					  RegType tinyint(1) NOT NULL,
					  HomeProvinceID int(11) DEFAULT NULL,
					  HomeCityID int(11) DEFAULT NULL,
					  LivingProvinceID int(11) DEFAULT NULL,
					  LivingCityID int(11) DEFAULT NULL,
					  LoginTime int(11) unsigned NOT NULL,
					  PRIMARY KEY (ID),
					  UNIQUE KEY ID_UNIQUE (ID),
					  KEY LOGS_REG_INDEX (UID,RegType,HomeProvinceID,HomeCityID,LivingProvinceID,LivingCityID,LoginTime)
				) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='登陆日志';
	`
	
	conn,err := LogsDB.GetMaster()
	if err != nil {
		*reply = false
		return  err
	}
	_, err = conn.Raw(sqlStr).Exec()
	if err != nil {
		*reply = false
		return  err
	}
	
	_,err = conn.Raw("INSERT INTO "+ self.tableName +"(UID,RegType,HomeProvinceID,HomeCityID,LivingProvinceID,LivingCityID,LoginTime) VALUES(?,?,?,?,?,?,?)",
	args.UID,args.RegType,args.HomeProvinceID,args.HomeCityID,args.LivingProvinceID,args.LivingCityID,args.LoginTime).Exec()
	if err != nil {
		*reply = false
		return  err
	}
	*reply = true
	return nil
	
}


