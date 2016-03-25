//@Description 系统消息日志(乡音团队)
//@Contact czw@outlook.com

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs"
)



type SysMsgLog struct {
	tableName string //hash到的表名

	dbConn orm.Ormer //hash到的数据库
	
	MaxID string //当前列表最大id  
	
	PageType int8 //翻页类型 1,上翻；2，翻页
}

func (u *SysMsgLog) initDBAndTable() {
	u.dbConn = ConnCommonDB()
	u.tableName = "system_msg"
}



	//最后的表
func (u *SysMsgLog) GetLastTableName() error{
	return nil
}
	
	//最前的表
func (u *SysMsgLog)	GetFirstTableName() error{
	return nil
}
	
	//检查表是否在最前或最后
func (u *SysMsgLog) CheckTableNameFirstOrLast() (bool,error){
	return true,nil
}
	
	//解析maxid
func (u *SysMsgLog) ParseMaxID() bool{
	return false
}
	
	//重新设置表
func (u *SysMsgLog)	ResetTableName() (bool,error){
	return false,nil
}
	
//获取日期表年月
func (u *SysMsgLog)	GetYearAndMonth() string {
	return ""
}

//设置翻页类型
func (u *SysMsgLog)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *SysMsgLog)	SetMaxID(id string) {
	u.MaxID = id
}

//获取翻页类型
func (u *SysMsgLog)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *SysMsgLog)	GetMaxID() string {
	return u.MaxID
}

//上页
func (u *SysMsgLog) PageUp() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID, Messages, PostTime  FROM %s WHERE ID > ?  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID).Values(&table)
	return table
	
}

//下页
func (u *SysMsgLog) PageDown() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID, Messages, PostTime  FROM %s WHERE ID < ?  ORDER BY ID DESC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)

	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID).Values(&table)
	return table
}
//最后页
func (u *SysMsgLog) PageEnd() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID, Messages, PostTime  FROM %s  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//首页
func (u *SysMsgLog) PageFirst() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID, Messages, PostTime  FROM  %s  ORDER BY ID Desc LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}


//最新的1条
func (u *SysMsgLog) NewsOne() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  ID, Messages, PostTime  FROM  %s  ORDER BY ID Desc LIMIT 1",u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//根据时间获取未读消息数

func (u *SysMsgLog) GetUnReadNUM(lastTime string) []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT  COUNT(ID) UnreadNum  FROM  %s  WHERE PostTime > ? ",u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,lastTime).Values(&table)
	return table
}





