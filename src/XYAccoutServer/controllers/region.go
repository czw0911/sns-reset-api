package controllers

import (
	"XYAPIServer/XYAccoutServer/models"
	//"XYAPIServer/XYAccoutServer/libs"
	//"fmt"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs"
)



type RegionController struct {
	BaseController
}

func (u *RegionController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}

//获取地区列表
func (u *RegionController) Get() {
	
	resp := XYLibs.RespStateCode["nil"]
	RegionDB := new(models.RegionList)
	regionType,_ := u.GetInt("RegionType")
	
	regionID,_ := u.GetInt("RegionID")
//	sign := u.GetString("Sign")
	
//	p := u.Ctx.Request.URL.Query()
//	p.Del("Sign")
//	content := p.Encode()
//	auth := XYLibs.CheckSign(content,sign)
//	if !auth {
//		resp = XYLibs.RespStateCode["sign_error"]
//		u.Data["json"] = resp
//		u.ServeJson()
//		return 
//	}
	//ip := u.Ctx.Input.IP()
	resp =   XYLibs.RespStateCode["ok"]
	switch regionType {
		case 0 :
			if len(RegionDB.List) == 0 {
				RegionDB.List = RegionDB.GetAll()
			}
			if len(RegionDB.List) == 0 {
				resp =   XYLibs.RespStateCode["region_null"]
			}
			resp.Info = RegionDB.List
		case 1 :
			RegionDB.RegionId = regionID
			resp.Info = RegionDB.GetProvince()
		case 2 :
			RegionDB.RegionId = regionID
			resp.Info = RegionDB.GetCity()
		case 3 :
			RegionDB.RegionId = regionID
			resp.Info = RegionDB.GetDistrict()
	}
	if _,ok := resp.Info.([]orm.Params) ; ok {
		if len(resp.Info.([]orm.Params)) == 0 {
			resp =   XYLibs.RespStateCode["region_null"]
		}
		
	}
	u.Data["json"] = resp
	u.ServeJson()
	
}


