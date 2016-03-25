// 存储
//@Contact czw@outlook.com

package models

import (
	"XYAPIServer/XYRPCServer/libs/storage"
	_ "XYAPIServer/XYRPCServer/libs/storage/mysql"

)

var (
	//通用库
	CommonDB storage.Storage
	//日志库
	LogsDB storage.Storage
)

func init() {
	var err error
	CommonDB,err = storage.NewStorage("xy_db_common")
	if err != nil {
		panic("xy_db_common database init fail")
	}
	
	LogsDB,err = storage.NewStorage("xy_db_logs")
	if err != nil {
		panic("xy_db_logs database init fail")
	}
}