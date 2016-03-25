//@Description 一致性hash数据库连接
//@Contact czw@outlook.com
package mysql

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"stathat.com/c/consistent"
	"fmt"
)

func NewHash() *Hash {
	return &Hash{consistent.New()}
}

type Hash struct {
	DB *consistent.Consistent 
}

func (c *Hash) ConverHashID(hashid interface{}) string {
	
	id := ""
	switch hashid.(type) {
		case int :
		 	id = fmt.Sprintf("%d",hashid)
		case int32 :
		 	id = fmt.Sprintf("%d",hashid)
		case int64 :
			id = fmt.Sprintf("%d",hashid)
		case string :
			id = hashid.(string)
	}
	
	return id
}

func (c *Hash) GetDB(dbPrefix string,hashid interface{}) ( orm.Ormer,error){
	
	id := c.ConverHashID(hashid)
	index, err := c.DB.Get(id)
	if err != nil {
		return  nil,err
	}
	name := fmt.Sprintf("%s%s",dbPrefix,index)
	d := orm.NewOrm()
	d.Using(name)
	return d,nil
}

func (c *Hash) PutDB(myqlDns , dbPrefix string,maxIdleConn,maxOpenConn int,hashid interface{}) {

		id := c.ConverHashID(hashid)	
		name := fmt.Sprintf("%s%s",dbPrefix,id)	
		orm.RegisterDataBase(name, "mysql", myqlDns)
		c.DB.Add(id)
		orm.SetMaxIdleConns(name,maxIdleConn)
		orm.SetMaxOpenConns(name,maxOpenConn)
}




