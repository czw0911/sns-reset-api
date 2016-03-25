//@Description 分页
//@Contact czw@outlook.com

package XYLibs

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"strconv"
)

type IPageing interface {
	
	//最后的表
	 GetLastTableName() error
	
	//最前的表
	 GetFirstTableName() error
	
	//检查表是否在最前或最后
	 CheckTableNameFirstOrLast() (bool,error)
	
	//解析maxid
	 ParseMaxID() bool
	
	//重新设置表
	 ResetTableName() (bool,error)
	
	//上翻页
	 PageUp() []orm.Params
	
	//下翻页
	 PageDown() []orm.Params
	
	//尾页
	 PageEnd() []orm.Params
	
	//首页
	 PageFirst() []orm.Params
	
	//获取翻页类型
	GetPageType() int8
	
	//设置翻页类型
	SetPageType(ptype int8)
	
	//获取最大id
	GetMaxID() string
	
	//设置最大id
	SetMaxID(id string)
	
	//获取日期表年月
	GetYearAndMonth() string
}

func NewPaging() *Paging {
	return new(Paging)
}

type Paging struct {}

//单表按id翻页
func (page *Paging)PageingSingleTable(actDB IPageing)(map[string]interface{},XYAPIResponse){
	
	arrData := make(map[string]interface{})
	
	if actDB.GetPageType() == PAGE_TYPE_UP {
	
		if actDB.GetMaxID() == "" {
			//首页

			dt := actDB.PageFirst()
			l := len(dt)
			if l == 0  {
				return arrData , RespStateCode["ok"]
			}
			
			arrData["MaxID"] = fmt.Sprintf("%s",dt[0]["ID"].(string))
			arrData["MinID"] = fmt.Sprintf("%s",dt[l -1]["ID"].(string))
			arrData["List"] = dt
			
		}else{
			//上翻页
			dt := actDB.PageUp()
			l := len(dt)
			if l == 0 {		
				return arrData , RespStateCode["ok"]
				
			}
			
			arrData["MaxID"] = fmt.Sprintf("%s",dt[l -1 ]["ID"].(string))
			arrData["MinID"] = fmt.Sprintf("%s",dt[0]["ID"].(string))
			arrData["List"] = dt
		}
		
	}else{

		if actDB.GetMaxID() == "" {
			
			//最后一页
			
			dt := actDB.PageEnd()
			l := len(dt)
			if l == 0  {
				return arrData , RespStateCode["ok"]
			}
			
			arrData["MaxID"] = fmt.Sprintf("%s",dt[l -1]["ID"].(string) )
			arrData["MinID"] = fmt.Sprintf("%s",dt[0]["ID"].(string))
			arrData["List"] = dt
				
		}else{
			//下翻页
			
			dt := actDB.PageDown()
			l := len(dt)
			if l == 0 {	
				return arrData , RespStateCode["ok"]
			}
			arrData["MaxID"] = fmt.Sprintf("%s",dt[0]["ID"].(string))
			arrData["MinID"] = fmt.Sprintf("%s",dt[l -1]["ID"].(string))
			arrData["List"] = dt
		}
	}
	
	return arrData , RespStateCode["ok"]
}

