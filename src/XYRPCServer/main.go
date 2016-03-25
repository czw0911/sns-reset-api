//@Description乡音rpc服务器
//@Contact czw@outlook.com

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"encoding/json"
	"XYAPIServer/XYRPCServer/models"
	"XYAPIServer/XYRPCServer/libs/storage"
)

var runPort int

func main() {
	
	
	rpc.Register(models.NewLogsReg())
	rpc.Register(models.NewLogsLogin())
	
	l , err := net.Listen("tcp",fmt.Sprintf(":%d",runPort))
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("rpc server runing...")
	
	for {
		conn,err := l.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func init() {
	cfg,err := storage.ConfigFileRead("app.json")
	if err != nil {
		panic(err)
	}
	var appCfg storage.AppConfig
	err = json.Unmarshal(cfg,&appCfg)
	if err != nil {
		panic(err)
	}
	if appCfg.RunPort < 1024 {
		panic("run port config error")
	}
	runPort =  appCfg.RunPort
}