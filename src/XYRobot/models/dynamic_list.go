//动态列表
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"
	"XYAPIServer/XYLibs"
	"XYAPIServer/XYRobot/libs"
)

const DYNAMIC_TABLE_SELECT_COLUMN_NAME = " ID,UID,DynamicID,DynamicContent,ViewNUM,CommentNum,ForwardNum,GoodNUM,Images,Voices,VoiceLen,PostTime " 

type DynamicList struct {

	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	
	  UID uint32 //用户id
	
	  HomeProvinceID int    //  家乡省id  
	
	  HomeCityID  int   //  家乡城市id   
	
	  LivingProvinceID  int  //  居住地省id  
	
	  LivingCityID   int //  居住地城市id  
	
	  DynamicID string //动态id
	
	  DynamicContent string //动态内容 

	  PostTime int64 //提交时间
	
	  GoodNUM int   //点赞数
	
	  //IsClickGood int8 //是否已点赞 0，没有 1，有
	
	  ViewNUM int //浏览数
	
	  CommentNum int //回复数
	
	  ForwardNum int //转发数
	
	  Images string //图片地址
	
	  Voices  string //音视频地址
	  
      VoiceLen int //音频长度
	  
	  YearAndMonth string //年月
	
	  MaxID string //当前列表最大id 格式:YearAndMonth - ID 
	
	  PageType int8 //翻页类型 1,上翻；2，翻页
	
	  DataIndex int64 //列表数据当前索引
	
	  prevTableName string //上次的表面
	
	
}



//月表按家乡省份hash
func (u *DynamicList) initDynamicDBAndTable() {
		u.dbConn =  ConnMasterDBDynamic(u.HomeProvinceID)
		if u.YearAndMonth == "" {
			u.YearAndMonth = time.Now().Format("200601")
		}
		u.tableName = fmt.Sprintf("dynamic_%d_%s",u.HomeProvinceID,u.YearAndMonth)	
}

//生成动态id  格式：时间戳 - 家乡省份编号 - uid 
func (u *DynamicList) GenerateDynamicID()  {
	id := XYLibs.GenerateNineteenPrefixID()
	u.DynamicID =  fmt.Sprintf("%s%d%d",id,u.HomeProvinceID,u.UID)
}

func (u *DynamicList) ParseDynamicID() bool {
	
	//fmt.Printf("%#v\n",arr)
	if len(u.DynamicID) < 25 {
		return false
	}
	
	u.HomeProvinceID,_ = strconv.Atoi(u.DynamicID[19:25])
	t,err := strconv.ParseInt(u.DynamicID[:19],10,64)
	if err != nil {
		return false
	}
	u.YearAndMonth = time.Unix(t / 1000000000,0).Format("200601")
	return true
}

//添加动态
func (u *DynamicList) Add() (bool,error) {
	u.initDynamicDBAndTable()
	u.GenerateDynamicID()
	
	sqlStr := `CREATE  TABLE  IF NOT EXISTS ` + u.tableName + `
	 (
	       ID  bigint(20) NOT NULL  AUTO_INCREMENT,
		   DynamicID  varchar(45) DEFAULT NULL COMMENT '动态id（时间戳_家乡省id_uid)',
		   UID  int(11) unsigned DEFAULT NULL COMMENT '提交人uid',
		   DynamicContent  varchar(140) DEFAULT NULL COMMENT '动态内容',
		   LastComment   varchar(140) DEFAULT NULL COMMENT '最后一条评论',
		   PostTime  int(11) unsigned DEFAULT NULL COMMENT '提交时间',
		   GoodNUM  int(11) DEFAULT 0 COMMENT '点赞数',
		   ForwardNUM  int(11) DEFAULT 0 COMMENT '转发数',
		   ViewNUM  int(11) DEFAULT 0 COMMENT '浏览数',
		   CommentNUM  int(11) DEFAULT 0 COMMENT '评论数',
           HomeCityID  int(11) DEFAULT NULL COMMENT '家乡城市id',
           LivingProvinceID  int(11) DEFAULT NULL COMMENT '居住地省id',
           LivingCityID  int(11) DEFAULT NULL COMMENT '居住地城市id',
		   Images  varchar(1000) DEFAULT NULL COMMENT '提交的图片，多图用逗号分隔',
		   Voices  varchar(1000) DEFAULT NULL COMMENT '提交的声音',
		   VoiceLen  int(11) DEFAULT NULL COMMENT '提交的声音长度',
		   IsShow tinyint(1) DEFAULT '0' COMMENT '是否显示', 
		  PRIMARY KEY ( ID ),
		  UNIQUE KEY DynamicID_UNIQUE (DynamicID),
		  KEY  dynamic_index  ( DynamicID ,  UID , PostTime,HomeCityID,LivingProvinceID, LivingCityID,IsShow)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家乡省hash动态月表';

	`
	_, err := u.dbConn.Raw(sqlStr).Exec()
	if err != nil {
		return false, err
	}
	_,err = u.dbConn.Raw("INSERT INTO "+ u.tableName +"(UID,DynamicID,DynamicContent,Images,Voices,VoiceLen,PostTime,HomeCityID,LivingProvinceID, LivingCityID) VALUES(?,?,?,?,?,?,?,?,?,?)",
	u.UID,u.DynamicID,u.DynamicContent,u.Images,u.Voices,u.VoiceLen,u.PostTime,u.HomeCityID,u.LivingProvinceID, u.LivingCityID).Exec()
	if err != nil {
		return false, err
	}
	
	_,err = u.SetTableNameIndex()
	if err != nil {
		return false, err
	}
	return true, nil
}