//根据hash值确定日期月份多表按id翻页
func (page *Paging)PaginMultiDateTableByHashIDName(actDB IPageing,dbHashIDName string,hstart,hstop int)(map[string]interface{},XYAPIResponse){
	
	arrData := make(map[string]interface{})
	
	arrTmp := make([]orm.Params,0,TABLE_LIMIT_NUM << 1)
	pageTypeBak := actDB.GetPageType()
	maxIDBak := actDB.GetMaxID()
	
	
	PAGE_RELOAD:
	
	p := actDB.ParseMaxID()
	if actDB.GetPageType() == PAGE_TYPE_UP {
		
		FIRST_PAGE:
		if !p {
			//首页
			
			actDB.GetFirstTableName()
			dt := actDB.PageFirst()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM  {
				//跳转下翻页凑行数
				actDB.SetPageType(PAGE_TYPE_DOWN)
				index := 0
				if l >0 {
					index = l -1
					actDB.SetMaxID(fmt.Sprintf("%s%s%s",actDB.GetYearAndMonth(),arrTmp[index][dbHashIDName].(string)[hstart:hstop],arrTmp[index]["ID"].(string)))
				}else{
					is,err := actDB.CheckTableNameFirstOrLast()
					if is {
						if err != nil {
							beego.Error(err)
						}
						return arrData,RespStateCode["ok"]
					}
				}			
				goto PAGE_RELOAD
			}
		
		}else{
			//上翻页
			
			PAGE_UP:
			dt := actDB.PageUp()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM  {
				
				is,err := actDB.CheckTableNameFirstOrLast()
				if !is {
					
					b,e := actDB.ResetTableName()
					if !b {
						beego.Error(e)
						p = false
						goto FIRST_PAGE
					}
					
					goto PAGE_UP
				}else{
					if err != nil {
						beego.Error(err)
					}
				}
				
			}

		}
		
	}else{
		LAST_PAGE:
		if !p {
			
			//最后一页
			
			actDB.GetLastTableName()
			dt := actDB.PageEnd()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM  {
				//跳转上翻页凑行数
				actDB.SetPageType(PAGE_TYPE_UP)
				index := 0
				if l >0 {
					index = l -1
					actDB.SetMaxID(fmt.Sprintf("%s%s%s",actDB.GetYearAndMonth(),arrTmp[index][dbHashIDName].(string)[hstart:hstop],arrTmp[index]["ID"].(string)))
				}else{
					is,err := actDB.CheckTableNameFirstOrLast()
					if is {
						if err != nil {
							beego.Error(err)
						}
						return arrData,RespStateCode["ok"]
					}
				}			
				goto PAGE_RELOAD
			}
				
		}else{
			//下翻页
			
			PAGE_DOWN:
			dt := actDB.PageDown()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM {
			
				is,err := actDB.CheckTableNameFirstOrLast()
				if !is {
					b,e := actDB.ResetTableName()
					if !b {
						beego.Error(e)
						p = false
						goto LAST_PAGE
					}	
					goto PAGE_DOWN
				}else{
					if err != nil {
						beego.Error(err)
					}
				}
			}
			
		}

	}
	
	
	datatLen := len(arrTmp)
	if datatLen == 0 {
		return arrData , RespStateCode["ok"]
	}
	index := datatLen - 1
	actDB.SetPageType(pageTypeBak)
	actDB.SetMaxID(maxIDBak)
	p = actDB.ParseMaxID()
	newMaxID := ""
	newMinID := ""
	if actDB.GetPageType() == PAGE_TYPE_UP {
		//上翻
		if !p {		
			d,_ := strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")
			newMaxID = fmt.Sprintf("%s%s%s",m,arrTmp[0][dbHashIDName].(string)[hstart:hstop],arrTmp[0]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s%s",m,arrTmp[index][dbHashIDName].(string)[hstart:hstop],arrTmp[index]["ID"].(string))
			
		}else{		
		
			d,_ := strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")	
			newMaxID = fmt.Sprintf("%s%s%s",m,arrTmp[index][dbHashIDName].(string)[hstart:hstop],arrTmp[index]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s%s",m,arrTmp[0][dbHashIDName].(string)[hstart:hstop],arrTmp[0]["ID"].(string))
		}
		
	}else{
		//下翻
		if !p {
			d,_ := strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")	
			newMaxID = fmt.Sprintf("%s%s%s",m,arrTmp[index][dbHashIDName].(string)[hstart:hstop],arrTmp[index]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s%s",m,arrTmp[0][dbHashIDName].(string)[hstart:hstop],arrTmp[0]["ID"].(string))
			
		}else{
			d,_ := strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")
			newMaxID = fmt.Sprintf("%s%s%s",m,arrTmp[0][dbHashIDName].(string)[hstart:hstop],arrTmp[0]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s%s",m,arrTmp[index][dbHashIDName].(string)[hstart:hstop],arrTmp[index]["ID"].(string))
		}
	}
	arrData["MaxID"] = newMaxID //格式：年月 + hashid + 流水号
	arrData["MinID"] = newMinID
	arrData["List"] = arrTmp
	
	return arrData,RespStateCode["ok"]
}


//日期月份多表按id翻页
func (page *Paging)PaginMultiDateTable(actDB IPageing)(map[string]interface{},XYAPIResponse){
	
	arrData := make(map[string]interface{})
	
	arrTmp := make([]orm.Params,0,TABLE_LIMIT_NUM << 1)
	pageTypeBak := actDB.GetPageType()
	maxIDBak := actDB.GetMaxID()
	
	
	PAGE_RELOAD:
	
	p := actDB.ParseMaxID()
	if actDB.GetPageType() == PAGE_TYPE_UP {
		
		FIRST_PAGE:
		if !p {
			//首页
			
			actDB.GetFirstTableName()
			dt := actDB.PageFirst()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM  {
				//跳转下翻页凑行数
				actDB.SetPageType(PAGE_TYPE_DOWN)
				index := 0
				if l >0 {
					index = l -1
					actDB.SetMaxID(fmt.Sprintf("%s%s",actDB.GetYearAndMonth(),arrTmp[index]["ID"].(string)))
				}else{
					is,err := actDB.CheckTableNameFirstOrLast()
					if is {
						if err != nil {
							beego.Error(err)
						}
						return arrData,RespStateCode["ok"]
					}
				}			
				goto PAGE_RELOAD
			}
		
		}else{
			//上翻页
			
			PAGE_UP:
			dt := actDB.PageUp()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM  {
				
				is,err := actDB.CheckTableNameFirstOrLast()
				if !is {
					
					b,e := actDB.ResetTableName()
					if !b {
						beego.Error(e)
						p = false
						goto FIRST_PAGE
					}
					
					goto PAGE_UP
				}else{
					if err != nil {
						beego.Error(err)
					}
				}
				
			}

		}
		
	}else{
		LAST_PAGE:
		if !p {
			
			//最后一页
			
			actDB.GetLastTableName()
			dt := actDB.PageEnd()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM  {
				//跳转上翻页凑行数
				actDB.SetPageType(PAGE_TYPE_UP)
				index := 0
				if l >0 {
					index = l -1
					actDB.SetMaxID(fmt.Sprintf("%s%s",actDB.GetYearAndMonth(),arrTmp[index]["ID"].(string)))
				}else{
					is,err := actDB.CheckTableNameFirstOrLast()
					if is {
						if err != nil {
							beego.Error(err)
						}
						return arrData,RespStateCode["ok"]
					}
				}			
				goto PAGE_RELOAD
			}
				
		}else{
			//下翻页
			
			PAGE_DOWN:
			dt := actDB.PageDown()
			if len(dt) > 0 {
				arrTmp = append(arrTmp,dt...)
			}
			l := len(arrTmp)
			if l  < TABLE_LIMIT_NUM {
			
				is,err := actDB.CheckTableNameFirstOrLast()
				if !is {
					b,e := actDB.ResetTableName()
					if !b {
						beego.Error(e)
						p = false
						goto LAST_PAGE
					}	
					goto PAGE_DOWN
				}else{
					beego.Error(err)
				}
			}
			
		}

	}
	
	
	datatLen := len(arrTmp)
	if datatLen == 0 {
		return arrData , RespStateCode["ok"]
	}
	index := datatLen - 1
	actDB.SetPageType(pageTypeBak)
	actDB.SetMaxID(maxIDBak)
	p = actDB.ParseMaxID()
	newMaxID := ""
	newMinID := ""
	if actDB.GetPageType() == PAGE_TYPE_UP {
		//上翻
		if !p {		
			d,_ := strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")
			newMaxID = fmt.Sprintf("%s%s",m,arrTmp[0]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s",m,arrTmp[index]["ID"].(string))
			
		}else{		
		
			d,_ := strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")	
			newMaxID = fmt.Sprintf("%s%s",m,arrTmp[index]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s",m,arrTmp[0]["ID"].(string))
		}
		
	}else{
		//下翻
		if !p {
			d,_ := strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")	
			newMaxID = fmt.Sprintf("%s%s",m,arrTmp[index]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s",m,arrTmp[0]["ID"].(string))
			
		}else{
			d,_ := strconv.ParseInt(arrTmp[0]["PostTime"].(string),10,64)
			m := time.Unix(d,0).Format("200601")
			newMaxID = fmt.Sprintf("%s%s",m,arrTmp[0]["ID"].(string))
			
			d,_ = strconv.ParseInt(arrTmp[index]["PostTime"].(string),10,64)
			m = time.Unix(d,0).Format("200601")
			newMinID = fmt.Sprintf("%s%s",m,arrTmp[index]["ID"].(string))
		}
	}
	arrData["MaxID"] = newMaxID
	arrData["MinID"] = newMinID
	arrData["List"] = arrTmp
	
	return arrData,RespStateCode["ok"]
}