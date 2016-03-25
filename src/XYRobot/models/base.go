// mysql存储
//@Contact czw@outlook.com

package models


import (
	"github.com/astaxie/beego/orm"
    "XYAPIServer/XYLibs/storage"
	"XYAPIServer/XYLibs/storage/mysql"
)

const TABLE_LIMIT_NUM = 10
var (
	//通用库
	CommonDB storage.Storage
	//动态库
	DynamicDB storage.Storage
	//账号库
	UserDB storage.Storage
	//活动库
	ActivityDB storage.Storage
	//日志库
	LogsDB storage.Storage
)



func ConnSlaveDBDynamic(HomeProvinceID int) orm.Ormer {
	DynamicDB.SetHashVal(HomeProvinceID)
	d ,err := DynamicDB.GetMaster()
	if err != nil {
		panic("DynamicDB master database get fail")
	}
	return d
}

func ConnMasterDBDynamic(HomeProvinceID int) orm.Ormer {
	DynamicDB.SetHashVal(HomeProvinceID)
	d ,err := DynamicDB.GetSalve()
	if err != nil {
		panic("DynamicDB salve database get fail")
	}
	return d
}

func ConnMasterDBActivity(ActivityID int64) orm.Ormer {
	ActivityDB.SetHashVal(ActivityID)
	d ,err := ActivityDB.GetMaster()
	if err != nil {
		panic("ActivityDB master database get fail")
	}
	return d
}

func ConnSlaveDBActivity(ActivityID int64) orm.Ormer {
	ActivityDB.SetHashVal(ActivityID)
	d ,err := ActivityDB.GetSalve()
	if err != nil {
		panic("ActivityDB salve database get fail")
	}
	return d
}


func ConnMasterDBUser(uid uint32) orm.Ormer {
	UserDB.SetHashVal(uid)
	d ,err := UserDB.GetMaster()
	if err != nil {
		panic("UserDB database get fail")
	}
	return d
}

func ConnCommonDB() orm.Ormer {
	d ,err := CommonDB.GetMaster()
	if err != nil {
		panic("CommonDB database get fail")
	}
	return d
}

func ConnLogsDB() orm.Ormer {
	d ,err := LogsDB.GetMaster()
	if err != nil {
		panic("LogsDB database get fail")
	}
	return d
}

func init() {
	//通用库
	storage.Register("xy_db_common",&mysql.Connect{
		"mysql_common.json",
		"xy_db_common_master",
		"xy_db_common_slave",
		mysql.NewHash(),
		mysql.NewHash(),
		"id",
		1,

	})
	var err error
	CommonDB,err = storage.NewStorage("xy_db_common")
	if err != nil {
		panic("xy_db_common database init fail")
	}
	
	//动态
	storage.Register("xy_db_dynamic",&mysql.Connect{
		"mysql_dynamic.json",
		"xy_db_dynamic_master",
		"xy_db_dynamic_slave",
		mysql.NewHash(),
		mysql.NewHash(),
		"id",
		1,
	})
	DynamicDB,err = storage.NewStorage("xy_db_dynamic")
	if err != nil {
		panic("xy_db_dynamic database init fail")
	}
	
	//用户账号
	storage.Register("xy_db_user",&mysql.ConnectOld{
		"mysql_user.json",
		"xy_db_user_master",
		"xy_db_user_slave",
		mysql.NewHash(),
		mysql.NewHash(),
		"id",
		1,
	})
	UserDB,err = storage.NewStorage("xy_db_user")
	if err != nil {
		panic("xy_db_user database init fail")
	}
	
	//活动
	storage.Register("xy_db_activity",&mysql.ConnectOld{
		"mysql_activity.json",
		"xy_db_activity_master",
		"xy_db_activity_slave",
		mysql.NewHash(),
		mysql.NewHash(),
		"id",
		1,
	})
	ActivityDB,err = storage.NewStorage("xy_db_activity")
	if err != nil {
		panic("xy_db_activity database init fail")
	}
	
	//log
	
	storage.Register("xy_db_logs",&mysql.Connect{
		"mysql_logs.json",
		"xy_db_logs_master",
		"xy_db_logs_slave",
		mysql.NewHash(),
		mysql.NewHash(),
		"id",
		1,
	})
	LogsDB,err = storage.NewStorage("xy_db_logs")
	if err != nil {
		panic("xy_db_logs database init fail")
	}
}