//表名称索引
func (u *DynamicList) SetTableNameIndex() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_TABLE_INDEX,u.HomeProvinceID)
	redisDB := libs.RedisDBDynamic
	_,err := redisDB.ZADD(noSQLKey,u.YearAndMonth,u.tableName)
	if err != nil {
			return false,err
	}
	
	noSQLKey = fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX,u.UID)
	_,err = redisDB.ZADD(noSQLKey,u.YearAndMonth,u.tableName)
	if err != nil {
			return false,err
	}
	return true , nil
}

func (u *DynamicList) ParseMaxID() bool{
	
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
	u.initDynamicDBAndTable()
	return true
}

//表是否为空
func (u *DynamicList) IsEmptyTableName()(bool,error){
	
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_TABLE_INDEX,u.HomeProvinceID)
	count,err := libs.RedisDBDynamic.ZCOUNT(noSQLKey,XYLibs.TABLE_NAME_START_DATE,time.Now().Format("200601"))
	if err != nil {
			return true,err
	}
	if count.(int64) > 0 {
		return false,nil
	}
	return true,nil
}

//设置翻页类型
func (u *DynamicList)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *DynamicList)	SetMaxID(id string) {
	u.MaxID = id
}

//获取日期表年月
func (u *DynamicList)	GetYearAndMonth() string {
	return u.YearAndMonth
}

//获取翻页类型
func (u *DynamicList)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *DynamicList)	GetMaxID() string {
	return u.MaxID
}

//最前的表
func (u *DynamicList) GetFirstTableName() error {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_TABLE_INDEX,u.HomeProvinceID)	
	tname,err := libs.RedisDBDynamic.ZREVRANGE(noSQLKey,"0","0")
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
func (u *DynamicList) GetLastTableName() error{
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_TABLE_INDEX,u.HomeProvinceID)
	
	tname,err := libs.RedisDBDynamic.ZREVRANGE(noSQLKey,"-1","-1")
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
func (u *DynamicList) CheckTableNameFirstOrLast() (bool,error){
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_TABLE_INDEX,u.HomeProvinceID)
	redisDB := libs.RedisDBDynamic
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
func (u *DynamicList) ResetTableName() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_TABLE_INDEX,u.HomeProvinceID)
	redisDB := libs.RedisDBDynamic
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
	u.YearAndMonth = u.tableName[15:]
	if u.prevTableName != u.tableName {
		u.DataIndex = 0
	}
	
	return true , nil
}



//点赞
func (u *DynamicList) ClickGood() (bool,error) {
	p := u.ParseDynamicID()
	if !p {
		return false,errors.New("动态id格式错误")
	}
	u.initDynamicDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET GoodNUM = GoodNUM + 1 WHERE DynamicID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.DynamicID).Exec()
	if err != nil {
		return false, err
	}
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_CLICK_GOOD_DYNAMIC_ID,u.UID)	
	_,err = libs.RedisDBDynamic.SADD(noSQLKey,u.DynamicID)
	if err != nil {
			return false,err
	}
	return true, nil
}

