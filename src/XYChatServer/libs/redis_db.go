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
)



func init() {

	ipAccount := beego.AppConfig.String("account_redis_ip_list")	
	if ipAccount == "" {
		panic("账号服务redis未配置")
	}
	RedisDBUser = XYLibs.NewRedisHash()
	RedisDBUser.ConnectRedis(ipAccount)
	
	
}