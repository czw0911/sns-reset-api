//@Description mysql服务器地址hash
//@Contact czw@outlook.com

package libs

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"stathat.com/c/consistent"
	//"hash/fnv"
	"fmt"
	"strconv"
	//"io"
	"os"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
)

var (
	MasteDB *MySQLHash
	SalveDB *MySQLHash
)

const (
	DB_MASTER_CONN_PREFIX = "db_master_conn_prefix_"
	DB_SALVE_CONN_PREFIX =  "db_slave_conn_prefix_"
	DB_COMMON_TAG = "db_common_tag"
)



type mysqlConf struct {
	MaxIdleConn int `json:"max_idle_conn"`
	MaxOpenConn int `json:"max_open_conn"`
	CommmonDB	string	`json:"common_db_list"`
	MasterDBList []mysqlDns `json:"db_list_master"`
	SalveDBList []mysqlDns `json:"db_list_salve"`
}

type mysqlDns struct {
	Dns string `json:"dns`
}

type MySQLHash struct {
	DBConsistentMaste *consistent.Consistent //主数据库
	DBConsistentSlave *consistent.Consistent //从数据库
}

func NewMySQLHash() *MySQLHash {
	m := new(MySQLHash)
	m.DBConsistentMaste = consistent.New()
	m.DBConsistentSlave = consistent.New()
	return m
}

//一致性hash

func (c *MySQLHash) GetConsistentMasterDB(hashid int) (string,error){
	id := fmt.Sprintf("%d",hashid)
	index, err := c.DBConsistentMaste.Get(id)
	if err != nil {
		return  "",err
	}
	connConf := fmt.Sprintf("%s%s",DB_MASTER_CONN_PREFIX,index)
	return connConf,nil
}

func (c *MySQLHash) GetConsistentSalveDB(hashid int) (string,error){
	id := fmt.Sprintf("%d",hashid)
	index, err := c.DBConsistentSlave.Get(id)
	if err != nil {
		return  "",err
	}
	connConf := fmt.Sprintf("%s%s",DB_SALVE_CONN_PREFIX,index)
	return connConf,nil
}


func init(){
	// config
	
	workPath,_ := os.Getwd()
	appPath,_ := filepath.Abs(workPath)
	
	confPath := filepath.Join(appPath,"conf","mysql.json")
	if _, err := os.Stat(confPath); err != nil {
		if os.IsNotExist(err) {
			panic("mysql config not read" + err.Error())
		}
	}
	
	conf,_ := ioutil.ReadFile(confPath)
	var confDB mysqlConf
	if err := json.Unmarshal(conf,&confDB); err != nil {
		panic("Unmarshal mysql.json config error."+err.Error())
	}
	
	MasteDB = NewMySQLHash()
	SalveDB = NewMySQLHash()
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev"{
		orm.Debug = true
	}
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	
	
	for i,v := range confDB.MasterDBList {
		name := fmt.Sprintf("%s%d",DB_MASTER_CONN_PREFIX,i)	
		orm.RegisterDataBase(name, "mysql", v.Dns)
		MasteDB.DBConsistentMaste.Add(strconv.Itoa(i))
		orm.SetMaxIdleConns(name,confDB.MaxIdleConn)
		orm.SetMaxOpenConns(name,confDB.MaxOpenConn)
	}
	
	for i,v := range confDB.SalveDBList {
		name := fmt.Sprintf("%s%d",DB_SALVE_CONN_PREFIX,i)	
		orm.RegisterDataBase(name, "mysql", v.Dns)
		SalveDB.DBConsistentSlave.Add(strconv.Itoa(i))
		orm.SetMaxIdleConns(name,confDB.MaxIdleConn)
		orm.SetMaxOpenConns(name,confDB.MaxOpenConn)
	}
	
	orm.RegisterDataBase(DB_COMMON_TAG,"mysql",confDB.CommmonDB)
	orm.RegisterDataBase("default", "mysql", confDB.MasterDBList[0].Dns)
	
	
}


