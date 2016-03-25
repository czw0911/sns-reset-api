//@Description 用户关注的组信息
//@Contact czw@outlook.com

package models

import (
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"fmt"
)


var (
	userFollowGroupTableNameHash = consistent.New()
)

type UserFollowGroup struct {

	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	  
	  UID uint32 //用户id
	
	  GroupID uint64    //  关注群组id 
	 
		
}

func (u *UserFollowGroup) initDBAndTable(){
		u.dbConn =  ConnMasterDB(u.UID)
		id := fmt.Sprintf("%d",u.UID)
		u.tableName,_ = userFollowGroupTableNameHash.Get(id)
}



func (u *UserFollowGroup) IsFollow() (bool,error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID FROM %s WHERE UID = ? AND GroupID = ? LIMIT 1",u.tableName)
	var res []orm.Params
	_, err := u.dbConn.Raw(sqlStr,u.UID,u.GroupID).Values(&res)
	if err != nil {
		return true,err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true,nil
}

func (u * UserFollowGroup) FollowGroup() (bool,error) {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("REPLACE INTO %s(UID,GroupID) VALUES(?,?)",u.tableName)
		_,err := u.dbConn.Raw(sqlStr,u.UID,u.GroupID).Exec()
		if err != nil {
			return false,err
		}
		return true,nil
}

func (u * UserFollowGroup) DelFollowGroupByUID() (bool,error) {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("DELETE FROM %s WHERE UID = ?",u.tableName)
		_,err := u.dbConn.Raw(sqlStr,u.UID).Exec()
		if err != nil {
			return false,err
		}
		return true,nil
}

func (u * UserFollowGroup) GetFollowGroupByUID() []orm.Params {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("SELECT GroupID FROM %s WHERE UID = ? ",u.tableName)
		var res []orm.Params
		_,_ = u.dbConn.Raw(sqlStr,u.UID).Values(&res)
		return res
}



func init(){
	for i:=0 ; i < USER_HASH_TABLE_NUM ; i++ {
		tn := fmt.Sprintf("user_follow_group_%d",i)
		userFollowGroupTableNameHash.Add(tn)
	}
}