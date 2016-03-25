//谈论活动
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"
	"XYAPIServer/XYLibs"
	"XYAPIServer/XYActivityServer/libs"
)

const  ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME = " ID,UID,TalkID,TalkContent,BadNUM,GoodNUM,CommentNUM,Images,Voices,VoiceLen,PostTime,LastComment,CommentUser"

type TalkList struct {

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
	
	  LastComment string //最后一条评论内容
	
	  CommentUser string //最后一条评论人
	  
	  YearAndMonth string //年月
	
	  MaxID string //当前列表最大id 格式:YearAndMonth - ID 
	
	  PageType int8 //翻页类型 1,上翻；2，翻页
	
	  DataIndex int64 //列表数据当前索引
	
	  prevTableName string //上次的表面
	
	 
	
	
}



//活动讨论月表
func (u *TalkList) initActivityTalkDBAndTable() {
		u.dbConn =  ConnMasterDB(u.ActivityID)
		if u.YearAndMonth == "" {
			u.YearAndMonth = time.Now().Format("200601")
		}
		u.tableName = fmt.Sprintf("talk_%d_%s",u.ActivityID,u.YearAndMonth)	
}

//生成讨论id  格式：时间戳 - 活动编号 - uid 
func (u *TalkList) GenerateTalkID()  {
	id := XYLibs.GenerateNineteenPrefixID()
	u.TalkID =  fmt.Sprintf("%s%d%d",id,u.ActivityID,u.UID)
}

func (u *TalkList) ParseTalkID() bool {
	
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

//谈论活动
func (u *TalkList) TalkActivity() (bool,error) {
	u.initActivityTalkDBAndTable()
	u.GenerateTalkID()
	
	sqlStr := `CREATE  TABLE  IF NOT EXISTS ` + u.tableName + `
	 (
	       ID  bigint(20) NOT NULL  AUTO_INCREMENT,
		   TalkID  varchar(45) DEFAULT NULL COMMENT '谈论id（时间戳_活动id_uid)',
		   UID  int(11) unsigned DEFAULT NULL COMMENT '提交人uid',
		   TalkContent  varchar(140) DEFAULT NULL COMMENT '谈论内容',
		   LastComment   varchar(140) DEFAULT NULL COMMENT '最后一条评论',
		   CommentUser   varchar(40) DEFAULT NULL COMMENT '最后一条评论人',
		   PostTime  int(11) unsigned DEFAULT NULL COMMENT '提交时间',
		   GoodNUM  int(11) DEFAULT 0 COMMENT '点赞数',
		   BadNUM  int(11) DEFAULT 0 COMMENT '吐槽数',
		   CommentNUM  int(11) DEFAULT 0 COMMENT '评论数',
		   Images  varchar(1000) DEFAULT NULL COMMENT '提交的图片，多图用逗号分隔',
		   Voices  varchar(1000) DEFAULT NULL COMMENT '提交的声音',
		   VoiceLen  int(11) DEFAULT NULL COMMENT '提交的声音长度',
		   IsShow tinyint(1) DEFAULT '0' COMMENT '是否显示',
		  PRIMARY KEY ( ID ),
		  UNIQUE KEY ActivityID_UNIQUE (TalkID),
		  KEY  activity_index  ( TalkID ,  UID , PostTime,IsShow )
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动谈论月表';

	`
	_, err := u.dbConn.Raw(sqlStr).Exec()
	if err != nil {
		return false, err
	}
	_,err = u.dbConn.Raw("INSERT INTO "+ u.tableName +"(UID,TalkID,TalkContent,Images,Voices,VoiceLen,PostTime) VALUES(?,?,?,?,?,?,?)",
	u.UID,u.TalkID,u.TalkContent,u.Images,u.Voices,u.VoiceLen,u.PostTime).Exec()
	if err != nil {
		return false, err
	}
	
	_,err = u.SetTableNameIndex()
	if err != nil {
		return false, err
	}
	return true, nil
}

//活动表名称索引
func (u *TalkList) SetTableNameIndex() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_TABLE_INDEX,u.ActivityID)
	redisDB := libs.RedisDBActivity
	_,err := redisDB.ZADD(noSQLKey,u.YearAndMonth,u.tableName)
	if err != nil {
			return false,err
	}
	
	noSQLKey = fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX,u.UID)
	_,err = redisDB.ZADD(noSQLKey,u.YearAndMonth,u.tableName)
	if err != nil {
			return false,err
	}
	return true , nil
}

func (u *TalkList) ParseMaxID() bool{
	
	u.DataIndex = 0
	if len(u.MaxID) < 7 {
	
		return false
	}
	
	tableIndex,err := strconv.Atoi(u.MaxID[0:6])
	if err != nil {
		return false
	}
	
	u.YearAndMonth = fmt.Sprintf("%d",tableIndex)
	
	id ,err  := strconv.ParseInt(u.MaxID[6:],10,64)
	if err != nil {
		return false
	}
	u.DataIndex = id
	u.initActivityTalkDBAndTable()
	return true
}

//活动表是否为空
func (u *TalkList) IsEmptyTableName()(bool,error){
	
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_TABLE_INDEX,u.ActivityID)
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
func (u *TalkList)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *TalkList)	SetMaxID(id string) {
	u.MaxID = id
}

//获取日期表年月
func (u *TalkList)	GetYearAndMonth() string {
	return u.YearAndMonth
}

//获取翻页类型
func (u *TalkList)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *TalkList)	GetMaxID() string {
	return u.MaxID
}

