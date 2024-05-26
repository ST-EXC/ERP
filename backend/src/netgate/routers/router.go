package routers

import (
	"EDA.Project.ERP.backend.netgate/controller"
	"EDA.Project.ERP.backend.netgate/tools/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig((logger.LoggerToFile())))
	r.Use(logger.Recover)

	//store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/register", controller.RegisterController{}.Register)
	r.POST("/login", controller.LoginController{}.Login)

	return r
}
