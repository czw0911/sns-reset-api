//@Description 用户详细信息
//@Contact czw@outlook.com

package models

import (
	"XYAPIServer/XYRobot/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"strings"
)

var (
	userDetailInfoTableNameHash = consistent.New()
)

type UserDetailInfo struct {
	tableName string //hash到的表名

	dbConn orm.Ormer //hash到的数据库

	UID uint32 //用户id

	NickName string //用户昵称

	Avatar string //用户头像

	Thumbnail string //头像缩略图
	
	DiySign string //个性签名

	ProfessionID int //职业所属行业id

	JobID int //职业id

	Gender int //性别 1=暖男，2=女神

	Birthday int //生日

	TagID string //标签，多个用逗号分隔

	HomeProvinceID int //  家乡省id

	HomeCityID int //  家乡城市id

	HomeDistrictID int //  家乡区县id

	LivingProvinceID int //  居住地省id

	LivingCityID int //  居住地城市id

	LivingDistrictID int //  居住地区县id

	HomeVoice string // 乡音音频地址

	VoiceLen int // 乡音音频长度

	NorthLatitude int //  北纬

	EastLongtude int //  东经

	MaxID string //当前列表最大id

	PageType int8 //翻页类型 1,上翻；2，翻页
	
	PushID string //推送id
	
	PushType int8 //推送类型,1:ios 2：android


}

//按家乡省hash
func (u *UserDetailInfo) initDBAndTable() {

	u.dbConn = ConnMasterDBUser(uint32(u.HomeProvinceID))
	id := fmt.Sprintf("%d", u.HomeProvinceID)
	u.tableName, _ = userDetailInfoTableNameHash.Get(id)
}

//获取hash到的数据库
func (u *UserDetailInfo) GetHashDBName() string {
	u.initDBAndTable()
	return u.dbConn.Driver().Name()
}

//获取hash到的表名
func (u *UserDetailInfo) GetHashTableName() string {
	u.initDBAndTable()
	return u.tableName
}

func (u *UserDetailInfo) GeteDetailInfoUID() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  NickName, Avatar, Thumbnail, TagID, DiySign, Gender, Birthday,HomeProvinceID, HomeCityID, HomeDistrictID, LivingProvinceID, LivingCityID, LivingDistrictID, HomeVoice, VoiceLen ,ProfessionID ,JobID  FROM %s WHERE UID = ?  LIMIT 1", u.tableName)
	var res []orm.Params
	_, _ = u.dbConn.Raw(sqlStr, u.UID).Values(&res)
	return res
}

func (u *UserDetailInfo) Reg() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("REPLACE INTO %s(NickName,Avatar,Thumbnail,UID,HomeProvinceID,HomeCityID,HomeDistrictID,LivingProvinceID,LivingCityID,LivingDistrictID,HomeVoice,VoiceLen,Birthday,ProfessionID ,JobID) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.NickName, u.Avatar, u.Thumbnail, u.UID, u.HomeProvinceID, u.HomeCityID, u.HomeDistrictID, u.LivingProvinceID, u.LivingCityID, u.LivingDistrictID, u.HomeVoice, u.VoiceLen,u.Birthday,u.ProfessionID,u.JobID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserDetailInfo) Replace(columnName []string, columnVal []interface{}) (bool, error) {
	u.initDBAndTable()
	l := len(columnName)
	z := make([]string,0,l)
	for i := 0 ; i< l ; i++ {
		z = append(z,"?")
	}

	sqlStr := fmt.Sprintf("REPLACE INTO %s(%s) VALUES(%s)", u.tableName, strings.Join(columnName, " , "),strings.Join(z," , "))
	_, err := u.dbConn.Raw(sqlStr, columnVal).Exec()
	if err != nil {
		return false, err
	}
	return true, nil

}


func (u *UserDetailInfo) Edit(columnName []string, columnVal []interface{}) (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s SET %s WHERE UID = %d", u.tableName, strings.Join(columnName, " , "),u.UID)
	_, err := u.dbConn.Raw(sqlStr, columnVal).Exec()
	if err != nil {
		return false, err
	}
	return true, nil

}

func (u *UserDetailInfo) IsUIDExist() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID FROM %s WHERE UID = ? LIMIT 1", u.tableName)
	var res []orm.Params
	_, err := u.dbConn.Raw(sqlStr, u.UID).Values(&res)
	if err != nil {
		return true, err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

func (u *UserDetailInfo) DelUID() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("DELETE FROM  %s  WHERE UID = ?", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserDetailInfo) UpdateHomeVoice() (bool, error) {

	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET  HomeVoice = ? , VoiceLen = ?  WHERE UID = ?", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.HomeVoice, u.VoiceLen, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}




func (u *UserDetailInfo) UpdatePUSHID() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET PushID = ? , PushType = ? WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.PushID, u.PushType, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserDetailInfo) GetPUSHID()[]orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  PushID  FROM  %s  WHERE UID = ?  LIMIT 1", u.tableName)
	var res []orm.Params
	_, _ = u.dbConn.Raw(sqlStr, u.UID).Values(&res)
	return res
}

