//@Description 注册数统计
//@Contact czw@outlook.com


package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
)

type RegNUM struct {
	
	tableName []string //表名
	
	UID uint32 //用户id
	
	HomeProvinceID int //  家乡省id

	HomeCityID int //  家乡城市id

	HomeDistrictID int //  家乡区县id

	LivingProvinceID int //  居住地省id

	LivingCityID int //  居住地城市id

	LivingDistrictID int //  居住地区县id
	
	RegType int8 // 注册类型,1:手机 2：微博 3:微信
	
	RegisterTime int64 //注册时间
}

func (r *RegNUM) initTable() {
		r.tableName = make([]string,0,10)
		var i int64
		for i = 2015 ; i <= r.RegisterTime ; i++ {
			r.tableName = append(r.tableName,fmt.Sprintf("reg_logs_%d",i))
		} 
}

//获取每天注册人数
func (r *RegNUM) GetEveryDay()[]orm.Params{
	db := ConnLogsDB()
	r.initTable()
	arrSql := make([]string,0,len(r.tableName))
	for _,v := range r.tableName {
		arrSql = append(arrSql,"SELECT ID ,RegisterTime FROM  "+ v )
	}
	strSql := "SELECT count(ID) Number,from_unixtime(RegisterTime,'%Y%-%m-%d') RegDate FROM ("+ strings.Join(arrSql,"  UNION ALL ")+")t"
	var res []orm.Params	
	_, _ = db.Raw(strSql).Values(&res)	
	return res
}


//按家乡分类获取所有注册人数
func (r *RegNUM) GetAllByHome()[]orm.Params{
	db := ConnLogsDB()
	r.initTable()
	arrSql := make([]string,0,len(r.tableName))
	for _,v := range r.tableName {
		arrSql = append(arrSql,"SELECT ID , HomeProvinceID FROM  "+ v )
	}
	strSql := "SELECT HomeProvinceID,count(ID) Number FROM ("+ strings.Join(arrSql,"  UNION ALL ")+")t  GROUP BY HomeProvinceID"
	var res []orm.Params	
	_, _ = db.Raw(strSql).Values(&res)	
	return res
}




