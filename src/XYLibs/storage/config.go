//@Description 存储配置
//@Contact czw@outlook.com
package storage
import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

type MysqlDns struct {
	Dns string `json:"dns"`
}


type MysqlConfig struct {
	//连接池最大空闲数
	MaxIdleConn int `json:"max_idle_conn"`
	//连接池最大连接数据
	MaxOpenConn int `json:"max_open_conn"`
	//hash表数量
	HashTableNum int `json:"table_hash_num"`
	//主库连接配置
	MasterDBList []MysqlDns `json:"db_list_master"`
	//从库连接配置
	SalveDBList  []MysqlDns `json:"db_list_salve"`
}

type AppConfig struct {
	
	RunPort int `json:"port"`
}


func ConfigFileRead(jsonConfig string) ([]byte,error) {
	workPath,_ := os.Getwd()
	appPath,_ := filepath.Abs(workPath)
	
	confPath := filepath.Join(appPath,"conf",jsonConfig)
	if _, err := os.Stat(confPath); err != nil {
		if os.IsNotExist(err) {
			return nil,fmt.Errorf("%s config file not find,%q" ,jsonConfig, err.Error())
		}
	}
	return ioutil.ReadFile(confPath)
}