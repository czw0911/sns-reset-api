//@Description 标签列表
//@Contact czw@outlook.com


package models

import (
	"github.com/astaxie/beego/orm"
)



type TagsList struct {

	TagID int 
	
	TagName string
	
	TagColor string
}


func (r *TagsList) GetAll() []orm.Params {
	
	var b []orm.Params
	db := ConnCommonDB()
	_, _ = db.Raw("SELECT TagID, TagName, TagColor FROM tags_list").Values(&b)
	return b
}


