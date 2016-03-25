// mysql存储
//@Contact czw@outlook.com

package models


import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"XYAPIServer/XYAccoutServer/libs"

)
const (
	DB_MASTER_CONN_PREFIX = "db_master_conn_prefix_"
	DB_SALVE_CONN_PREFIX =  "db_slave_conn_prefix_"
	DB_COMMON_TAG = "db_common_tag"
)

var (
	USER_HASH_TABLE_NUM = 8
	
	TABLE_LIMIT_NUM = 10
)


type Base struct {}

//func ConnSlaveDB(uid string) orm.Ormer {
//	d := orm.NewOrm()
//	index := fmt.Sprintf("%s%d",DB_SALVE_CONN_PREFIX,dbHashByUID.GetMySQLHashByUID(0,uid))
//	d.Using(index)
//	return d
//}

func ConnMasterDB(uid uint32) orm.Ormer {
	d := orm.NewOrm()
	dbHashByUID  := libs.NewMySQLHash()
	index,err := dbHashByUID.GetConsistentByUID(uid)
	if err != nil {
		beego.Error(err)
	}
	d.Using(index)
	return d
}

func ConnCommonDB() orm.Ormer {
	d := orm.NewOrm()
	d.Using(DB_COMMON_TAG)
	return d
}

func init() {
	USER_HASH_TABLE_NUM = libs.UserHashTableNum
	// config
	
//	workPath,_ := os.Getwd()
//	appPath,_ := filepath.Abs(workPath)
	
//	confPath := filepath.Join(appPath,"conf","mysql.json")
//	if _, err := os.Stat(confPath); err != nil {
//		if os.IsNotExist(err) {
//			panic("mysql config not read",err)
//		}
//	}
	
//	conf,_ := ioutil.ReadFile(confPath)
//	var confDB mysqlConf
//	if err := json.Unmarshal(conf,&confDB); err != nil {
//		panic("Unmarshal mysql.json config error."+err.Error())
//	}
	
	
//	orm.Debug = true
//	orm.RegisterDriver("mysql", orm.DR_MySQL)
//	//
//	//主库连接
//	//

//	//dbHashByUID.DBCountMaste = len(confDB.UserDBList)
//	dbNUM := len(confDB.UserDBList)
//	orm.SetMaxIdleConns(confDB.MaxIdleConn)
//	orm.SetMaxOpenConns(confDB.MaxOpenConn)
//	for i := 0 ; i < dbNUM ; i++ {
//		name := fmt.Sprintf("%s%d",DB_MASTER_CONN_PREFIX,i)
		
//		orm.RegisterDataBase(name, "mysql", arrDB[i])
		
//	}
//	orm.RegisterDataBase(DB_COMMON_TAG,"mysql",confDB.CommmonDB)
//	orm.RegisterDataBase("default", "mysql", arrDB[0])
	
	//
	//从库连接
	//
//	slaveConnDB := beego.AppConfig.String("mysql_slave_db")
//	if slaveConnDB == "" {
//		panic("slave db config not read")
//	}
	
//	arrDBS := strings.Split(slaveConnDB,",")
//	dbHashByUID.DBCountSlave = len(arrDBS)
//	for i := 0 ; i < dbHashByUID.DBCountSlave ; i++ {
//		name := fmt.Sprintf("%s%d",DB_SALVE_CONN_PREFIX,i)
//		orm.RegisterDataBase(name, "mysql", arrDBS[i])
//	}
	
	//SlectDB.SlaveDB.Raw("SET sql_mode='NO_UNSIGNED_SUBTRACTION'").Exec()
}