//是否点赞
func (u *DynamicList) IsClickGood() (int8,error) {
		noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_CLICK_GOOD_DYNAMIC_ID,u.UID)
		val := fmt.Sprintf("%s",u.DynamicID)
		res,err := libs.RedisDBDynamic.SISMEMBER(noSQLKey,val)
		if err != nil || res == nil {
				return 0,err
		}
		return int8(res.(int64)),nil
}

//转发
func (u *DynamicList) ClickForward() (bool,error) {
	p := u.ParseDynamicID()
	if !p {
		return false,errors.New("动态id格式错误")
	}
	u.initDynamicDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET ForwardNUM = ForwardNUM + 1 WHERE DynamicID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.DynamicID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//评论
func (u *DynamicList) ClickComment() (bool,error) {
	p := u.ParseDynamicID()
	if !p {
		return false,errors.New("动态id格式错误")
	}
	u.initDynamicDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET CommentNUM = CommentNUM + 1 WHERE DynamicID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.DynamicID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

// 浏览（待优化）
func (u *DynamicList) ClickView() (bool,error) {
	p := u.ParseDynamicID()
	if !p {
		return false,errors.New("动态id格式错误")
	}
	u.initDynamicDBAndTable()
	
	sqlStr := fmt.Sprintf("UPDATE %s SET ViewNUM = ViewNUM + 1 WHERE DynamicID = ?",u.tableName)
	_, err := u.dbConn.Raw(sqlStr,u.DynamicID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//设置浏览数（待优化）
func (u *DynamicList) SetViewNum(arrDynamicID []string) (bool,error) {
	arrKey := make([]interface{},0,len(arrDynamicID))
	for _,v := range arrDynamicID {
		noSQLKey := fmt.Sprintf("%s:%s",XYLibs.NO_SQL_DYNAMIC_VIEW_NUM,v)
		arrKey = append(arrKey,noSQLKey)
	}
	err := libs.RedisDBDynamic.PipeliningINCR(arrKey)
	if err != nil {
			return false,err
	}
	return true , nil
}

//获取浏览数（待优化）
func (u *DynamicList) GetViewNum(arrDynamicID []string) (bool,[]string,error) {
	
	arrKey := make([]interface{},0)
	for _,v := range arrDynamicID {
		arrKey = append(arrKey,fmt.Sprintf("%s:%s",XYLibs.NO_SQL_DYNAMIC_VIEW_NUM,v))
	}
	data,err := libs.RedisDBDynamic.MGET(arrKey)
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


func (u *DynamicList) GenPageParam() string {
	 arr := make([]string,0,4)
	 
	if u.HomeCityID > 0 {
		arr = append(arr,fmt.Sprintf("HomeCityID = %d",u.HomeCityID))
	}
	
	if u.LivingProvinceID > 0 {
		arr = append(arr,fmt.Sprintf("LivingProvinceID = %d",u.LivingProvinceID))
	}
	
	if u.LivingCityID > 0 {
		arr = append(arr,fmt.Sprintf("LivingCityID = %d",u.LivingCityID))
	}
	res := ""
	if len(arr) > 0 {
		res = strings.Join(arr," AND ")
	}
	return res
}


//上页
func (u *DynamicList) PageUp() []orm.Params {
	u.dbConn =  ConnSlaveDBDynamic(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " AND " + strWhere
	}
	
	sqlStr := fmt.Sprintf("SELECT  %s FROM %s WHERE ID > ?  %s ORDER BY ID ASC LIMIT %d",DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex).Values(&table)
	return table
	
}

//下页
func (u *DynamicList) PageDown() []orm.Params {
	u.dbConn =  ConnSlaveDBDynamic(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " AND " + strWhere
	}
	sqlStr := ""
	if u.DataIndex > 0 {
		sqlStr = fmt.Sprintf("SELECT %s FROM %s WHERE ID < ? %s ORDER BY ID DESC LIMIT %d",DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	}else{
		sqlStr = fmt.Sprintf("SELECT %s FROM %s WHERE ID > ? %s ORDER BY ID DESC LIMIT %d",DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	}
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex).Values(&table)
	return table
}
//最后页
func (u *DynamicList) PageEnd() []orm.Params {
	u.dbConn =  ConnSlaveDBDynamic(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " WHERE " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  %s ORDER BY ID ASC LIMIT %d",DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//第一页
func (u *DynamicList) PageFirst() []orm.Params {
	u.dbConn =  ConnSlaveDBDynamic(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " WHERE " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  %s ORDER BY ID DESC LIMIT %d",DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}


