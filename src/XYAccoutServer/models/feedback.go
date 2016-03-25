//@Description 反馈
//@Contact czw@outlook.com

package models



type Feedback struct {

	UID uint32
	
	Contents string //反馈内容
	 
	Contact  string //联系人
	
	PostTime int64  //反馈时间
		
}


func (u * Feedback) Add() (bool,error) {
		db := ConnCommonDB()
		_, err := db.Raw("INSERT INTO feedback(Contents, Contact, PostTime) VALUES(?,?,?)",u.Contents,u.Contact,u.PostTime).Exec()
		if err != nil {
			return false,err
		}
		return true,nil
}