func (u *UserDetailInfo) SetCacheIOSPUSHID() (bool,error) {
		noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_USER_IOS_PUSHI_ID,u.UID)
		val := fmt.Sprintf("%s",u.PushID)
		_,err := libs.RedisDBUser.MSET(noSQLKey,val)
		if err != nil {
				return false,err
		}
		return true , nil
}



func (u *UserDetailInfo) GetMultiCacheIOSPUSHID(arrUID []string) ([]string,error) {
		arrKey := make([]interface{}, 0)
		for _, v := range arrUID {
			arrKey = append(arrKey, fmt.Sprintf("%s:%s", XYLibs.NO_SQL_USER_IOS_PUSHI_ID, v))
		}
		data, err :=libs.RedisDBUser.MGET(arrKey)
		if err != nil {
				return nil,err
		}
		arrResult := make([]string,0,len(arrUID))
		if d ,ok := data.([]interface {});ok {
			for _,v := range d {
				if id,ok := v.([]uint8);ok {
					arrResult = append(arrResult,string(id))
				}
			}
		}
		return arrResult , nil
}

//获取所有推送id
func (u *UserDetailInfo) GetAllPushID()[]orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  UID,PushID  FROM %s WHERE UID != ? AND PushID != '' ", u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}


//保存乡音文件到已录列表
func (u *UserDetailInfo) SaveHomeVoiceToRevordList() (bool, error) {

	noSQLKey := fmt.Sprintf("%s", XYLibs.NO_SQL_USER_ALREADY_RECORD_VOICE_LIST)
	_, err := libs.RedisDBUser.SADD(noSQLKey, fmt.Sprintf("%d", u.UID))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserDetailInfo) GenPageParam() string {
	arr := make([]string, 0, 4)
	if u.HomeProvinceID > 0 {
		arr = append(arr, fmt.Sprintf("HomeProvinceID = %d", u.HomeProvinceID))
	}

	if u.HomeCityID > 0 {
		arr = append(arr, fmt.Sprintf("HomeCityID = %d", u.HomeCityID))
	}

	if u.LivingProvinceID > 0 {
		arr = append(arr, fmt.Sprintf("LivingProvinceID = %d", u.LivingProvinceID))
	}

	if u.LivingCityID > 0 {
		arr = append(arr, fmt.Sprintf("LivingCityID = %d", u.LivingCityID))
	}
	res := ""
	if len(arr) > 0 {
		res = strings.Join(arr, " AND ")
	}
	return res
}

//最后的表
func (u *UserDetailInfo) GetLastTableName() error {
	return nil
}

//最前的表
func (u *UserDetailInfo) GetFirstTableName() error {
	return nil
}

//检查表是否在最前或最后
func (u *UserDetailInfo) CheckTableNameFirstOrLast() (bool, error) {
	return true, nil
}

//解析maxid
func (u *UserDetailInfo) ParseMaxID() bool {
	return false
}

//重新设置表
func (u *UserDetailInfo) ResetTableName() (bool, error) {
	return false, nil
}

//获取日期表年月
func (u *UserDetailInfo) GetYearAndMonth() string {
	return ""
}

//设置翻页类型
func (u *UserDetailInfo) SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *UserDetailInfo) SetMaxID(id string) {
	u.MaxID = id
}

//获取翻页类型
func (u *UserDetailInfo) GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *UserDetailInfo) GetMaxID() string {
	return u.MaxID
}

//上页
func (u *UserDetailInfo) PageUp() []orm.Params {
	u.initDBAndTable()
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " AND " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT  ID, UID  FROM %s WHERE  ID > ? %s  ORDER BY ID ASC LIMIT %d", u.tableName, strWhere, TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr, u.MaxID).Values(&table)
	return table

}

//下页
func (u *UserDetailInfo) PageDown() []orm.Params {
	u.initDBAndTable()

	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " AND " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT ID, UID  FROM %s WHERE ID < ? %s  ORDER BY ID DESC LIMIT %d", u.tableName, strWhere, TABLE_LIMIT_NUM)

	var table []orm.Params
	u.dbConn.Raw(sqlStr, u.MaxID).Values(&table)
	return table
}

//最后页
func (u *UserDetailInfo) PageEnd() []orm.Params {
	u.initDBAndTable()
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " WHERE " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT ID, UID  FROM %s  %s  ORDER BY ID ASC LIMIT %d", u.tableName, strWhere, TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//首页
func (u *UserDetailInfo) PageFirst() []orm.Params {
	u.initDBAndTable()
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " WHERE " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT ID, UID  FROM %s  %s   ORDER BY ID Desc LIMIT %d", u.tableName, strWhere, TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//获取1万个用户
func (u *UserDetailInfo) GetLimitTenThousand() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  UID  FROM %s WHERE UID != ? ORDER BY ID ASC LIMIT 10000", u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}

func init() {
	num := UserDB.GetHashTableNum()
	for i := 0; i < num; i++ {
		tn := fmt.Sprintf("user_detail_info_%d", i)
		userDetailInfoTableNameHash.Add(tn)
	}
}
