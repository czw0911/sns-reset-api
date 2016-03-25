//群组
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	//"time"
	//"errors"
)


//群组
type Groups struct {
	
	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	
	  UID uint32 //用户id

      GroupID int64 //群组id
	
	  ActivityID string //活动id
}


//群组用户表
func (u *Groups) initGroupUidDBAndTable() {
		u.dbConn =  ConnMasterDB(u.GroupID)
		u.tableName = fmt.Sprintf("group_user_%d",u.GroupID)	
}

//加入群组
func (u *Groups) JoinGroup()(bool,error) {
	u.initGroupUidDBAndTable()
	sqlStr := `CREATE  TABLE  IF NOT EXISTS ` + u.tableName + `
	 (
	  ID int(11) NOT NULL AUTO_INCREMENT,
	  UID int(11) unsigned DEFAULT NULL COMMENT '用户uid',
	  PRIMARY KEY (ID),
	  UNIQUE KEY UID_UNIQUE (UID)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='群组用户表'
	`
	_, err := u.dbConn.Raw(sqlStr).Exec()
	if err != nil {
		return false, err
	}
	_,err = u.dbConn.Raw("REPLACE INTO "+ u.tableName +"(UID) VALUES(?)",u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//退出群组
func (u *Groups) ExitGroup()(bool,error){
	u.initGroupUidDBAndTable()
	sqlStr := fmt.Sprintf("DELETE FROM %s WHERE UID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//群组是否存在
func (u *Groups) IsExistGroup()(bool,error){
	db := ConnCommonDB()
	sqlStr := fmt.Sprintf("SELECT ID  FROM %s WHERE GroupID = ?","group_list")
	var res []orm.Params
	_, err := db.Raw(sqlStr,u.GroupID).Values(&res)
	if err != nil {
		return false, err
	}
	
	if len(res) > 0 {
		return true,nil
	}
	return false, nil
}


//群组的用户
func (u *Groups) GetUsers() []orm.Params {
	u.initGroupUidDBAndTable()
	sqlStr := fmt.Sprintf("SELECT UID FROM %s",u.tableName)
	var res []orm.Params
	_,_ = u.dbConn.Raw(sqlStr).Values(&res)
	return res
}

