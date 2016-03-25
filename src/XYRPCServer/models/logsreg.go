// 注册日志
//@Contact czw@outlook.com

package models

import (
	"fmt"
	"time"
)

func NewLogsReg() *LogsReg {
	return new(LogsReg)
}

type LogsReg struct {
	
	tableName string //hash到的表名
	
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

func (self *LogsReg) initTableName() {
	 self.tableName =  fmt.Sprintf("reg_logs_%s",time.Now().Format("2006"))
}

func (self *LogsReg) Add(args *LogsReg, reply *bool) error {
	
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
					  RegisterTime int(11) unsigned NOT NULL,
					  PRIMARY KEY (ID),
					  UNIQUE KEY ID_UNIQUE (ID),
					  KEY LOGS_REG_INDEX (UID,RegType,HomeProvinceID,HomeCityID,LivingProvinceID,LivingCityID,RegisterTime)
				) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='注册日志';
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
	
	_,err = conn.Raw("INSERT INTO "+ self.tableName +"(UID,RegType,HomeProvinceID,HomeCityID,LivingProvinceID,LivingCityID,RegisterTime) VALUES(?,?,?,?,?,?,?)",
	args.UID,args.RegType,args.HomeProvinceID,args.HomeCityID,args.LivingProvinceID,args.LivingCityID,args.RegisterTime).Exec()
	if err != nil {
		*reply = false
		return  err
	}
	*reply = true
	return nil
	
}


