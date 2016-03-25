//@Description 乡音认证
//@Contact czw@outlook.com

package XYLibs

import (
	"fmt"
	"errors"
	"strconv"
)

const (
	
	AUTH_HOME_VOICE_TYPE_SEND = 0 //认证别人
	
	AUTH_HOME_VOICE_TYPE_RECV = 1 //被人认证 
	
)


func NewVoiceAuthOK(rdb *RedisHash) *VoiceAuthOK {
	auth := new(VoiceAuthOK)
	auth.redisDB = rdb
	return auth
}

//认证战绩排行榜
type VoiceRecordRanking struct {
	
		AuthSendNum int //认证他人次数
		
		AuthRecvNum int //被认证次数
		
		AccuracyRate float32 //验证准确率
		
		Ranking int //认证排行
}

type VoiceAuthOK struct {

	  redisDB *RedisHash
	
	  UID uint32 //用户id

      AuthType int8 // 认证类型（0，认证别人；1，被认证）
	
	  AuthUID string // 认证我的人或被我认证uid
	 
	  AuthTime int //认证时间
	
	 
}



///设置已认证别人列表索引
func (u *VoiceAuthOK) SetSendIndex() (bool,error) {
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_SEND_VOICE_AUTH_LIST,u.UID)
		val := fmt.Sprintf("%s",u.AuthUID)
		_,err := u.redisDB.SADD(noSQLKey,val)
		if err != nil {
				return false,err
		}
		return true , nil
}

///获取已认证别人列表索引
func (u *VoiceAuthOK) GetSendIndex() (map[uint32]uint32,error) {
		res := make(map[uint32]uint32)
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_SEND_VOICE_AUTH_LIST,u.UID)
		data,err := u.redisDB.SMEMBERS(noSQLKey)
		if err != nil || data == nil{
				return res,err
		}
		for _,v := range data.([]interface{}) {
			uid,_ := strconv.ParseInt(string(v.([]uint8)),10,64)
			i := uint32(uid)
			res[i] = i
		}
		return res , nil
}

///设置被认证次数
func (u *VoiceAuthOK) SetRecvNum() (bool,error) {
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_RECV_VOICE_AUTH_NUM,u.UID)
		err := u.redisDB.INCR(noSQLKey)
		if err != nil {
				return false,err
		}
		return true , nil
}

//获取被认证次数
func (u *VoiceAuthOK) GetRecvNum() (int,error) {
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_RECV_VOICE_AUTH_NUM,u.UID)
		res := 0
		d,err := u.redisDB.Get(noSQLKey)
		if err != nil || d == nil {
				return res,err
		}
		if data,ok := d.([]uint8); ok {
			res,_ = strconv.Atoi(string(data))
		}else{
			//test
			fmt.Printf("%#v",d)
		}
		
		return res,nil
}

///设置认证错误次数
func (u *VoiceAuthOK) SetSendErrorNum() (bool,error) {
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_SEND_VOICE_AUTH_ERROR_NUM,u.UID)
		err := u.redisDB.INCR(noSQLKey)
		if err != nil {
				return false,err
		}
		return true , nil
}


///获取认证错误次数
func (u *VoiceAuthOK) GetSendErrorNum() (int,error) {
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_SEND_VOICE_AUTH_ERROR_NUM,u.UID)
		res,err := u.redisDB.Get(noSQLKey)
		d := 0
		if err != nil || res == nil {
				return d,err
		}
		switch res.(type) {
			case int64 :
				d = int(res.(int64))
			case []uint8 :
				d,_ = strconv.Atoi(string(res.([]uint8)))
		}
		
		return d  , nil
}

//是否认证过他
func (u *VoiceAuthOK) IsAuth() (int8,error) {
		noSQLKey := fmt.Sprintf("%s:%d",NO_SQL_USER_SEND_VOICE_AUTH_LIST,u.UID)
		val := fmt.Sprintf("%s",u.AuthUID)
		res,err := u.redisDB.SISMEMBER(noSQLKey,val)
		d := int8(0)
		if err != nil || res == nil {
				return d,err
		}
		if x,ok := res.(int64); ok {
			d = int8(x)
		}				
		return d,nil
}

//随机取没有认证过的用户(待优化)
func (u *VoiceAuthOK) RandomUID() []string {
	allKey := fmt.Sprintf("%s",NO_SQL_USER_ALREADY_RECORD_VOICE_LIST)
	isAuth := fmt.Sprintf("%s:%d",NO_SQL_USER_SEND_VOICE_AUTH_LIST,u.UID)
	res,_ := u.redisDB.SDIFF(allKey,isAuth)
	arrData := make([]string,0,9)
	if res != nil {
		if all,ok := res.([]interface{}); ok {
			for _,v := range all {
				if len(arrData) > 8 {	
					break
				}
				if vv,ok := v.([]uint8);ok{
					arrData = append(arrData,string(vv))
				}	
			}
		}
	}
	return arrData
}

//设置认证排名 （排名SCORE等于认证他人的次数）
func (u *VoiceAuthOK) SetRanking() (bool,error) {
		noSQLKey := fmt.Sprintf("%s",NO_SQL_USER_VOICE_AUTH_RANKING_LIST)
		val := fmt.Sprintf("%s",u.UID)
		res,err := u.redisDB.ZINCRBY(noSQLKey,"1",val)
		if err != nil  {
				return true,err
		}
		if res == nil {
			return false,nil
		}
		return true,errors.New("VoiceAuthOK->IsAuth error")
}

//获取认证排名
func (u *VoiceAuthOK) GetRanking() (int,error) {
		noSQLKey := fmt.Sprintf("%s",NO_SQL_USER_VOICE_AUTH_RANKING_LIST)
		val := fmt.Sprintf("%s",u.UID)
		res,err := u.redisDB.ZREVRANK(noSQLKey,val)
		rank := 0
		if err != nil {
				return rank,err
		}
		if res != nil {
			switch res.(type){
				case int64 :
					rank  = int(res.(int64))
				case []uint8 :
				  rank ,_ = strconv.Atoi(string(res.([]uint8)))
			}
			
			return rank + 1,nil
		}
		return rank,errors.New("VoiceAuthOK->GetRanking error")
}

//获取认证他人的次数（排名SCORE等于次数）
func (u *VoiceAuthOK) GetSendNum() (int,error) {
		noSQLKey := fmt.Sprintf("%s",NO_SQL_USER_VOICE_AUTH_RANKING_LIST)
		val := fmt.Sprintf("%s",u.UID)
		num := 0
		res,err := u.redisDB.ZSCORE(noSQLKey,val)
		if err != nil {
				return num,err
		}
		if res != nil {
			switch res.(type) {
				case int64 :
				num = int(res.(int64))
				case []uint8:
				num,_ = strconv.Atoi(string(res.([]uint8)))
			}
			
			return num,nil
		}
		return num,errors.New("VoiceAuthOK->GetSendNum error")
}

//获取战绩榜
func (u *VoiceAuthOK) GetRordRanking() (VoiceRecordRanking) {
		obj := VoiceRecordRanking{0,0,0,0}
		obj.Ranking, _ = u.GetRanking()		
		errnum,_ := u.GetSendErrorNum()	
		obj.AuthSendNum ,_ = u.GetSendNum()
		obj.AuthRecvNum,_ = u.GetRecvNum()
		count := errnum + obj.AuthSendNum
		
		if count > 0 {
			obj.AccuracyRate  =  float32(obj.AuthSendNum) / float32(count)
		}	
		return obj
}

