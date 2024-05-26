package main

import (
	"fmt"
	"net"

	"EDA.Project.ERP.backend.register/config"
	Register "EDA.Project.ERP.backend.register/proto/register"
	"EDA.Project.ERP.backend.register/service"
	"google.golang.org/grpc"
)

func main() {
	//监听地址为127.0.0.1:8972，采用tcp连接的方式，该函数会返回一个8792端口的监听器
	l, err := net.Listen(config.RPCMethod, config.RPCAddress)

	if err != nil {
		fmt.Printf("server net err:%v\n", err)
		return
	}

	//创建RPC服务端
	s := grpc.NewServer()

	//注册服务RegisterService到RPC服务端
	Register.RegisterRegisterServer(s, &service.RegisterService{})

	//启动RPC服务，并
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("server serve err:%v\n", err)
		return
	}
}
