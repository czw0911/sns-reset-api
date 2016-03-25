//@Description 数据库连接
//@Contact czw@outlook.com
package mysql

import (
	"fmt"
	"sync"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs/storage"
)

var once sync.Once

type Connect struct {
	//json连接配置文件名
	JsonConfig  string
	//主库orm Reg 前缀
	MasterPrefix string
	//从库orm reg 前缀
	SlavePrefix string
	//主库
	MasterDB *Hash
	//从库
	SlaveDB *Hash
	//哈希id
	HashID interface{}
	//哈希表数量
	HashTableNum int
}

func (c *Connect) SetHashVal(val interface{}) {
	c.HashID = val
}

func (c *Connect) GetMaster() (orm.Ormer,error) {
	return c.MasterDB.GetDB(c.MasterPrefix,c.HashID)
}

func (c *Connect) GetSalve() (orm.Ormer,error) {
	return c.SlaveDB.GetDB(c.SlavePrefix,c.HashID)
}

func (c *Connect) GetHashTableNum() int {
	return c.HashTableNum
}

func (c *Connect) Start() error {
	
	conf,err := storage.ConfigFileRead(c.JsonConfig)
	if err != nil {
		return err
	}
	var confDB storage.MysqlConfig
	err = json.Unmarshal(conf,&confDB)
	if err != nil {
		return fmt.Errorf("Unmarshal %s config error :%q",c.JsonConfig,err.Error())
	}
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	
	c.HashTableNum = confDB.HashTableNum
	
	if len(confDB.MasterDBList) == 0 {
		return fmt.Errorf("not config master db")
	}
	
	for i,v := range confDB.MasterDBList {
		c.MasterDB.PutDB(v.Dns,c.MasterPrefix,confDB.MaxIdleConn,confDB.MaxOpenConn,i)
	}
	
	once.Do(func(){
		orm.RegisterDataBase("default", "mysql", confDB.MasterDBList[0].Dns)
	})
	
	
	for i,v := range confDB.SalveDBList {		
		c.SlaveDB.PutDB(v.Dns,c.SlavePrefix,confDB.MaxIdleConn,confDB.MaxOpenConn,i)	
	}
	orm.Debug = true
	return nil
}

