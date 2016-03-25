//修改用户信息
package controllers

import (
	"XYAPIServer/XYRobot/libs"
	"XYAPIServer/XYRobot/models"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"time"
	"strconv"
	"strings"
	"net/url"
)

type Size interface {
    Size() int64
}

type UpdateUserInfoController struct {
	BaseController
}

func (u *UpdateUserInfoController) Post() {

	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()
}

func (u *UpdateUserInfoController) Get() {
	detailDB := new(models.UserDetailInfo)
	resp := XYLibs.RespStateCode["ok"]
	uid, _ := u.GetInt64("UID")
	detailDB.UID = uint32(uid)
	detailDB.HomeProvinceID, _ = u.GetInt("HomeProvinceID")
	detailDB.HomeCityID, _ = u.GetInt("HomeCityID")
	detailDB.HomeDistrictID, _ = u.GetInt("HomeDistrictID")
	detailDB.LivingProvinceID, _ = u.GetInt("LivingProvinceID")
	detailDB.LivingCityID, _ = u.GetInt("LivingCityID")
	detailDB.LivingDistrictID, _ = u.GetInt("LivingDistrictID")
	detailDB.NickName,_ = url.QueryUnescape(u.GetString("NickName"))
	detailDB.ProfessionID, _ = u.GetInt("ProfessionID")
	detailDB.JobID, _ = u.GetInt("JobID")
	detailDB.Gender, _ = u.GetInt("Gender")
	detailDB.Birthday, _ = u.GetInt("Birthday")
	detailDB.TagID,_ = url.QueryUnescape(u.GetString("TagID"))
	detailDB.DiySign,_ = url.QueryUnescape(u.GetString("DiySign"))
	imgFile, imgHeader, _ := u.Ctx.Request.FormFile("Avatar")


	if imgSizeInterface, ok := imgFile.(Size); ok {
		beego.Info("上传img文件的大小为: %d mime:%s", imgSizeInterface.Size(), imgHeader.Header.Get("Content-Type"))

		if imgSizeInterface.Size() > XYLibs.UPLOAD_FILE_MAX_SIZE {
			resp = XYLibs.RespStateCode["upload_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}

		ym := time.Now().Format("200601")
		fName, fPath := XYLibs.GetUserUpLoadFileNameAndPath(detailDB.UID, ym, "png")
		_, err := os.Stat(fPath)
		if err != nil {
			err = os.MkdirAll(beego.AppConfig.String("upload_path")+"/"+fPath, 0666)
			if err != nil {
				beego.Error(err)
			}

		}
		saveName := fmt.Sprintf("%s/%s/%s", beego.AppConfig.String("upload_path"), fPath, fName)
		saveRes := u.SaveToFile("Avatar", saveName)
		if saveRes != nil {
			beego.Error(saveRes)
			resp = XYLibs.RespStateCode["upload_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
		detailDB.Avatar = fName
	}
	acatar := XYLibs.NewUserAvatar(libs.RedisDBUser)
	data, err := acatar.Get(detailDB.UID)
	if err != nil {
		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_homevoice_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}

	columnName := make([]string, 0, 15)
	columnValue := make([]interface{}, 0, 15)
	oldNickName := data.NickName
	
	if detailDB.UID != 0 {
		columnName = append(columnName, "UID = ?")
		columnValue = append(columnValue,detailDB.UID)
	}
	
	if detailDB.NickName != "" {
		columnName = append(columnName, "NickName = ?")
		columnValue = append(columnValue,detailDB.NickName)
		data.NickName =  detailDB.NickName
	}

	if detailDB.Gender != 0 {
		columnName = append(columnName, "Gender = ?")
		columnValue = append(columnValue,detailDB.Gender)
		data.Gender =  detailDB.Gender
	}

	if detailDB.Birthday != 0 {
		columnName = append(columnName, "Birthday = ?")
		columnValue = append(columnValue,detailDB.Birthday)
		data.Birthday =  detailDB.Birthday
	}

	if detailDB.ProfessionID != 0 {
		columnName = append(columnName, "ProfessionID = ?")
		columnValue = append(columnValue,detailDB.ProfessionID)
		data.ProfessionID =  detailDB.ProfessionID
	}

	if detailDB.JobID != 0 {
		columnName = append(columnName, "JobID = ?")
		columnValue = append(columnValue,detailDB.JobID)
		data.JobID =  detailDB.JobID
	}

	if detailDB.HomeProvinceID != 0 {
		columnName = append(columnName, "HomeProvinceID = ?")
		columnValue = append(columnValue,detailDB.HomeProvinceID)
		data.HomeProvinceID =  detailDB.HomeProvinceID
	}

	if detailDB.HomeCityID != 0 {
		columnName = append(columnName, "HomeCityID = ?")
		columnValue = append(columnValue,detailDB.HomeCityID)
		data.HomeCityID =  detailDB.HomeCityID
	}

	if detailDB.HomeDistrictID != 0 {
		columnName = append(columnName, "HomeDistrictID = ?")
		columnValue = append(columnValue,detailDB.HomeDistrictID)
		data.HomeDistrictID =  detailDB.HomeDistrictID
	}

	if detailDB.LivingProvinceID != 0 {
		columnName = append(columnName, "LivingProvinceID = ?")
		columnValue = append(columnValue,detailDB.LivingProvinceID)
		data.LivingProvinceID =  detailDB.LivingProvinceID
	}

	if detailDB.LivingCityID != 0 {
		columnName = append(columnName, "LivingCityID = ?")
		columnValue = append(columnValue,detailDB.LivingCityID)
		data.LivingCityID =  detailDB.LivingCityID
	}

	if detailDB.LivingDistrictID != 0 {
		columnName = append(columnName, "LivingDistrictID = ?")
		columnValue = append(columnValue,detailDB.LivingDistrictID)
		data.LivingDistrictID =  detailDB.LivingDistrictID
	}


	if detailDB.DiySign != "" {
		columnName = append(columnName, "DiySign = ?")
		columnValue = append(columnValue,detailDB.DiySign)
		
	}

	if detailDB.Avatar != "" {
		columnName = append(columnName, "Avatar = ?")
		columnValue = append(columnValue,detailDB.Avatar)
		data.Avatar = detailDB.Avatar
	}

	if detailDB.TagID != "" {
		columnName = append(columnName, "TagID = ?")
		columnValue = append(columnValue,detailDB.TagID)
		data.TagsID =  strings.Split(detailDB.TagID,",")
	}

	if len(columnName) == 0 {
			resp = XYLibs.RespStateCode["user_update_user_info_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
	}
	data.SetRedisConnect(libs.RedisDBUser)
	_, err = data.Set()
	if err != nil {

		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_user_info_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	if detailDB.HomeVoice != "" {
		_, err = detailDB.SaveHomeVoiceToRevordList()
		if err != nil {
			beego.Error(err)
			resp = XYLibs.RespStateCode["user_update_user_info_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
	}
	
	objUserBase := new(models.UserBase)
	objUserBase.UID = detailDB.UID
	userData := objUserBase.GeteHomeProvinceIDByUID()
	if  len(userData) == 0 {
		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_user_info_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	oldHomeProvinceID,ok := userData[0]["HomeProvinceID"].(string)
	if !ok {
		resp = XYLibs.RespStateCode["user_update_user_info_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	
	copyObject := new(models.UserDetailInfo)
	OldHomePID,_ := strconv.Atoi(oldHomeProvinceID)
	copyObject.HomeProvinceID = OldHomePID
	OldDbName := copyObject.GetHashDBName()
	OldTabName := copyObject.GetHashTableName()
	NewDbName := detailDB.GetHashDBName()
	NewTabName := detailDB.GetHashTableName()
	
	if detailDB.HomeProvinceID != 0 && (OldDbName != NewDbName ||  OldTabName != NewTabName){

		//转移到新表
		copyObject.UID = detailDB.UID	
		copyData := copyObject.GeteDetailInfoUID()
		if len(copyData) > 0 {
			id := fmt.Sprintf("%d", detailDB.UID)
			colName :=make([]string,0,18)
			colName = append(colName,"UID")
								
			colVal  := make([]interface{},0,18)
			colVal = append(colVal,id)
			for k,v := range copyData[0] {
				if v != nil {
				    colName = append(colName,k)
					colVal = append(colVal,v)
				}
			}
			_, err = detailDB.Replace(colName, colVal)
			if err != nil {
				beego.Error(err)
				resp = XYLibs.RespStateCode["user_update_user_info_fail"]
				u.Data["json"] = resp
				u.ServeJson()
				return
			}
		}
		//更新家乡省
		objUserBase.HomeProvinceID = detailDB.HomeProvinceID
		_,err := objUserBase.SeteHomeProvinceIDByUID()
		if err != nil {
			beego.Error(err)
			resp = XYLibs.RespStateCode["user_update_user_info_fail"]
			u.Data["json"] = resp
			u.ServeJson()
			return
		}
	}
	
	
	_, err = detailDB.Edit(columnName, columnValue)	
	if err != nil {
		if detailDB.HomeProvinceID != 0 &&  (OldDbName != NewDbName ||  OldTabName != NewTabName) {
			detailDB.DelUID()
		}
		objUserBase.HomeProvinceID = OldHomePID
		objUserBase.SeteHomeProvinceIDByUID()
		data.HomeProvinceID = OldHomePID
		data.NickName = oldNickName
		data.Set()
		beego.Error(err)
		resp = XYLibs.RespStateCode["user_update_user_info_fail"]
		u.Data["json"] = resp
		u.ServeJson()
		return
	}
	if detailDB.HomeProvinceID != 0 &&  (OldDbName != NewDbName ||  OldTabName != NewTabName) {
		copyObject.DelUID()	
	}

	u.Data["json"] = resp
	u.ServeJson()

}
