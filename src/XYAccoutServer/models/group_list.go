//@Description 群组类型列表
//@Contact czw@outlook.com


package models

import (
	//"github.com/astaxie/beego/orm"
	"fmt"
)

type GroupList struct {
	GroupId string
	GroupName string
}




func (r *GroupList) Reg()(bool,error){
		db := ConnCommonDB()
		sqlStr := fmt.Sprintf("REPLACE INTO %s(GroupID,GroupName) VALUES(?,?)","group_list")
		_,err := db.Raw(sqlStr,r.GroupId,r.GroupName).Exec()
		if err != nil {
			return false,err
		}
		return true,nil

}


