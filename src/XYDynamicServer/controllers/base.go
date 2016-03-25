// base
package controllers

import (
	"github.com/astaxie/beego"
	"XYAPIServer/XYDynamicServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	//"sync"
)

var (
	uploadFileVoiceType = []string{"caf"}
	//baseLock = new(sync.Mutex)
)

type BaseController struct {
	beego.Controller
}



func GetLoginToken(uid uint32) string {
	//baseLock.Lock()
	//defer baseLock.Unlock()
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_USER_LOGIN_TOKEN,uid)
	redisDB := libs.RedisDBUser
	//println(noSQLKey)
	token , err := redisDB.Get(noSQLKey)
	if err != nil {
		beego.Error(err)
		return ""
	}
	if token != nil {
		return string(token.([]uint8))
	}else{
		return ""
	}
}
