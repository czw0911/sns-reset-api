//@Description 用户头像
//@Contact czw@outlook.com

package XYLibs

import (
	"encoding/json"
	"fmt"
)

func NewUserAvatar(rdb *RedisHash) *UserAvatar {
	user := new(UserAvatar)
	user.redisDB = rdb
	return user
}

type UserAvatar struct {
	redisDB *RedisHash

	UID uint32 //用户id

	NickName string //用户昵称

	Avatar string //用户头像

	Thumbnail string //头像缩略图

	HomeProvinceID int //  家乡省id

	HomeCityID int //  家乡城市id

	HomeDistrictID int //  家乡区县id

	LivingProvinceID int //  居住地省id

	LivingCityID int //  居住地城市id

	LivingDistrictID int //  居住地区县id

	HomeVoice string // 乡音音频地址

	VoiceLen int // 乡音音频长度

	ProfessionID int //职业所属行业id

	JobID int //职业id

	IsFollow int8 //是否关注0:未关注；1：已关注；2:被关注;3：互相关注

	Birthday int //生日

	Gender int //性别

	IsMember int8 // 是否会员

	IsGuess int8 //是否猜测过乡音

	AuthRecvNum int // 获得乡音认证次数

	TagsID []string //标签信息id

	LastLoginTime int64 //最后登录时间

}

func (u *UserAvatar) SetRedisConnect(rdb *RedisHash) {
	u.redisDB = rdb
}

//获取ios推送id
func (u *UserAvatar) GetCacheIOSPUSHID(uid string) (string,error) {
		noSQLKey := fmt.Sprintf("%s:%s",NO_SQL_USER_IOS_PUSHI_ID,uid)
		data, err := u.redisDB.Get(noSQLKey)
		res := ""
		if err != nil {
				return res,err
		}
		if id,ok := data.([]uint8);ok {
			res = string(id)
		}
		return res , nil
}

///设置注册用户uid索引
func (u *UserAvatar) SetAllRegisterUID() (bool,error) {
		noSQLKey := fmt.Sprintf("%s",NO_SQL_USER_ALL_REGISTER_UID)
		val := fmt.Sprintf("%d",u.UID)
		_,err := u.redisDB.SADD(noSQLKey,val)
		if err != nil {
				return false,err
		}
		return true , nil
}

//随机获取用户
func (u *UserAvatar) GetRandomUID() ([]string,error) {
		noSQLKey := fmt.Sprintf("%s",NO_SQL_USER_ALL_REGISTER_UID)
		res := make([]string,0,3)
		data,err := u.redisDB.SRANDMEMBER(noSQLKey,"3")
		if err != nil {
				return res,err
		}
		//fmt.Printf("%#v\n",data)
		if d,ok := data.([]interface{});ok{
			
			for _,v := range d {
				res = append(res,string(v.([]uint8)))
			}
		}
		return res , nil
}

func (u *UserAvatar) Set() (bool, error) {
	noSQLKey := fmt.Sprintf("%s:%d", NO_SQL_USER_AVATER_INFO, u.UID)

	data, err := json.Marshal(u)
	if err != nil {
		return false, err
	}
	_, err = u.redisDB.MSETByte(noSQLKey, data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserAvatar) Get(uid uint32) (UserAvatar, error) {

	var avatar UserAvatar
	key := fmt.Sprintf("%s:%d", NO_SQL_USER_AVATER_INFO, uid)
	data, err := u.redisDB.Get(key)
	if err != nil || data == nil {
		return avatar, err
	}
	if d,ok := data.([]uint8);ok {
		json.Unmarshal(d, &avatar)
	}else{
		//test
		fmt.Printf("get UserAvatar error:%#v\n",data)
	}
	return avatar, nil
}

func (u *UserAvatar) GetAll(arrUID []string, fileServerAddr string) (bool, map[uint32]UserAvatar, error) {
	arrKey := make([]interface{}, 0)
	for _, v := range arrUID {
		arrKey = append(arrKey, fmt.Sprintf("%s:%s", NO_SQL_USER_AVATER_INFO, v))
	}
	data, err := u.redisDB.MGET(arrKey)
	if err != nil || data == nil {
		return false, nil, err
	}
	size := len(data.([]interface{}))
	arrRes := make(map[uint32]UserAvatar, size)
	objVoiceAuthOK := NewVoiceAuthOK(u.redisDB)
	objVoiceAuthOK.UID = u.UID
	arrIsAuth, _ := objVoiceAuthOK.GetSendIndex()
	objFollow := NewFollow(u.redisDB)
	objFollow.UID = u.UID
	l := len(arrIsAuth)
	for _, j := range data.([]interface{}) {

		if j == nil {
			continue
		}
		var user UserAvatar
		json.Unmarshal(j.([]uint8), &user)

		objVoiceAuthOK.UID = user.UID
		user.IsGuess = 0
		if l > 0 {
			if _, ok := arrIsAuth[user.UID]; ok {
				user.IsGuess = 1
			}
		}
		user.AuthRecvNum, _ = objVoiceAuthOK.GetRecvNum()

		if user.HomeVoice != "" {
			user.HomeVoice = fmt.Sprintf("%s=%s", fileServerAddr, user.HomeVoice)
		}

		if user.Avatar != "" {
			user.Avatar = fmt.Sprintf("%s=%s", fileServerAddr, user.Avatar)
		}

		if user.Thumbnail != "" {
			user.Thumbnail = fmt.Sprintf("%s=%s", fileServerAddr, user.Thumbnail)
		}
		objFollow.FollowUID = user.UID
		objFollow.GetFollowState()
		user.IsFollow = objFollow.FollowState
		arrRes[user.UID] = user
	}

	return true, arrRes, nil
}
