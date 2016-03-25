//@Description Redis服务
//@Contact czw@outlook.com

package libs

import (
	"github.com/astaxie/beego"
	"XYAPIServer/XYLibs"
)

var (
	//分享reids
	RedisDBShare *XYLibs.RedisHash
)



func init() {

	ipAccount := beego.AppConfig.String("share_redis_ip_list")	
	if ipAccount == "" {
		panic("分享服务redis未配置")
	}
	RedisDBShare = XYLibs.NewRedisHash()
	RedisDBShare.ConnectRedis(ipAccount)
	
	
}