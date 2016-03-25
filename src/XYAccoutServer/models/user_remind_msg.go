//@Description 用户提醒消息
//@Contact czw@outlook.com

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"
	"strconv"
	"XYAPIServer/XYLibs"
	"time"
)

var (
	userRemindMsgTableNameHash = consistent.New()
)

type UserRemindMsg struct {
	
	tableName string //hash到的表名

	dbConn orm.Ormer //hash到的数据库

	UID uint32 //用户id
	
	MsgTypeID string // 消息类型
	
	UnreadNum int //未读消息数;乡音团队（系统）消息用时间戳判断
	
	ReadNum  int //已读数
	
	LastMsg string //最后一条消息
	
	LastTime int64 //最后一后消息时间
	
}

func (u *UserRemindMsg) initDBAndTable() {
	u.dbConn = ConnMasterDB(u.UID)
	id := fmt.Sprintf("%d", u.UID)
	u.tableName, _ = userRemindMsgTableNameHash.Get(id)
}


func (u *UserRemindMsg) Add() (bool, error){ 
	u.initDBAndTable()
	UnreadNum := 1
	if u.MsgTypeID != XYLibs.REMIND_MESSAGE_TYPE_C {
		sqlStr := fmt.Sprintf("SELECT UnreadNum  FROM %s  WHERE UID = ? AND MsgTypeID = ?  LIMIT 1", u.tableName)
		var data []orm.Params
		u.dbConn.Raw(sqlStr,u.UID, u.MsgTypeID).Values(&data)
		
		if len(data) > 0 {
			t ,_ := strconv.Atoi(data[0]["UnreadNum"].(string))
			UnreadNum = t + 1
		}
	}else{
		UnreadNum = u.UnreadNum
	}
	
	sqlStr := fmt.Sprintf("REPLACE INTO %s(UID,MsgTypeID,LastMsg,LastTime,UnreadNum) VALUES(?,?,?,?,?)", u.tableName)
	_, err := u.dbConn.Raw(sqlStr,  u.UID, u.MsgTypeID,u.LastMsg,u.LastTime,UnreadNum).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}




//设置已读
func (u *UserRemindMsg) SetReadNum() (bool, error){ 
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE %s SET UnreadNum = UnreadNum - ?  WHERE UID = ? AND MsgTypeID = ? ", u.tableName)
	if u.MsgTypeID == XYLibs.REMIND_MESSAGE_TYPE_C {
		//乡音团队消息用时间判断
		sqlStr = fmt.Sprintf("UPDATE %s SET UnreadNum = ?  WHERE UID = ? AND MsgTypeID = ? ", u.tableName)
		u.ReadNum = int(time.Now().Unix())
	}
	_, err := u.dbConn.Raw(sqlStr, u.ReadNum, u.UID, u.MsgTypeID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserRemindMsg) Get() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT MsgTypeID,UnreadNum,LastMsg,LastTime FROM %s  WHERE UID = ? ", u.tableName)
	var data []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&data)
	return data
}

//获取最后已读的乡音团队消息
func (u *UserRemindMsg) GetLastReadSysMsg() string {	
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT UnreadNum  FROM %s  WHERE UID = ? AND MsgTypeID = ?  LIMIT 1", u.tableName)
	var data []orm.Params
	u.dbConn.Raw(sqlStr,u.UID, XYLibs.REMIND_MESSAGE_TYPE_C).Values(&data)
	UnreadNum := "1433520000"
	if len(data) > 0 {
		UnreadNum = data[0]["UnreadNum"].(string)
	}
	return UnreadNum
}

//获取最新的乡音团队消息
func (u *UserRemindMsg) GetNewsSysMsg() []orm.Params {	
	sys := new(SysMsgLog)
	return sys.NewsOne()
}

//获取最新的乡音团队未读消息数
func (u *UserRemindMsg) GetSysMsgUnreadNum() int {	
	lastTime := u.GetLastReadSysMsg()
	sys := new(SysMsgLog)
	data := sys.GetUnReadNUM(lastTime)
	unreadNum := 0
	if len(data) > 0 {
		unreadNum , _ = strconv.Atoi(data[0]["UnreadNum"].(string))
	}
	return unreadNum
}


func init() {
	for i := 0; i < USER_HASH_TABLE_NUM; i++ {
		tn := fmt.Sprintf("user_remind_msg_%d", i)
		userRemindMsgTableNameHash.Add(tn)
	}
}