//最前的表
func (u *TalkList) GetFirstTableName() error {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_TABLE_INDEX,u.ActivityID)	
	tname,err := libs.RedisDBActivity.ZREVRANGE(noSQLKey,"0","0")
	if err != nil {
			return err
	}
	t := tname.([]interface{})
	u.tableName = string(t[0].([]uint8))
	a := strings.Split(u.tableName,"_")
	u.YearAndMonth = a[2]
	return nil
}

//最后的表
func (u *TalkList) GetLastTableName() error{
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_TABLE_INDEX,u.ActivityID)
	
	tname,err := libs.RedisDBActivity.ZREVRANGE(noSQLKey,"-1","-1")
	if err != nil {
			return err
	}
	t := tname.([]interface{})
	u.tableName = string(t[0].([]uint8))
	a := strings.Split(u.tableName,"_")
	u.YearAndMonth = a[2]
	return nil
}

//检查活动表是否在最前或最后
func (u *TalkList) CheckTableNameFirstOrLast() (bool,error){
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_TABLE_INDEX,u.ActivityID)
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
func (u *TalkList) ResetTableName() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_TALK_TABLE_INDEX,u.ActivityID)
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
	u.YearAndMonth = u.tableName[14:]
	if u.prevTableName != u.tableName {
		u.DataIndex = 0
	}
	
	return true , nil
}


//是否点赞或吐槽过,返回 0:没有点击过;1点击过点赞；2点击过吐槽
func (u *TalkList) IsClickGoodOrBad() (string,error) {
		noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_CLICK_GOOD_OR_BAD_ID,u.UID)
		val := fmt.Sprintf("%s",u.TalkID)
		res,err := libs.RedisDBActivity.ZSCORE(noSQLKey,val)
		if err != nil || res == nil {
				return "0",err
		}
		return string(res.([]uint8)),nil
}

//点赞
func (u *TalkList) ClickGood() (bool,error) {
	p := u.ParseTalkID()
	if !p {
		return false,errors.New("谈论id格式错误")
	}
	u.initActivityTalkDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET GoodNUM = GoodNUM + 1 WHERE TalkID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.TalkID).Exec()
	if err != nil {
		return false, err
	}
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_CLICK_GOOD_OR_BAD_ID,u.UID)	
	_,err = libs.RedisDBActivity.ZADD(noSQLKey,"1",u.TalkID)
	if err != nil {
			return false,err
	}
	return true, nil
}

//吐槽
func (u *TalkList) ClickBad() (bool,error) {
	p := u.ParseTalkID()
	if !p {
		return false,errors.New("谈论id格式错误")
	}
	u.initActivityTalkDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET BadNUM = BadNUM + 1 WHERE TalkID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.TalkID).Exec()
	if err != nil {
		return false, err
	}
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_ACTIVITY_CLICK_GOOD_OR_BAD_ID,u.UID)	
	_,err = libs.RedisDBActivity.ZADD(noSQLKey,"2",u.TalkID)
	if err != nil {
			return false,err
	}
	return true, nil
}

//删除
func (u *TalkList) ClickDelete () (bool,error) {
	p := u.ParseTalkID()
	if !p {
		return false,errors.New("谈论id格式错误")
	}
	u.initActivityTalkDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE  %s  SET  IsShow = 1   WHERE TalkID = ?  AND  UID = ? ",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.TalkID,u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//评论
func (u *TalkList) ClickComment() (bool,error) {
	p := u.ParseTalkID()
	if !p {
		return false,errors.New("谈论id格式错误")
	}
	u.initActivityTalkDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET CommentNUM = CommentNUM + 1 , LastComment = ? ,CommentUser = ?  WHERE TalkID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.LastComment,u.CommentUser,u.TalkID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}


//评论删除
func (u *TalkList) ClickDelComment() (bool,error) {
	p := u.ParseTalkID()
	if !p {
		return false,errors.New("谈论id格式错误")
	}
	u.initActivityTalkDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET CommentNUM = CommentNUM - 1   WHERE TalkID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.TalkID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}



//上页
func (u *TalkList) PageUp() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT  %s FROM %s WHERE ID > ? AND IsShow = 0  ORDER BY ID ASC LIMIT %d",ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex).Values(&table)
	return table
	
}

//下页
func (u *TalkList) PageDown() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := ""
	if u.DataIndex > 0 {
		sqlStr = fmt.Sprintf("SELECT %s  FROM %s WHERE ID < ? AND IsShow = 0  ORDER BY ID DESC LIMIT %d",ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	}else{
		sqlStr = fmt.Sprintf("SELECT %s  FROM %s WHERE ID > ? AND IsShow = 0  ORDER BY ID DESC LIMIT %d",ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	}
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex).Values(&table)
	return table
}
//最后页
func (u *TalkList) PageEnd() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT %s  FROM %s WHERE IsShow = 0  ORDER BY ID ASC LIMIT %d",ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//第一页
func (u *TalkList) PageFirst() []orm.Params {
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  WHERE IsShow = 0 ORDER BY ID DESC LIMIT %d",ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//获取指定内容
func (u *TalkList) GetDataByTalkID() []orm.Params {
	u.initActivityTalkDBAndTable()
	u.dbConn =  ConnMasterDB(u.ActivityID)
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  WHERE  TalkID = ?  ORDER BY ID DESC LIMIT 1",ACTIVITY_TALK_TABLE_SELECT_COLUMN_NAME,u.tableName)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.TalkID).Values(&table)
	return table
}
