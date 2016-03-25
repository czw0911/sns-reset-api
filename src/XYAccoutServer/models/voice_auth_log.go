//@Description 乡音认证日志
//@Contact czw@outlook.com

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"XYAPIServer/XYLibs"
)

var (
	voiceAuthLogTableNameHash = consistent.New()
)

type VoiceAuthLog struct {
	tableName string //hash到的表名

	dbConn orm.Ormer //hash到的数据库

	UID uint32 //用户id

    AuthType int8 // 认证类型（0，认证别人；1，被认证）
	
	AuthUID string // 认证我的人或被我认证uid
	 
	AuthTime int64 //认证时间
	
	MaxID string //当前列表最大id  
	
	PageType int8 //翻页类型 1,上翻；2，翻页
}

func (u *VoiceAuthLog) initDBAndTable() {
	u.dbConn = ConnMasterDB(u.UID)
	id := fmt.Sprintf("%d", u.UID)
	u.tableName, _ = voiceAuthLogTableNameHash.Get(id)
}


func (u *VoiceAuthLog) Add() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("REPLACE INTO %s(UID,AuthType,AuthUID,AuthTime) VALUES(?,?,?,?)", u.tableName)
	_, err := u.dbConn.Raw(sqlStr,  u.UID, u.AuthType, u.AuthUID, u.AuthTime).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}



	//最后的表
func (u *VoiceAuthLog) GetLastTableName() error{
	return nil
}
	
	//最前的表
func (u *VoiceAuthLog)	GetFirstTableName() error{
	return nil
}
	
	//检查表是否在最前或最后
func (u *VoiceAuthLog) CheckTableNameFirstOrLast() (bool,error){
	return true,nil
}
	
	//解析maxid
func (u *VoiceAuthLog) ParseMaxID() bool{
	return false
}
	
	//重新设置表
func (u *VoiceAuthLog)	ResetTableName() (bool,error){
	return false,nil
}
	
//获取日期表年月
func (u *VoiceAuthLog)	GetYearAndMonth() string {
	return ""
}

//设置翻页类型
func (u *VoiceAuthLog)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *VoiceAuthLog)	SetMaxID(id string) {
	u.MaxID = id
}

//获取翻页类型
func (u *VoiceAuthLog)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *VoiceAuthLog)	GetMaxID() string {
	return u.MaxID
}

//上页
func (u *VoiceAuthLog) PageUp() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID,AuthType,AuthUID,AuthTime  FROM %s WHERE ID > ? AND UID = ?  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID,u.UID).Values(&table)
	return table
	
}

//下页
func (u *VoiceAuthLog) PageDown() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,AuthType,AuthUID,AuthTime  FROM %s WHERE ID < ?  AND UID = ?  ORDER BY ID DESC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)

	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID,u.UID).Values(&table)
	return table
}
//最后页
func (u *VoiceAuthLog) PageEnd() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,AuthType,AuthUID,AuthTime  FROM %s  WHERE UID = ?  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}

//首页
func (u *VoiceAuthLog) PageFirst() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,AuthType,AuthUID,AuthTime  FROM %s  WHERE UID = ?    ORDER BY ID Desc LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}






func init() {
	for i := 0; i < USER_HASH_TABLE_NUM; i++ {
		tn := fmt.Sprintf("voice_auth_log_%d", i)
		voiceAuthLogTableNameHash.Add(tn)
	}
}
