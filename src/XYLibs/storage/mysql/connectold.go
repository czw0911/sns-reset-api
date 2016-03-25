//@Description 兼容以前的数据库连接
//@Contact czw@outlook.com
package mysql

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"XYAPIServer/XYLibs/storage"
)


type ConnectOld struct {
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

func (c *ConnectOld) SetHashVal(val interface{}) {
	c.HashID = val
}

func (c *ConnectOld) GetMaster() (orm.Ormer,error) {
	return c.MasterDB.GetDBold(c.MasterPrefix,c.HashID)
}

func (c *ConnectOld) GetSalve() (orm.Ormer,error) {
	return c.SlaveDB.GetDBold(c.SlavePrefix,c.HashID)
}

func (c *ConnectOld) GetHashTableNum() int {
	return c.HashTableNum
}

func (c *ConnectOld) Start() error {
	
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
		c.MasterDB.PutDBold(v.Dns,c.MasterPrefix,"db_master_conn_prefix_",confDB.MaxIdleConn,confDB.MaxOpenConn,i)
	}
	
	once.Do(func(){
		orm.RegisterDataBase("default", "mysql", confDB.MasterDBList[0].Dns)
	})
	
	
	for i,v := range confDB.SalveDBList {		
		c.SlaveDB.PutDBold(v.Dns,c.SlavePrefix,"db_slave_conn_prefix_",confDB.MaxIdleConn,confDB.MaxOpenConn,i)	
	}
	orm.Debug = true
	return nil
}

