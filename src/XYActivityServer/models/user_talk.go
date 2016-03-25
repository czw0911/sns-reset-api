//我的谈论活动
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"strings"
	"strconv"
	//"errors"
	"XYAPIServer/XYLibs"
	"XYAPIServer/XYActivityServer/libs"
)
const USER_TALK_TABLE_SELECT_COLUMN_NAME = " ID,UID,TalkID,TalkContent,BadNUM,GoodNUM,CommentNUM,Images,Voices,VoiceLen,PostTime " 


type UserTalkList struct {

	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	
	  ActivityID  int64 //活动编号
	  
	  UID uint32 //用户id
	
	  TalkID string //谈论id
	
	  TalkContent string //谈论内容 

	  PostTime int64 //提交时间
	
	  GoodNUM int   //点赞数
	
	  BadNUM int //吐槽数
	
	  CommentNum int //回复数
	
	  Images string //图片地址
	
	  Voices  string //音视频地址
	  
      VoiceLen int //音频长度
	  
	  YearAndMonth string //年月
	
	  MaxID string //当前列表最大id 格式:YearAndMonth - ID 
	
	  PageType int8 //翻页类型 1,上翻；2，翻页
	
	  DataIndex int64 //列表数据当前索引
	
	  prevTableName string //上次的表面
	
	
}



//活动讨论月表
func (u *UserTalkList) initActivityTalkDBAndTable() {
		u.dbConn =  ConnMasterDB(u.ActivityID)
		if u.YearAndMonth == "" {
			u.YearAndMonth = time.Now().Format("200601")
		}
		u.tableName = fmt.Sprintf("talk_%d_%s",u.ActivityID,u.YearAndMonth)	
}

//生成讨论id  格式：时间戳 - 活动编号 - uid 
func (u *UserTalkList) GenerateTalkID()  {
	id := XYLibs.GenerateNineteenPrefixID()
	u.TalkID =  fmt.Sprintf("%s%d%d",id,u.ActivityID,u.UID)
}

func (u *UserTalkList) ParseTalkID() bool {
	
	//fmt.Printf("%#v\n",arr)
	if len(u.TalkID) < 27 {
		return false
	}
	u.ActivityID,_ = strconv.ParseInt(u.TalkID[19:27],10,64)
	t,err := strconv.ParseInt(u.TalkID[:19],10,64)
	if err != nil {
		return false
	}
	u.YearAndMonth = time.Unix(t / 1000000000,0).Format("200601")
	return true
}





func (u *UserTalkList) ParseMaxID() bool{
	
	u.DataIndex = 0
	if len(u.MaxID) < 7 {
	
		return false
	}
	
	tableIndex,err := strconv.Atoi(u.MaxID[0:6])
	if err != nil {
		return false
	}
	
	u.YearAndMonth = fmt.Sprintf("%d",tableIndex)
	u.ActivityID,err =  strconv.ParseInt(u.MaxID[6:14],10,64)
	if err != nil {
		return false
	}
	id ,err  := strconv.ParseInt(u.MaxID[14:],10,64)
	if err != nil {
		return false
	}
	u.DataIndex = id
	u.initActivityTalkDBAndTable()
	return true
}


//是否点赞或吐槽过,返回 0:没有点击过;1点击过点赞；2点击过吐槽
func (u *UserTalkList) IsClickGoodOrBad() (string,error) {
		noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_CLICK_GOOD_OR_BAD_ID,u.UID)
		val := fmt.Sprintf("%s",u.TalkID)
		res,err := libs.RedisDBActivity.ZSCORE(noSQLKey,val)
		if err != nil || res == nil {
				return "0",err
		}
		return string(res.([]uint8)),nil
}

//活动表是否为空
func (u *UserTalkList) IsEmptyTableName()(bool,error){
	
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX,u.UID)
	count,err := libs.RedisDBActivity.ZCOUNT(noSQLKey,XYLibs.TABLE_NAME_START_DATE,time.Now().Format("200601"))
	if err != nil {
			return true,err
	}
	if count.(int64) > 0 {
		return false,nil
	}
	return true,nil
}

//设置翻页类型
func (u *UserTalkList)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *UserTalkList)	SetMaxID(id string) {
	u.MaxID = id
}

//获取日期表年月
func (u *UserTalkList)	GetYearAndMonth() string {
	return u.YearAndMonth
}

//获取翻页类型
func (u *UserTalkList)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *UserTalkList)	GetMaxID() string {
	return u.MaxID
}

