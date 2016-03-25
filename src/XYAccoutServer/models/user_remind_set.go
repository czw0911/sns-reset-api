//@Description 用户提醒设置
//@Contact czw@outlook.com

package models

import (
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"fmt"
)


var (
	userRemindSetTableNameHash = consistent.New()
)

type UserRemindSet struct {

	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	  
	  UID uint32 //用户id
	
	  Comment int8    //评论提醒开关 (0关。1开)
	 
	  Follow int8 //关注提醒开关
	
	  Activity int8 //活动提醒开关）
		
	  Message int8  //留言或聊天提醒开关
}

func (u *UserRemindSet) initDBAndTable(){
		u.dbConn =  ConnMasterDB(u.UID)
		id := fmt.Sprintf("%d",u.UID)
		u.tableName,_ = userRemindSetTableNameHash.Get(id)
}




func (u * UserRemindSet) Set() (bool,error) {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("REPLACE INTO %s(UID,Comment,Follow,Activity,Message) VALUES(?,?,?,?,?)",u.tableName)
		_,err := u.dbConn.Raw(sqlStr,u.UID,u.Comment,u.Follow,u.Activity,u.Message).Exec()
		if err != nil {
			return false,err
		}
		return true,nil
}


func (u * UserRemindSet) Get() []orm.Params {
		u.initDBAndTable()
		sqlStr := fmt.Sprintf("SELECT Comment,Follow,Activity,Message FROM %s WHERE UID = ? ",u.tableName)
		var res []orm.Params
		_,_ = u.dbConn.Raw(sqlStr,u.UID).Values(&res)
		return res
}



func init(){
	for i:=0 ; i < USER_HASH_TABLE_NUM ; i++ {
		tn := fmt.Sprintf("user_remind_set_%d",i)
		userRemindSetTableNameHash.Add(tn)
	}
}