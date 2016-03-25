//@Description 用户的好友
//@Contact czw@outlook.com

package models

import (
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"fmt"
)


var (
	userFriendsTableNameHash = consistent.New()
)

type UserFriends struct {

	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	  
	  UID uint32 //用户id
	
	  FUID uint32    //  好友uid
	 

		
}

func (u *UserFriends) initDBAndTable(){
		u.dbConn =  ConnMasterDB(u.UID)
		id := fmt.Sprintf("%d",u.UID)
		u.tableName,_ = userFriendsTableNameHash.Get(id)
}



func (u *UserFriends) IsFriends() (bool,error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID FROM %s WHERE UID = ? AND FUID = ? LIMIT 1",u.tableName)
	var res []orm.Params
	_, err := u.dbConn.Raw(sqlStr,u.UID,u.FUID).Values(&res)
	if err != nil {
		return true,err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true,nil
}

func (u * UserFriends) AddFriends() (bool,error) {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("INSERT INTO %s(UID,FUID) VALUES(?,?)",u.tableName)
		_,err := u.dbConn.Raw(sqlStr,u.UID,u.FUID).Exec()
		if err != nil {
			return false,err
		}
		return true,nil
}

func (u * UserFriends) GetFriendsByUID() []orm.Params {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("SELECT FUID FROM %s WHERE UID = ? ",u.tableName)
		var res []orm.Params
		_,_ = u.dbConn.Raw(sqlStr,u.UID).Values(&res)
		return res
}



func init(){
	for i:=0 ; i < USER_HASH_TABLE_NUM ; i++ {
		tn := fmt.Sprintf("user_friends_%d",i)
		userFriendsTableNameHash.Add(tn)
	}
}