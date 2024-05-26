package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"net"

	"google.golang.org/grpc"
)

// 定义对应于服务Greeter的结构体Greeter，ps：该结构体名字可以不是Greeter
// 注意：该结构体必须实现UnimplementedGreeterServer结构体，也就是Greeter结构体未实现所需实现的方法时的后备方法
type Greeter struct {
	pb.UnimplementedGreeterServer
}

// 实现Greeter服务的方法SayHello
func (s *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	//监听地址为127.0.0.1:8972，采用tcp连接的方式，该函数会返回一个8793端口的监听器
	l, err := net.Listen("tcp", "127.0.0.1:8992")

	if err != nil {
		fmt.Printf("server net err:%v\n", err)
		return
	}

	//创建RPC服务端
	s := grpc.NewServer()

	//注册服务Greeter到RPC服务端
	pb.RegisterGreeterServer(s, &Greeter{})

	//启动RPC服务，并
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("server serve err:%v\n", err)
		return
	}
}
