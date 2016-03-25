//@Description 求乡音认证
//@Contact czw@outlook.com

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"XYAPIServer/XYLibs"
)

var (
	voiceAuthRequestTableNameHash = consistent.New()
)

type VoiceAuthRequest struct {
	tableName string //hash到的表名

	dbConn orm.Ormer //hash到的数据库

	UID uint32 //用户id
	
	ReuestUID uint32 // 求认证uid	 
	
	MaxID string //当前列表最大id  
	
	PageType int8 //翻页类型 1,上翻；2，翻页
}

func (u *VoiceAuthRequest) initDBAndTable() {
	u.dbConn = ConnMasterDB(u.UID)
	id := fmt.Sprintf("%d", u.UID)
	u.tableName, _ = voiceAuthRequestTableNameHash.Get(id)
}


func (u *VoiceAuthRequest) Add() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("REPLACE INTO %s(UID,ReuestUID) VALUES(?,?)", u.tableName)
	_, err := u.dbConn.Raw(sqlStr,  u.UID, u.ReuestUID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *VoiceAuthRequest) Del() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("DELETE FROM %s  WHERE UID = ? AND ReuestUID = ?", u.tableName)
	_, err := u.dbConn.Raw(sqlStr,  u.UID, u.ReuestUID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}



	//最后的表
func (u *VoiceAuthRequest) GetLastTableName() error{
	return nil
}
	
	//最前的表
func (u *VoiceAuthRequest)	GetFirstTableName() error{
	return nil
}
	
	//检查表是否在最前或最后
func (u *VoiceAuthRequest) CheckTableNameFirstOrLast() (bool,error){
	return true,nil
}
	
	//解析maxid
func (u *VoiceAuthRequest) ParseMaxID() bool{
	return false
}
	
	//重新设置表
func (u *VoiceAuthRequest)	ResetTableName() (bool,error){
	return false,nil
}
	
//获取日期表年月
func (u *VoiceAuthRequest)	GetYearAndMonth() string {
	return ""
}

//设置翻页类型
func (u *VoiceAuthRequest)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *VoiceAuthRequest)	SetMaxID(id string) {
	u.MaxID = id
}

//获取翻页类型
func (u *VoiceAuthRequest)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *VoiceAuthRequest)	GetMaxID() string {
	return u.MaxID
}

//上页
func (u *VoiceAuthRequest) PageUp() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID,ReuestUID  FROM %s WHERE ID > ? AND UID = ?  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID,u.UID).Values(&table)
	return table
	
}

//下页
func (u *VoiceAuthRequest) PageDown() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,ReuestUID  FROM %s WHERE ID < ?  AND UID = ?  ORDER BY ID DESC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)

	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID,u.UID).Values(&table)
	return table
}
//最后页
func (u *VoiceAuthRequest) PageEnd() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,ReuestUID  FROM %s  WHERE UID = ?  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}

//首页
func (u *VoiceAuthRequest) PageFirst() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,ReuestUID  FROM %s  WHERE UID = ?    ORDER BY ID Desc LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}






func init() {
	for i := 0; i < USER_HASH_TABLE_NUM; i++ {
		tn := fmt.Sprintf("voice_auth_request_%d", i)
		voiceAuthRequestTableNameHash.Add(tn)
	}
}
