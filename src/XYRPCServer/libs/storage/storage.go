//@Description 数据存储适配器
//@Contact czw@outlook.com
package storage

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Storage interface {
	//获取主库
	GetMaster() (orm.Ormer,error)
	//获取从库
	GetSalve()  (orm.Ormer,error)
	//获取hash表数量
	GetHashTableNum() int
	//存储初始化与运行
	Start() error
}

var adapters = make(map[string]Storage)

//注册存储
func Register(name string ,adapter Storage) {
	
	if adapter == nil {
		panic("Storage: adapter is null")
	}
	
	if _,ok := adapters[name]; ok {
		panic("Storage: name existing")
	}
	
	adapters[name] =  adapter
}

//获取存储
func NewStorage(adapterName string) (adapter Storage, err error) {
	
	adapter,ok := adapters[adapterName]
	
	if !ok {
		err = fmt.Errorf("Storage: unknown adapter name %q",adapterName)
	}
	
	err = adapter.Start()
	if err != nil {
		adapter = nil
	}
	return
}