package service

import (
	"context"
	"fmt"
	"time"

	"code.gubai.com/backend/hello_client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	C      proto.GreeterClient
	Ctx    context.Context
	Cancel context.CancelFunc
)

func init() {
	//建立tcp连接
	conn, err := grpc.NewClient("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("grpc.Dial failed,err:%v\n", err)
		return
	}

	// 创建Greeter服务的客户端，也就是 c
	// 在Greeter服务中，定义了一个SayHello方法，通过建立该服务的客户端，可以使用其自动生成的代码
	C = proto.NewGreeterClient(conn)

	//获取context以传递信息
	Ctx, Cancel = context.WithTimeout(context.Background(), 100*time.Second)

}
