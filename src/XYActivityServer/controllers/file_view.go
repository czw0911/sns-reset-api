//查看文件
package controllers

import (
	//"XYAPIServer/XYGroupsServer/models"
	//"XYAPIServer/XYGroupsServer/libs"
	"XYAPIServer/XYLibs"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"io/ioutil"
)



type FileViewController struct {
	BaseController
}

func (u *FileViewController) Post() {
	
	resp := XYLibs.RespStateCode["method_not_find"]
	u.Data["json"] = resp
	u.ServeJson()

}


func (u *FileViewController) Get() {
	
	
	//resp := XYLibs.RespStateCode["nil"]
	//uidT,_ := u.GetInt64("UID")
	//uid := uint32(uidT)
	fileName := u.GetString("FileName")
	
	arr := strings.Split(fileName,".")
	
	if len(arr) != 2 {
		u.Data["json"] = XYLibs.RespStateCode["view_file_format_fail"]
		u.ServeJson()
		return
	}
	arrPath := strings.Split(arr[0],"_")
	if len(arrPath) != 6 {
		u.Data["json"] = XYLibs.RespStateCode["view_file_format_fail"]
		u.ServeJson()
		return
	}
	
	
	a , err := strconv.Atoi(arrPath[0])
	if err != nil {
		
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	b , err := strconv.ParseInt(arrPath[1],10,64)
	if err != nil {
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	c , err := strconv.Atoi(arrPath[2])
	if err != nil {
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	d , err := strconv.ParseInt(arrPath[3],10,64)
	if err != nil {
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	e , err := strconv.Atoi(arrPath[4])
	if err != nil {
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	f , err := strconv.ParseInt(arrPath[5],10,64)
	if err != nil {
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	fName := ""
	name := fmt.Sprintf("%d_%d_%d_%d_%d_%d",a,b,c,d,e,f)
	
	switch arr[1] {
		case "caf" :
			fName = fmt.Sprintf("%s.caf",name)
			
		case "png" :
			
			fName = fmt.Sprintf("%s.png",name)
		case "mp3" :
			
			fName = fmt.Sprintf("%s.mp3",name)
		case "aac" :	
			fName = fmt.Sprintf("%s.aac",name)
			
		default :
		u.Data["json"] = XYLibs.RespStateCode["view_file_name_fail"]
		u.ServeJson()
		return
	}
	
	path := fmt.Sprintf("%s/%d/%d/%d/%d/%d/%s",beego.AppConfig.String("upload_path"),a,b,c,d,e,fName)
	//println(path)
	data,err := ioutil.ReadFile(path)
	if err != nil {
		beego.Error("read file error:",err)
		u.Data["json"] = XYLibs.RespStateCode["view_file_not_find"]
		u.ServeJson()
		return
	}
	
	u.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	u.Ctx.Output.Header("Content-Type", "application/octet-stream")
	u.Ctx.Output.Body(data)
		
}

