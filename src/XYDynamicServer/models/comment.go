//评论
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"XYAPIServer/XYLibs"
)


type Comments struct {
	
	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	
	  ID  int64 
	
	  CommentID string //评论id = 动态id _ id
	
	  UID uint32 //评论用户id
	
	  HomeProvinceID int    //  家乡省id 
	
	  DynamicID string //动态id
	
	  Contents string //评论内容
	
	  PostTime int64 //评论时间
	
	  MaxID string //当前列表最大id 格式:YearAndMonth - ID 
	
	  PageType int8 //翻页类型 1,上翻；2，翻页

	  YearAndMonth string //年月
	

}



//评论表
func (u *Comments) initCommentDBAndTable() {
		u.dbConn =  ConnMasterDB(u.HomeProvinceID)
		u.tableName = fmt.Sprintf("comment_%d_%s",u.HomeProvinceID,u.YearAndMonth)	
}


func (u *Comments) Add() (bool,error) {
	u.initCommentDBAndTable()
		sqlStr := `CREATE  TABLE  IF NOT EXISTS ` + u.tableName + `
	 (
	       ID  bigint(20) NOT NULL  AUTO_INCREMENT,
		   DynamicID  varchar(45) DEFAULT NULL COMMENT '动态id',
		   UID  int(11) unsigned DEFAULT NULL COMMENT '提交人uid',
		   Contents  varchar(140) DEFAULT NULL COMMENT '评论内容',
		   PostTime  int(11) unsigned DEFAULT NULL COMMENT '提交时间',
		   IsShow tinyint(1) DEFAULT '0' COMMENT '是否显示',
		  PRIMARY KEY ( ID ),
		  KEY  activity_index  ( DynamicID ,  UID , PostTime, IsShow)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='动态的评论月表';

	`
	_, err := u.dbConn.Raw(sqlStr).Exec()
	if err != nil {
		return false, err
	}
	_,err = u.dbConn.Raw("INSERT INTO "+ u.tableName +"(UID,DynamicID,Contents,PostTime) VALUES(?,?,?,?)",
	u.UID,u.DynamicID,u.Contents,u.PostTime).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//评论是否存在
func (u *Comments) IsComments() bool {
	u.initCommentDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID FROM %s  WHERE ID = ? AND UID = ? AND IsShow = 0 LIMIT 1",u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.ID,u.UID).Values(&table)
	if len(table) > 0 {
		return true
	}
	return false
}

//删除
func (u *Comments) ClickDelete() (bool,error) {
	u.initCommentDBAndTable()	
	sqlStr := fmt.Sprintf("UPDATE %s SET IsShow = 1  WHERE ID = ? AND UID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.ID,u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

	//最后的表
func (u *Comments) GetLastTableName() error{
	return nil
}
	
	//最前的表
func (u *Comments)	GetFirstTableName() error{
	return nil
}
	
	//检查表是否在最前或最后
func (u *Comments) CheckTableNameFirstOrLast() (bool,error){
	return true,nil
}
	
	//解析maxid
func (u *Comments) ParseMaxID() bool{
	return false
}
	
	//重新设置表
func (u *Comments)	ResetTableName() (bool,error){
	return false,nil
}

//获取日期表年月
func (u *Comments)	GetYearAndMonth() string {
	return ""
}

//设置翻页类型
func (u *Comments)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *Comments)	SetMaxID(id string) {
	u.MaxID = id
}

//获取翻页类型
func (u *Comments)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *Comments)	GetMaxID() string {
	return u.MaxID
}

//上页
func (u *Comments) PageUp() []orm.Params {
	u.initCommentDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID,  UID,DynamicID,Contents,PostTime  FROM %s WHERE ID > ? AND DynamicID = ? AND IsShow = 0 ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID,u.DynamicID).Values(&table)
	return table
	
}

//下页
func (u *Comments) PageDown() []orm.Params {
	u.initCommentDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID, UID,DynamicID,Contents,PostTime  FROM %s WHERE ID < ?  AND DynamicID = ?  AND IsShow = 0 ORDER BY ID DESC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)

	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.MaxID,u.DynamicID).Values(&table)
	return table
}

//最后页
func (u *Comments) PageEnd() []orm.Params {
	u.initCommentDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID, UID,DynamicID,Contents,PostTime  FROM %s  WHERE DynamicID = ? AND IsShow = 0  ORDER BY ID ASC LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DynamicID).Values(&table)
	return table
}

//首页
func (u *Comments) PageFirst() []orm.Params {
	u.initCommentDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID, UID,DynamicID,Contents,PostTime FROM %s  WHERE DynamicID = ?  AND IsShow = 0  ORDER BY ID Desc LIMIT %d",u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DynamicID).Values(&table)
	return table
}

