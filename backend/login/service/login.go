package service

import (
	"context"
	"login/proto/login"
)

type LoginService struct {
	login.UnimplementedLoginServer
}

func (s *LoginService) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	// 登录逻辑

	success := false

	return &login.LoginResponse{Success: success}, nil
}
