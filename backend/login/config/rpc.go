package config

import (
	"google.golang.org/grpc"
)

func ConnectToRPC() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
