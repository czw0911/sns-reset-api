/*
*@Description 配置
*@Contact czw@outlook.com
*/

package libs

import (
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

func NewAppConfig(cfg string) (*AppConfig,error){
	app := new(AppConfig)
	err := app.Load(cfg)
	if err != nil {
		return nil,err
	}
	return app,nil
}

type AppConfig struct {
	//运行端口
	RunPort int `json:"port"`
	//访问域名
	RunDomain string `json:"domain"`
	QRRedisAddr string `json:"qr_redis"`
	MySqlAdmin string `json:"mysql_admin"`
	LoginCallbackUrl string `json:"login_callback_url"`
}

func (self *AppConfig) Load (cfg string) error {
	
	data,err := ConfigFileRead(cfg)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data,self)
	if err != nil {
		return err
	}
	return nil
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