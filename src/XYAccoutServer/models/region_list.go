//@Description 地区列表
//@Contact czw@outlook.com


package models

import (
	"github.com/astaxie/beego/orm"
)

type RegionList struct {
	RegionId int
	RegionName string
	List []*Province
}

//省
type Province struct {
	RegionId string 	`json:"RegionID"`
	RegionName string	`json:"RegionName"`
	RegionSub []*City	`json:"RegionSub"`
	
}
//市
type City struct {
	CityID string			`json:"RegionID"`
	CityName string			`json:"RegionName"`
	CitySub []*District		`json:"RegionSub"`
}
//区县
type District struct {
	DistrictID string		`json:"RegionID"`
	DistrictName string		`json:"RegionName"`
}


func (r *RegionList) GetProvince()[]orm.Params{
	db := ConnCommonDB()
	var res []orm.Params
	if r.RegionId != 0 {
		_, _ = db.Raw("SELECT ProvinceID AS RegionID,ProvinceName AS RegionName  FROM  region_province WHERE ProvinceID = ?",r.RegionId).Values(&res)
	}else{
		_, _ = db.Raw("SELECT ProvinceID AS RegionID,ProvinceName AS RegionName  FROM  region_province").Values(&res)

	}
	return res
}


func (r *RegionList) GetCity()[]orm.Params{
	db := ConnCommonDB()
	var res []orm.Params
	if r.RegionId != 0 {
			_, _ = db.Raw("SELECT CityID AS RegionID,CityName AS RegionName  FROM  region_city WHERE ProvinceID = ?",r.RegionId).Values(&res)
	}else{
			_, _ = db.Raw("SELECT CityID AS RegionID,CityName AS RegionName  FROM  region_city ").Values(&res)

	}
	return res
}


func (r *RegionList) GetCityByID()[]orm.Params{
	db := ConnCommonDB()
	var res []orm.Params
	_, _ = db.Raw("SELECT CityID AS RegionID,CityName AS RegionName  FROM  region_city WHERE CityID = ?",r.RegionId).Values(&res)
	return res
}


func (r *RegionList) GetDistrict()[]orm.Params{
	db := ConnCommonDB()
	var res []orm.Params
	if r.RegionId != 0 {
			_, _ = db.Raw("SELECT DistrictID AS RegionID,DistrictName AS RegionName   FROM  region_district WHERE CityID = ?",r.RegionId).Values(&res)
	}else{
			_, _ = db.Raw("SELECT DistrictID AS RegionID,DistrictName AS RegionName   FROM  region_district ").Values(&res)

	}
	return res
}

func (r *RegionList) GetAll() []*Province {
	var p []orm.Params
	var c []orm.Params
	var d []orm.Params
	db := ConnCommonDB()
	_, _ = db.Raw("SELECT ProvinceID AS RegionID,ProvinceName AS RegionName  FROM  region_province ").Values(&p)
	_, _ = db.Raw("SELECT CityID AS RegionID,CityName AS RegionName,ProvinceID  FROM  region_city ").Values(&c)
	_, _ = db.Raw("SELECT DistrictID AS RegionID,DistrictName AS RegionName,CityID  FROM  region_district ").Values(&d)
	
	
	province := make([]*Province,0)
	
	for _,v := range p {
		pe := new(Province)
		pe.RegionId = v["RegionID"].(string)
		pe.RegionName = v["RegionName"].(string)
		pe.RegionSub = make([]*City,0)
		for _,x := range c {
				if pe.RegionId == x["ProvinceID"].(string) {
					cy := new(City)
					cy.CityID = x["RegionID"].(string)
					cy.CityName = x["RegionName"].(string)
					cy.CitySub = make([]*District,0)
					 
					for _,y := range d {
						if cy.CityID == y["CityID"].(string) {
							dt := new(District)
							dt.DistrictID = y["RegionID"].(string)
							dt.DistrictName = y["RegionName"].(string)
							cy.CitySub = append(cy.CitySub,dt)
						}
					}
					pe.RegionSub = append(pe.RegionSub, cy)
				}
		}
		province = append(province,pe)
	}
	

	return province
}


