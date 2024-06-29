package main

import (
	"log"
	"login/config"
	"login/data"
	"login/proto/login"
	"login/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 初始化数据库连接
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	data.InitDB(db)

	// 设置gRPC服务器
	grpcServer := grpc.NewServer()

	// 注册服务
	login.RegisterLoginServer(grpcServer, &service.LoginService{})

	// 启动监听
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Login Server is running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
