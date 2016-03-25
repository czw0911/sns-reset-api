//关注
package XYLibs

import (
	"fmt"
	"time"
	"strconv"
)

func NewFollow(rdb *RedisHash) *Follow {
	f := new(Follow)
	f.redisDB = rdb
	return f
}

type Follow struct {
	
	  redisDB *RedisHash
	
	  UID uint32 //我的id
	
	  FollowType int8 // 0:我关注 1：关注我
	
	  FollowUID uint32    //我关注或关注我的uid 
	
	  FollowState int8  // 关注状态（0:未关注；1：已关注；2:被关注;3：互相关注
	
	  MaxID int //当前列表最大id
	
	  PageType int8 //翻页类型 1,上翻；2，翻页
	

}

//设置我关注
func (u *Follow) SetFollowYou()(bool,error){
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_YOU,u.UID)
		score := fmt.Sprintf("%d",time.Now().Unix())
		val := fmt.Sprintf("%d",u.FollowUID)
		_,err := u.redisDB.ZADD(noSQLKey,score,val)
		if err != nil {
				return false,err
		}
		return true , nil
}
//删除我关注
func (u *Follow) RemFollowYou()(bool,error){
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_YOU,u.UID)
		val := fmt.Sprintf("%d",u.FollowUID)
		_,err := u.redisDB.ZREM(noSQLKey,val)
		if err != nil {
				return false,err
		}
		return true , nil
}

//设置被关注
func (u *Follow) SetFollowMe()(bool,error){
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_ME,u.UID)
		score := fmt.Sprintf("%d",time.Now().Unix())
		val := fmt.Sprintf("%d",u.FollowUID)
		_,err := u.redisDB.ZADD(noSQLKey,score,val)
		if err != nil {
				return false,err
		}
		return true , nil
}

//删除被关注
func (u *Follow) RemFollowMe()(bool,error){
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_ME,u.UID)
		val := fmt.Sprintf("%d",u.FollowUID)
		_,err := u.redisDB.ZREM(noSQLKey,val)
		if err != nil {
				return false,err
		}
		return true , nil
}

//获取关注状态
func (u *Follow) GetFollowState() {
	 u.FollowState = 0
	 followYou := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_YOU,u.UID)
	 followMe :=  fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_ME,u.UID)
	 val := fmt.Sprintf("%d",u.FollowUID)
	 res,err := u.redisDB.ZSCORE(followYou,val)
	 fy := int8(0) 
	 if err == nil && res != nil{
				fy = 1
	 }
	
	 res,err = u.redisDB.ZSCORE(followMe,val)
	 fm := int8(0) 
	 if err == nil && res != nil{
				fm = 2
	 }
	 u.FollowState = fy + fm
}


func (u *Follow) GetNoSQLKey() string {
	
	noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_YOU,u.UID)
	if u.FollowType != 0 {
		noSQLKey = fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_ME,u.UID)
	}
	return noSQLKey
}

//关注我的所有用户
func (u *Follow) GetAllFollowMe() ([]string,error){
	var res []string
	noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_FOLLOW_ME,u.UID)
	data,err := u.redisDB.ZREVRANGE(noSQLKey,"0","-1")
	if err != nil {
		return res,err
	}
	if arr,ok := data.([]interface{});ok{
		res = make([]string,0,len(arr))
		for _,v := range arr {
			if id ,ok := v.([]uint8);ok {
				res = append(res,string(id))
			}
		}
	}
	return res,nil
}

//首页
func (u *Follow) PageFirst() (map[string]interface{},error) {
	res := make(map[string]interface{},3)
	res["MaxID"] = "0"
	res["MinID"] = strconv.Itoa(TABLE_LIMIT_NUM - 1)
	
	list := make([]string,0,TABLE_LIMIT_NUM)
	key := u.GetNoSQLKey()
    data,err := u.redisDB.ZREVRANGE(key,res["MaxID"].(string),res["MinID"].(string))
	if err != nil {
				return res,err
	}
	if arr,ok := data.([]interface{});ok{
		for _,v := range arr {
			list = append(list,string(v.([]uint8)))
		}
	}
	res["List"] =  list
	return res,nil
}

func (u *Follow) PageDown()  (map[string]interface{},error) {
	res := make(map[string]interface{},3)
	start := u.MaxID + 1
	stop := start + TABLE_LIMIT_NUM - 1
	res["MaxID"] = strconv.Itoa(start)
	res["MinID"] = strconv.Itoa(stop)
	list := make([]string,0,TABLE_LIMIT_NUM)
	key := u.GetNoSQLKey()
    data,err := u.redisDB.ZREVRANGE(key,res["MaxID"].(string),res["MinID"].(string))
	if err != nil {
				return res,err
	}
	if arr,ok := data.([]interface{});ok{
		for _,v := range arr {
			list = append(list,string(v.([]uint8)))
		}
	}
	res["List"] =  list
	return res,nil
}

func (u *Follow) PageUp() (map[string]interface{},error) {
	res := make(map[string]interface{},3)
	start := u.MaxID - TABLE_LIMIT_NUM
	if start < 0 {
		start = 0
	}
	stop := start + TABLE_LIMIT_NUM - 1
	res["MaxID"] = strconv.Itoa(start)
	res["MinID"] = strconv.Itoa(stop)
	list := make([]string,0,TABLE_LIMIT_NUM)
	key := u.GetNoSQLKey()
    data,err := u.redisDB.ZREVRANGE(key,res["MaxID"].(string),res["MinID"].(string))
	if err != nil {
				return res,err
	}
	if arr,ok := data.([]interface{});ok{
		for _,v := range arr {
			list = append(list,string(v.([]uint8)))
		}
	}
	res["List"] =  list
	return res,nil
}

func (u *Follow) PageEnd() (map[string]interface{},error) {
	key := u.GetNoSQLKey()
	res := make(map[string]interface{},3)
	count,err := u.redisDB.ZCARD(key)
	if err != nil {
		return res,err
	}
	end,ok := count.(int64)
	if !ok {	
		return res,err
	}
	stop := int(end)
	start := stop - TABLE_LIMIT_NUM
	if start < 0 {
		start = 0
	}
	res["MaxID"] = strconv.Itoa(start)
	res["MinID"] = strconv.Itoa(stop)
	list := make([]string,0,TABLE_LIMIT_NUM)
	
    data,err := u.redisDB.ZREVRANGE(key,res["MaxID"].(string),res["MinID"].(string))
	if err != nil {
				return res,err
	}
	if arr,ok := data.([]interface{});ok{
		for _,v := range arr {
			list = append(list,string(v.([]uint8)))
		}
	}
	res["List"] =  list
	return res,nil
}

func (u *Follow) Pageing() (map[string]interface{},error) {
	res := make(map[string]interface{},3)
	res["MaxID"] = ""
	res["MinID"] = ""
	res["List"] = make([]string,0,TABLE_LIMIT_NUM)
	var err error
	if u.PageType == PAGE_TYPE_UP {
		//上翻页
		if u.MaxID == 0 {
			//首页
			res,err = u.PageFirst()
		}else{
			//上一页
			res,err = u.PageUp()
		}
		
	}else{
		//下翻页
		if u.MaxID == 0 {
			//尾页
			res,err = u.PageEnd()
		}else{
			//下一页
			res,err = u.PageDown()
		}
	}
	return res,err
	
}