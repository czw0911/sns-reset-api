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
	//"strconv"
	//"io"
	"os"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
)

const (
	DB_MASTER_CONN_PREFIX = "db_master_conn_prefix_"
	DB_SALVE_CONN_PREFIX =  "db_slave_conn_prefix_"
	DB_COMMON_TAG = "db_common_tag"
)

var (
	mysqlConnList = consistent.New()
	UserHashTableNum = 8
)


type mysqlConf struct {
	MaxIdleConn int `json:"max_idle_conn"`
	MaxOpenConn int `json:"max_open_conn"`
	CommmonDB	string	`json:"common_db_list"`
	UserHashTableNum	int	`json:"activity_table_hash_num"`
	UserDBList []mysqlDns `json:"activity_db_list"`
}

type mysqlDns struct {
	Dns string `json:"dns`
}

type mySQLHash struct {
	DBCountMaste int //主数据库数量
	DBCountSlave int //从数据库数量
}

func NewMySQLHash() *mySQLHash {
	return new(mySQLHash)
}

//一致性hash

func (c *mySQLHash) GetConsistentByGroupID(gid int64) (string,error){
	id := fmt.Sprintf("%d",gid)
	connConf, err := mysqlConnList.Get(id)
	if err != nil {
		return  "",err
	}
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
	UserHashTableNum = confDB.UserHashTableNum
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev"{
		orm.Debug = true
	}
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	
	dbNUM := len(confDB.UserDBList)
	
	for i := 0 ; i < dbNUM ; i++ {
		name := fmt.Sprintf("%s%d",DB_MASTER_CONN_PREFIX,i)	
		orm.RegisterDataBase(name, "mysql", confDB.UserDBList[i].Dns)
		mysqlConnList.Add(name)
		orm.SetMaxIdleConns(name,confDB.MaxIdleConn)
		orm.SetMaxOpenConns(name,confDB.MaxOpenConn)
	}
	orm.RegisterDataBase(DB_COMMON_TAG,"mysql",confDB.CommmonDB)
	orm.RegisterDataBase("default", "mysql", confDB.UserDBList[0].Dns)
	
	
}


