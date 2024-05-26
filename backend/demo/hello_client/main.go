package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"code.gubai.com/backend/hello_client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//建立tcp连接
	conn, err := grpc.NewClient("127.0.0.1:8992", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("grpc.Dial failed,err:%v\n", err)
		return
	}

	// 创建Greeter服务的客户端，也就是 c
	// 在Greeter服务中，定义了一个SayHello方法，通过建立该服务的客户端，可以使用其自动生成的代码
	c := proto.NewGreeterClient(conn)

	//获取context以传递信息,并且限制超时时间为100s
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	resp, err := c.SayHello(ctx, &proto.HelloRequest{Name: "golang"})
	if err != nil {
		fmt.Printf("c.SayHello failed ,err:%v\n", err)
		return
	}
	//拿到了RPC响应，获取对应的字段值应当使用Get`字段名`()方法
	log.Printf("resp is :%v", resp.GetReply())
}
