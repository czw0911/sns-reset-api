// base
package controllers

import (
	"github.com/astaxie/beego"
	"XYAPIServer/XYRobot/libs"
	"fmt"
	"XYAPIServer/XYLibs"
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
	
	token , err := libs.RedisDBUser.Get(noSQLKey)
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
