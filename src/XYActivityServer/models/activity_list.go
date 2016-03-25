//活动主题列表
package models

import (
	//"encoding/json"
	"fmt"
	"XYAPIServer/XYLibs"
	"XYAPIServer/XYActivityServer/libs"
	"github.com/astaxie/beego/orm"
	
	
)

const tablename_activity_list_column = "ID, ActivityID, ActivityName, ActivityDesImg,Url,ShareImgUrl,ShareContent"

func  NewActivityList() *ActivityList {
	return &ActivityList{tableName:"activity_list"}
}

type ActivityList struct {

	tableName string //"activity_theme_list"
	
	dbConn  orm.Ormer
	
	ActivityID  int64 //活动编号
	
	ActivityName  string //活动名称
	
	ActivityDesImg  string //活动描述
	
	JoinNum int //参加人数
	
	MaxID string //当前列表最大id  
	
	PageType int8 //翻页类型 1,上翻；2，翻页
	
	ShareImgUrl string //分享URL
	
	ShareContent string //分享内容
}



func (u *ActivityList) IsExist()(bool,error){
	db := ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT ID  FROM %s WHERE ActivityID = ? LIMIT 1",u.tableName)
	var res []orm.Params
	_, err := db.Raw(sqlStr,u.ActivityID).Values(&res)
	if err != nil {
		return false, err
	}
	
	if len(res) > 0 {
		return true,nil
	}
	return false, nil
}

//设置活动参与人数
func (u *ActivityList) SetJoinNum() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_JOIN_NUM,u.ActivityID)

	err := libs.RedisDBActivity.INCR(noSQLKey)
	if err != nil {
			return false,err
	}
	return true , nil
}



//获取活动参与人数
func (u *ActivityList) GetAllJoinNum(arrActivityID []string) (bool,[]string,error) {
	
	arrKey := make([]interface{},0)
	for _,v := range arrActivityID {
		arrKey = append(arrKey,fmt.Sprintf("%s:%s",XYLibs.NO_SQL_ACTIVITY_JOIN_NUM,v))
	}
	data,err := libs.RedisDBActivity.MGET(arrKey)
	if err != nil {
			return false,nil,err
	}
	size := len(data.([]interface {}))
	arrRes := make([]string,0,size)
	for _,j := range data.([]interface {}) {
		if j != nil {	
			arrRes = append(arrRes,string(j.([]byte)))		
		}else{
			arrRes = append(arrRes,"0")
		}
		
	}
	
	return true , arrRes, nil
}

	//最后的表
func (u *ActivityList) GetLastTableName() error{
	return nil
}
	
	//最前的表
func (u *ActivityList)	GetFirstTableName() error{
	return nil
}
	
	//检查表是否在最前或最后
func (u *ActivityList) CheckTableNameFirstOrLast() (bool,error){
	return true,nil
}
	
	//解析maxid
func (u *ActivityList) ParseMaxID() bool{
	return false
}
	
	//重新设置表
func (u *ActivityList)	ResetTableName() (bool,error){
	return false,nil
}
	
//获取日期表年月
func (u *ActivityList)	GetYearAndMonth() string {
	return ""
}

//设置翻页类型
func (u *ActivityList)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *ActivityList)	SetMaxID(id string) {
	u.MaxID = id
}

//获取翻页类型
func (u *ActivityList)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *ActivityList)	GetMaxID() string {
	return u.MaxID
}

//上页
func (u *ActivityList) PageUp() []orm.Params {
	u.dbConn =  ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT  %s  FROM %s WHERE ID > ? AND IsShow = 0 ORDER BY ID ASC LIMIT %d",tablename_activity_list_column,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID).Values(&table)
	return table
	
}

//下页
func (u *ActivityList) PageDown() []orm.Params {
	u.dbConn =  ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT %s  FROM %s WHERE ID < ?  AND IsShow = 0 ORDER BY ID DESC LIMIT %d",tablename_activity_list_column,u.tableName,XYLibs.TABLE_LIMIT_NUM)

	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID).Values(&table)
	return table
}
//最后页
func (u *ActivityList) PageEnd() []orm.Params {
	u.dbConn =  ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT %s  FROM %s  WHERE IsShow = 0 ORDER BY ID ASC LIMIT %d",tablename_activity_list_column,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//首页
func (u *ActivityList) PageFirst() []orm.Params {
	u.dbConn =  ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT %s  FROM %s  WHERE IsShow = 0   ORDER BY ID Desc LIMIT %d",tablename_activity_list_column,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}


//获取所有活动
func (u *ActivityList) GetAllActivity() []orm.Params {
	u.dbConn =  ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT ID, ActivityID, ActivityName, ActivityDesImg,Url  FROM %s  WHERE IsShow = 0   ORDER BY ID Desc",u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

// 获取活动名
func (u *ActivityList) GetActivityName() []orm.Params {
	u.dbConn =  ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT ID, ActivityID, ActivityName, ActivityDesImg,Url  FROM %s  WHERE IsShow = 0  AND  ActivityID = ? ORDER BY ID Desc LIMIT 1",u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.ActivityID).Values(&table)
	return table
}

