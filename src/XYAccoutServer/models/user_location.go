////@Description 用户位置信息
////@Contact czw@outlook.com

package models

//import (
//	"github.com/astaxie/beego/orm"
//	"stathat.com/c/consistent"
//	"fmt"
//	"XYAPIServer/XYLibs"
//	"XYAPIServer/XYAccoutServer/libs"
//	"strings"
//)


//var (
//	userLocationTableNameHash = consistent.New()
//)

//type UserLocation struct {

//	  tableName string //hash到的表名
	
//	  dbConn  orm.Ormer //hash到的数据库
	  
//	  UID uint32 //用户id
	
//	  HomeProvinceID int    //  家乡省id  
	
//	  HomeProvinceName  string   //  家乡省名  
	
//	  HomeCityID  int   //  家乡城市id  
	
//	  HomeCityName  string  //  家乡城市名  
	
//	  HomeDistrictID  int  //  家乡区县id  
	
//	  HomeDistrictName   string //  家乡区县名  
	
//	  LivingProvinceID  int  //  居住地省id  
	
//	  LivingProvinceName  string  //  居住地省名  
	
//	  LivingCityID   int //  居住地城市id  
	
//	  LivingCityName   string //  居住地城市名 
	 
//	  LivingDistrictID  int //  居住地区县id  
	
//	  LivingDistrictName  string   //  居住地区县名  
	
//	  ProfessionID int //职业所属行业id
	
//	  JobID int //职业id
	
//	  HomeVoice string // 乡音音频地址
	
//	  VoiceLen int  // 乡音音频长度 
	
//	  NorthLatitude  int  //  北纬  
	
//	  EastLongtude  int  //  东经  
	
//	  MaxID string //当前列表最大id  
	
//	  PageType int8 //翻页类型 1,上翻；2，翻页

		
//}

//func (u *UserLocation) initDBAndTable(){
//		u.dbConn =  ConnMasterDB(uint32(u.HomeProvinceID))
//		id := fmt.Sprintf("%d",u.HomeProvinceID)
//		u.tableName,_ = userLocationTableNameHash.Get(id)
//}



//func (u *UserLocation) IsUIDExist() (bool,error) {
//	u.initDBAndTable()
//	sqlStr := fmt.Sprintf("SELECT ID FROM %s WHERE UID = ? LIMIT 1",u.tableName)
//	var res []orm.Params
//	_, err := u.dbConn.Raw(sqlStr,u.UID).Values(&res)
//	if err != nil {
//		return true,err
//	}
//	if len(res) == 0 {
//		return false, nil
//	}
//	return true,nil
//}

//func (u * UserLocation) Reg() (bool,error) {
//		u.initDBAndTable()
//		sqlStr := fmt.Sprintf("REPLACE INTO %s(UID,HomeProvinceID,HomeCityID,HomeDistrictID,LivingProvinceID,LivingCityID,LivingDistrictID,HomeVoice,VoiceLen) VALUES(?,?,?,?,?,?,?,?,?)",u.tableName)
//		_,err := u.dbConn.Raw(sqlStr,u.UID,u.HomeProvinceID,u.HomeCityID,u.HomeDistrictID,u.LivingProvinceID,u.LivingCityID,u.LivingDistrictID,u.HomeVoice,u.VoiceLen).Exec()
//		if err != nil {
//			return false,err
//		}
//		return true,nil
//}

//func (u * UserLocation) DelLocationByUID() (bool,error) {
//		u.initDBAndTable()
//		sqlStr := fmt.Sprintf("DELETE FROM  %s  WHERE UID = ?)",u.tableName)
//		_,err := u.dbConn.Raw(sqlStr,u.UID).Exec()
//		if err != nil {
//			return false,err
//		}
//		return true,nil
//}


//func (u * UserLocation) GetLocationByUID() []orm.Params {
//		u.initDBAndTable()
//		sqlStr := fmt.Sprintf("SELECT HomeProvinceID, HomeCityID, HomeDistrictID, LivingProvinceID, LivingCityID, LivingDistrictID, HomeVoice, VoiceLen ,ProfessionID ,JobID  FROM %s WHERE UID = ?  LIMIT 1",u.tableName)
//		var res []orm.Params
//		_,_ = u.dbConn.Raw(sqlStr,u.UID).Values(&res)
//		return res
//}

//func (u * UserLocation) UpdateHomeVoice() (bool,error) {
	
//		u.initDBAndTable()
//		sqlStr := fmt.Sprintf("UPDATE  %s  SET  HomeVoice = ? , VoiceLen = ?  WHERE UID = ?",u.tableName)
//		_,err := u.dbConn.Raw(sqlStr,u.HomeVoice,u.VoiceLen,u.UID).Exec()
//		if err != nil {
//			return false,err
//		}
//		return true,nil
//}
////保存乡音文件到已录列表
//func (u * UserLocation) SaveHomeVoiceToRevordList() (bool,error) {
	
//		noSQLKey := fmt.Sprintf("%s",XYLibs.NO_SQL_USER_ALREADY_RECORD_VOICE_LIST)	
//		_,err := libs.RedisDBUser.SADD(noSQLKey,fmt.Sprintf("%d",u.UID))
//		if err != nil {
//				return false,err
//		}
//		return true , nil
//}

//func (u * UserLocation) GenerateHometownGroup() ([]string,error) {
//		db := ConnCommonDB()
//		gID := make([]string,3)
//		sqlStr := fmt.Sprintf("REPLACE INTO %s(GroupID,GroupName) VALUES(?,?)","group_list")
//		gID[0] = fmt.Sprintf("%s%d000000",XYLibs.GROUP_TYPE_BY_LIVING,u.HomeProvinceID)
//		aname :=fmt.Sprintf("%s人大家庭",u.HomeProvinceName)
//		p,err := db.Raw(sqlStr).Prepare()
//		if err != nil {
//			return gID,err
//		}
//		defer p.Close()
//		_,err = p.Exec(gID[0],aname)
//		if err != nil {
//			return gID,err
//		}
		
