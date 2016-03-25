//用户动态列表
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"strings"
	"strconv"
	"XYAPIServer/XYLibs"
	"XYAPIServer/XYDynamicServer/libs"
	"errors"
)

const USER_DYNAMIC_TABLE_SELECT_COLUMN_NAME = " ID,UID,DynamicID,DynamicContent,ViewNUM,CommentNum,ForwardNum,GoodNUM,Images,Voices,VoiceLen,PostTime " 

type UserDynamic struct {

	  tableName string //hash到的表名
	
	  dbConn  orm.Ormer //hash到的数据库
	
	  UID uint32 //用户id
	
	  ViewUID uint32 //查看的用户id
	
	  HomeProvinceID int    //  家乡省id  	
	
	  DynamicID string //动态id	
	  
	  YearAndMonth string //年月
	
	  MaxID string //当前列表最大id 格式:YearAndMonth - ID 
	
	  PageType int8 //翻页类型 1,上翻；2，翻页
	
	  DataIndex int64 //列表数据当前索引
	
	  prevTableName string //上次的表面
	
	
}

//月表按家乡省份hash
func (u *UserDynamic) initDynamicDBAndTable() {
		u.dbConn =  ConnMasterDB(u.HomeProvinceID)
		if u.YearAndMonth == "" {
			u.YearAndMonth = time.Now().Format("200601")
		}
		u.tableName = fmt.Sprintf("dynamic_%d_%s",u.HomeProvinceID,u.YearAndMonth)	
}

func (u *UserDynamic) ParseDynamicID() bool {
	
	//fmt.Printf("%#v\n",arr)
	if len(u.DynamicID) != 25 {
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

// 浏览（待优化）
func (u *UserDynamic) ClickView() (bool,error) {
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

func (u *UserDynamic) ParseMaxID() bool{
	
	u.DataIndex = 0
	if len(u.MaxID) < 13 {
	
		return false
	}
	
	tableIndex,err := strconv.Atoi(u.MaxID[0:6])
	if err != nil {
		return false
	}
	u.YearAndMonth = fmt.Sprintf("%d",tableIndex)
	
	u.HomeProvinceID,err = strconv.Atoi(u.MaxID[6:12])
	if err != nil {
		return false
	}
	
	id ,err  := strconv.ParseInt(u.MaxID[12:],10,64)
	if err != nil {
		return false
	}
	u.DataIndex = id
	u.initDynamicDBAndTable()
	return true
}

//是否点赞
func (u *UserDynamic) IsClickGood() (int8,error) {
		noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_CLICK_GOOD_DYNAMIC_ID,u.UID)
		val := fmt.Sprintf("%s",u.DynamicID)
		res,err := libs.RedisDBDynamic.SISMEMBER(noSQLKey,val)
		if err != nil || res == nil {
				return 0,err
		}
		return int8(res.(int64)),nil
}

//表是否为空
func (u *UserDynamic) IsEmptyTableName()(bool,error){
	
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX,u.ViewUID)
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
func (u *UserDynamic)	SetPageType(ptype int8) {
	u.PageType = ptype
}

//设置最大id
func (u *UserDynamic)	SetMaxID(id string) {
	u.MaxID = id
}

//获取日期表年月
func (u *UserDynamic)	GetYearAndMonth() string {
	return u.YearAndMonth
}

//获取翻页类型
func (u *UserDynamic)	GetPageType() int8 {
	return u.PageType
}

//获取最大id
func (u *UserDynamic)	GetMaxID() string {
	return u.MaxID
}

//最前的表
func (u *UserDynamic) GetFirstTableName() error {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX,u.ViewUID)	
	tname,err := libs.RedisDBDynamic.ZREVRANGE(noSQLKey,"0","0")
	if err != nil {
			return err
	}
	t := tname.([]interface{})
	u.tableName = string(t[0].([]uint8))
	a := strings.Split(u.tableName,"_")
	u.HomeProvinceID,_ = strconv.Atoi(a[1])
	u.YearAndMonth = a[2]
	return nil
}

//最后的表
func (u *UserDynamic) GetLastTableName() error{
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX,u.ViewUID)
	
	tname,err := libs.RedisDBDynamic.ZREVRANGE(noSQLKey,"-1","-1")
	if err != nil {
			return err
	}
	t := tname.([]interface{})
	u.tableName = string(t[0].([]uint8))
	a := strings.Split(u.tableName,"_")
	u.HomeProvinceID,_ = strconv.Atoi(a[1])
	u.YearAndMonth = a[2]
	return nil
}

//检查活动表是否在最前或最后
func (u *UserDynamic) CheckTableNameFirstOrLast() (bool,error){
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX,u.ViewUID)
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
func (u *UserDynamic) ResetTableName() (bool,error) {
	noSQLKey := fmt.Sprintf("%s:%d",XYLibs.NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX,u.ViewUID)
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
	u.HomeProvinceID,_ = strconv.Atoi(u.tableName[8:14])
	u.YearAndMonth = u.tableName[15:]
	if u.prevTableName != u.tableName {
		u.DataIndex = 0
	}
	
	return true , nil
}




func (u *UserDynamic) GenPageParam() string {
	
		return fmt.Sprintf("UID = %d",u.ViewUID)

}


//上页
func (u *UserDynamic) PageUp() []orm.Params {
	
	u.dbConn =  ConnSlaveDB(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " AND " + strWhere
	}
	
	sqlStr := fmt.Sprintf("SELECT  %s FROM %s WHERE ID > ?  %s ORDER BY ID ASC LIMIT %d",USER_DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex).Values(&table)
	return table
	
}

//下页
func (u *UserDynamic) PageDown() []orm.Params {
	
	u.dbConn =  ConnSlaveDB(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " AND " + strWhere
	}
	sqlStr := ""
	if u.DataIndex > 0 {
		sqlStr = fmt.Sprintf("SELECT %s FROM %s WHERE ID < ? %s ORDER BY ID DESC LIMIT %d",USER_DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	}else{
		sqlStr = fmt.Sprintf("SELECT %s FROM %s WHERE ID > ? %s ORDER BY ID DESC LIMIT %d",USER_DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	}
	var table []orm.Params
	u.dbConn.Raw(sqlStr,u.DataIndex).Values(&table)
	return table
}
//最后页
func (u *UserDynamic) PageEnd() []orm.Params {
	
	u.dbConn =  ConnSlaveDB(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " WHERE " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  %s ORDER BY ID ASC LIMIT %d",USER_DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}

//第一页
func (u *UserDynamic) PageFirst() []orm.Params {
	
	u.dbConn =  ConnSlaveDB(u.HomeProvinceID)
	strWhere := u.GenPageParam()
	if strWhere != "" {
		strWhere = " WHERE " + strWhere
	}
	sqlStr := fmt.Sprintf("SELECT %s FROM %s  %s ORDER BY ID DESC LIMIT %d",USER_DYNAMIC_TABLE_SELECT_COLUMN_NAME,u.tableName,strWhere,XYLibs.TABLE_LIMIT_NUM)
	var table []orm.Params
	u.dbConn.Raw(sqlStr).Values(&table)
	return table
}