//最前的表
func (u *UserTalkList) GetFirstTableName() error {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX,u.UID)	
	tname,err := libs.RedisDBActivity.ZREVRANGE(noSQLKey,"0","0")
	if err != nil {
			return err
	}
	t := tname.([]interface{})
	u.tableName = string(t[0].([]uint8))
	a := strings.Split(u.tableName,"_")
	u.ActivityID,_ = strconv.ParseInt(a[1],10,64)
	u.YearAndMonth = a[2]
	return nil
}

//最后的表
func (u *UserTalkList) GetLastTableName() error{
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX,u.UID)
	
	tname,err := libs.RedisDBActivity.ZREVRANGE(noSQLKey,"-1","-1")
	if err != nil {
			return err
	}
	t := tname.([]interface{})
	u.tableName = string(t[0].([]uint8))
	a := strings.Split(u.tableName,"_")
	u.ActivityID,_ = strconv.ParseInt(a[1],10,64)
	u.YearAndMonth = a[2]
	return nil
}

//检查活动表是否在最前或最后
func (u *UserTalkList) CheckTableNameFirstOrLast() (bool,error){
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX,u.UID)
	redisDB := libs.RedisDBActivity
	count,err := redisDB.ZCOUNT(noSQLKey,XYLibs.TABLE_NAME_START_DATE,time.Now().Format("200601"))
	if err != nil || count == nil {
			return true,err
	}
	if count.(int64) == 1 {
			return true,nil
	}
	rank,err := redisDB.ZREVRANK(noSQLKey,u.tableName)
	if err != nil || rank == nil {
		return false,err
	}
	name := ""
	if u.PageType == 1 {
		tname,err := redisDB.ZREVRANGE(noSQLKey,"0","0")
		if err != nil  || tname == nil {
				return true,err
		}
		t := tname.([]interface{})
		name = string(t[0].([]uint8))
	}else{
		tname,err := redisDB.ZREVRANGE(noSQLKey,"-1","-1")
		if err != nil || tname == nil {
				return true,err
		}
		t := tname.([]interface{})
		name = string(t[0].([]uint8))
	}
	if u.tableName == name {
		return true,nil
	}	
	return false,nil
}

//重新设置活动表名称
func (u *UserTalkList) ResetTableName() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX,u.UID)
	redisDB := libs.RedisDBActivity
	rank,err := redisDB.ZREVRANK(noSQLKey,u.tableName)
	if err != nil || rank == nil {
		return false,err
	}
	
	start := "0"
	stop := "0"
	index := int64(0)
	
	switch u.PageType {
		case 1:
		//上一页				
				if rank.(int64) > 0 {
					index = rank.(int64) - 1
				}
				
		default:
		//下一页
				count,err := redisDB.ZCOUNT(noSQLKey,XYLibs.TABLE_NAME_START_DATE,time.Now().Format("200601"))
				if err != nil  {
					return false,err
				}
				
		  		index = rank.(int64) + 1
				
				if index >=  count.(int64) {
					index = rank.(int64)
				}	
		
	}
	
	start = fmt.Sprintf("%d",index)
	stop = start
	
	tname,err := redisDB.ZREVRANGE(noSQLKey,start,stop)
	if err != nil {
			return false,err
	}
	t := tname.([]interface{})
	u.prevTableName = u.tableName
	u.tableName = string(t[0].([]uint8))
	u.ActivityID,_ = strconv.ParseInt(u.tableName[5:13],10,64)
	u.YearAndMonth = u.tableName[14:]
	if u.prevTableName != u.tableName {
		u.DataIndex = 0
	}
	
	return true , nil
}








//上页
func (u *UserTalkList) PageUp() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT %s FROM %s WHERE ID > ?  AND UID = ?  ORDER BY ID ASC LIMIT %d",USER_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex,u.UID).Values(&table)
	return table
	
}

//下页
func (u *UserTalkList) PageDown() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := ""
	if u.DataIndex > 0 {
		sqlStr = fmt.Sprintf("SELECT %s FROM %s WHERE ID < ?  AND UID = ? ORDER BY ID DESC LIMIT %d",USER_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	}else{
		sqlStr = fmt.Sprintf("SELECT %s FROM %s WHERE ID > ?  AND UID = ?  ORDER BY ID DESC LIMIT %d",USER_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	}
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex,u.UID).Values(&table)
	return table
}
//最后页
func (u *UserTalkList) PageEnd() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  WHERE UID = ?  ORDER BY ID ASC LIMIT %d",USER_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}

//第一页
func (u *UserTalkList) PageFirst() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT %s FROM %s   WHERE UID = ?  ORDER BY ID DESC LIMIT %d",USER_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.UID).Values(&table)
	return table
}


