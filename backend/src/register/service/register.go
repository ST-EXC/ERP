package service

import (
	"context"
	"fmt"

	"EDA.Project.ERP.backend.register/models"
	Register "EDA.Project.ERP.backend.register/proto/register"
	"EDA.Project.ERP.backend.register/tools/encode"
)

type RegisterService struct {
	Register.UnimplementedRegisterServer
}

// 实现Register服务的方法
func (r *RegisterService) IsExist(ctx context.Context, in *Register.IsUserExistRequest) (*Register.IsUserExistResponse, error) {
	username := in.GetUsername()
	fmt.Printf("username=%v\n", username)
	if username == "" {
		return &Register.IsUserExistResponse{Exist: false}, nil
	}

	_, err := models.User{}.GetUserInfoByUsername(username)
	if err != nil {
		return &Register.IsUserExistResponse{Exist: false}, err
	}

	return &Register.IsUserExistResponse{Exist: true}, nil
}

func (r *RegisterService) Register(ctx context.Context, in *Register.RegisterRequest) (*Register.RegisterResponse, error) {

	username := in.GetUsername()
	password := in.GetPassword()

	if username == "" || password == "" {
		return &Register.RegisterResponse{Id: -1}, nil
	}

	user, _ := models.User{}.GetUserInfoByUsername(username)
	if user.Id != 0 {
		return &Register.RegisterResponse{Id: int32(user.Id)}, nil
	}

	UserId, err := models.User{}.AddUser(username, encode.EncryMd5(password))
	if err != nil {
		return &Register.RegisterResponse{Id: -1}, err
	}

	return &Register.RegisterResponse{Id: int32(UserId)}, nil
}