//		gID[1] = fmt.Sprintf("%s%d%d",XYLibs.GROUP_TYPE_BY_LIVING,u.HomeProvinceID,u.LivingCityID)
//		bname :=fmt.Sprintf("%s人在%s",u.HomeProvinceName,u.LivingCityName)
//		_,err = p.Exec(gID[1],bname)
//		if err != nil {
//			return gID,err
//		}
		
//		gID[2] = fmt.Sprintf("%s%d%d",XYLibs.GROUP_TYPE_BY_LIVING,u.HomeCityID,u.LivingCityID)
//		cname :=fmt.Sprintf("%s人在%s",u.HomeCityName,u.LivingCityName)
//		_,err = p.Exec(gID[2],cname)
//		if err != nil {
//			return gID,err
//		}
		
//		return gID,nil
//}

//func (u *UserLocation) GenPageParam() string {
//	 arr := make([]string,0,4)
//	 if u.HomeProvinceID > 0 {
//		arr = append(arr,fmt.Sprintf("HomeProvinceID = %d",u.HomeProvinceID))
//	 }
	
//	if u.HomeCityID > 0 {
//		arr = append(arr,fmt.Sprintf("HomeCityID = %d",u.HomeCityID))
//	}
	
//	if u.LivingProvinceID > 0 {
//		arr = append(arr,fmt.Sprintf("LivingProvinceID = %d",u.LivingProvinceID))
//	}
	
//	if u.LivingCityID > 0 {
//		arr = append(arr,fmt.Sprintf("LivingCityID = %d",u.LivingCityID))
//	}
//	res := ""
//	if len(arr) > 0 {
//		res = strings.Join(arr," AND ")
//	}
//	return res
//}


//	//最后的表
//func (u *UserLocation) GetLastTableName() error{
//	return nil
//}
	
//	//最前的表
//func (u *UserLocation)	GetFirstTableName() error{
//	return nil
//}
	
//	//检查表是否在最前或最后
//func (u *UserLocation) CheckTableNameFirstOrLast() (bool,error){
//	return true,nil
//}
	
//	//解析maxid
//func (u *UserLocation) ParseMaxID() bool{
//	return false
//}
	
//	//重新设置表
//func (u *UserLocation)	ResetTableName() (bool,error){
//	return false,nil
//}

////获取日期表年月
//func (u *UserLocation)	GetYearAndMonth() string {
//	return ""
//}

////设置翻页类型
//func (u *UserLocation)	SetPageType(ptype int8) {
//	u.PageType = ptype
//}

////设置最大id
//func (u *UserLocation)	SetMaxID(id string) {
//	u.MaxID = id
//}

////获取翻页类型
//func (u *UserLocation)	GetPageType() int8 {
//	return u.PageType
//}

////获取最大id
//func (u *UserLocation)	GetMaxID() string {
//	return u.MaxID
//}

////上页
//func (u *UserLocation) PageUp() []orm.Params {
//	u.initDBAndTable()
//	strWhere := u.GenPageParam()
//	if strWhere != "" {
//		strWhere = " AND " + strWhere
//	}
//	sqlStr := fmt.Sprintf("SELECT  ID, UID  FROM %s WHERE  ID > ? %s  ORDER BY ID ASC LIMIT %d",u.tableName,strWhere,TABLE_LIMIT_NUM)
//	var table []orm.Params
//	u.dbConn.Raw(sqlStr,u.MaxID).Values(&table)
//	return table
	
//}

////下页
//func (u *UserLocation) PageDown() []orm.Params {
//	u.initDBAndTable()
	
//	strWhere := u.GenPageParam()
//	if strWhere != "" {
//		strWhere = " AND " + strWhere
//	}
//	sqlStr := fmt.Sprintf("SELECT ID, UID  FROM %s WHERE ID < ? %s  ORDER BY ID DESC LIMIT %d",u.tableName,strWhere,TABLE_LIMIT_NUM)

//	var table []orm.Params
//	u.dbConn.Raw(sqlStr,u.MaxID).Values(&table)
//	return table
//}

////最后页
//func (u *UserLocation) PageEnd() []orm.Params {
//	u.initDBAndTable()
//	strWhere := u.GenPageParam()
//	if strWhere != "" {
//		strWhere = " WHERE " + strWhere
//	}
//	sqlStr := fmt.Sprintf("SELECT ID, UID  FROM %s  %s  ORDER BY ID ASC LIMIT %d",u.tableName,strWhere,TABLE_LIMIT_NUM)
//	var table []orm.Params
//	u.dbConn.Raw(sqlStr).Values(&table)
//	return table
//}

////首页
//func (u *UserLocation) PageFirst() []orm.Params {
//	u.initDBAndTable()
//	strWhere := u.GenPageParam()
//	if strWhere != "" {
//		strWhere = " WHERE " + strWhere
//	}
//	sqlStr := fmt.Sprintf("SELECT ID, UID  FROM %s  %s   ORDER BY ID Desc LIMIT %d",u.tableName,strWhere,TABLE_LIMIT_NUM)
//	var table []orm.Params
//	u.dbConn.Raw(sqlStr).Values(&table)
//	return table
//}



//func init(){
//	for i:=0 ; i < USER_HASH_TABLE_NUM ; i++ {
//		tn := fmt.Sprintf("user_location_%d",i)
//		userLocationTableNameHash.Add(tn)
//	}
//}