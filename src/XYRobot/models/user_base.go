//@Description 用户基本信息
//@Contact czw@outlook.com

package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"stathat.com/c/consistent"

)

var (
	userBaseTableNameHash = consistent.New()
)

type UserBase struct {
	tableName string //hash到的表名

	dbConn orm.Ormer //hash到的数据库

	UID uint32 //用户id

	Account string //用户账号

	PassWord string //登录密码

	RegType int8 // 注册类型,1:手机登陆 2：微博登陆 3:微信登陆

	IsHometown int8 // 是否已建立家乡档案 0:没有，1：有

	IsFollow int8 //是否关注0:未关注；1：已关注；2：互相关注

	IsMember int8 // 是否会员

	AuthSendNum int // 认证他人乡音次数

	AuthRecvNum int // 获得乡音认证次数

	RemainDays int // vip剩余天数

	GrowNum int //成长值

	RPNum int // 人品值

	MedalNum int // 勋章个数

	RegisterTime int64 //注册时间

	BindPhone uint64 //绑定手机号

	HomeProvinceID int //  家乡省id
}

func (u *UserBase) initDBAndTable() {
	u.dbConn = ConnMasterDBUser(u.UID)
	id := fmt.Sprintf("%d", u.UID)
	u.tableName, _ = userBaseTableNameHash.Get(id)
}

func (u *UserBase) IsUIDExist() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT ID FROM %s WHERE UID = ? LIMIT 1", u.tableName)
	var res []orm.Params
	_, err := u.dbConn.Raw(sqlStr, u.UID).Values(&res)
	if err != nil {
		return true, err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true, errors.New(u.Account + " 账号已存在")
}

func (u *UserBase) Reg() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("INSERT INTO %s(RegType,UID,Account,PassWord,RegisterTime,BindPhone,HomeProvinceID) VALUES(?,?,?,?,?,?,?)", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.RegType, u.UID, u.Account, u.PassWord, u.RegisterTime, u.BindPhone, u.HomeProvinceID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserBase) Login() (bool, []orm.Params, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT UID, Account,  IsMember, AuthSendNum, AuthRecvNum, RemainDays, GrowNum, RPNum, MedalNum ,BindPhone,HomeProvinceID FROM %s WHERE UID =? AND PassWord = ? LIMIT 1", u.tableName)
	var res []orm.Params
	_, err := u.dbConn.Raw(sqlStr, u.UID, u.PassWord).Values(&res)
	if err != nil {
		return false, nil, err
	}
	if len(res) == 1 {
		return true, res, nil
	}
	return false, nil, nil
}

func (u *UserBase) GeteUserBaseByUID() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT UID, Account, IsMember, AuthSendNum, AuthRecvNum, RemainDays, GrowNum, RPNum, MedalNum ,BindPhone,HomeProvinceID FROM %s WHERE UID =?  LIMIT 1", u.tableName)
	var res []orm.Params
	_, _ = u.dbConn.Raw(sqlStr, u.UID).Values(&res)
	return res
}

func (u *UserBase) GeteHomeProvinceIDByUID() []orm.Params {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("SELECT HomeProvinceID FROM %s WHERE UID =?  LIMIT 1", u.tableName)
	var res []orm.Params
	_, _ = u.dbConn.Raw(sqlStr, u.UID).Values(&res)
	return res
}

func (u *UserBase) SeteHomeProvinceIDByUID() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET HomeProvinceID = ?  WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.HomeProvinceID, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}


func (u *UserBase) Delete() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("DELETE FROM  %s  WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserBase) UpdatePassWD() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET PassWord = ? WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.PassWord, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserBase) UpdateBindPhone() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET BindPhone = ? WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.BindPhone, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserBase) UpdateAuthSendNum() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET AuthSendNum = AuthSendNum + ? WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.AuthSendNum, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil

}

func (u *UserBase) UpdateAuthRecvNum() (bool, error) {
	u.initDBAndTable()
	sqlStr := fmt.Sprintf("UPDATE  %s  SET AuthRecvNum = AuthRecvNum + ? WHERE UID = ? ", u.tableName)
	_, err := u.dbConn.Raw(sqlStr, u.AuthRecvNum, u.UID).Exec()
	if err != nil {
		return false, err
	}
	return true, nil

}


func init() {
	num := UserDB.GetHashTableNum()
	for i := 0; i < num; i++ {
		tn := fmt.Sprintf("user_base_%d", i)
		userBaseTableNameHash.Add(tn)
	}
}
