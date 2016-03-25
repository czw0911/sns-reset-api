//@Description Redis服务
//@Contact czw@outlook.com

package libs

import (
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
)

var (
	//用户账号redis
	RedisDBUser *XYLibs.RedisHash
	
	//活动reids
	RedisDBActivity *XYLibs.RedisHash
)




func init() {

	ipAccount := beego.AppConfig.String("account_redis_ip_list")	
	if ipAccount == "" {
		panic("账号服务redis未配置")
	}
	RedisDBUser = XYLibs.NewRedisHash()
	RedisDBUser.HashID = "account"
	RedisDBUser.ConnectRedis(ipAccount)
	
	
	ipActivity := beego.AppConfig.String("activity_redis_ip_list")
	if ipActivity == "" {
		panic("活动服务redis未配置")
	}
	RedisDBActivity = XYLibs.NewRedisHash()
	RedisDBActivity.ConnectRedis(ipActivity)
	
	
}