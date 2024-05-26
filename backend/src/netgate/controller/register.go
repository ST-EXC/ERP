package controller

import (
	"context"

	"EDA.Project.ERP.backend.netgate/proto/Register"
	"EDA.Project.ERP.backend.netgate/service"
	"github.com/gin-gonic/gin"
)

type RegisterController struct{}

func (u RegisterController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		Error(c, 500, "register infomation error")
		return
	}

	resp, _ := service.RegisterClient.IsUserExist(service.RegisterCtx, &Register.IsUserExistRequest{Username: username})
	exist := resp.GetExist()

	if exist {
		Error(c, 501, "user existed")
		return
	}

	_, err := service.RegisterClient.Register(context.Background(), &Register.RegisterRequest{Username: username, Password: password})

	if err != nil {
		Error(c, 503, "register error")
		return
	}

	Success(c, 200, "user add success", username)
}
