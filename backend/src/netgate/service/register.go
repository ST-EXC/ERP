package service

import (
	"context"
	"fmt"

	"EDA.Project.ERP.backend.netgate/config"
	"EDA.Project.ERP.backend.netgate/proto/Register"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	RegisterClient Register.RegisterClient
	RegisterCtx    context.Context
)

func init() {
	conn, err := grpc.NewClient(config.RegisterAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("grpc.Dial failed,err:%v\n", err)
		return
	}
	RegisterCtx = context.Background()
	//建立连接
	RegisterClient = Register.NewRegisterClient(conn)
}
