package routes

import (
	"github.com/Thewalkers2012/DOJ/api"
	"github.com/Thewalkers2012/DOJ/logger"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 设置路由组
	v1 := r.Group("/api/v1")

	// 注册用户的业务
	v1.POST("/signup", api.SignUpHandler)

	return r
}
