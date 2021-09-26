package routes

import (
	"net/http"

	"github.com/Thewalkers2012/DOJ/api"
	"github.com/Thewalkers2012/DOJ/logger"
	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.CorsMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 设置路由组
	v1 := r.Group("/api/v1")

	// 注册用户的业务
	v1.POST("/register", api.SignUpHandler)
	v1.POST("/login", api.LoginHandler)
	v1.GET("/info", middleware.JWTAuthorMiddleware(), api.InfoHandler)

	// 题目相关的业务
	v1.POST("/problem", middleware.JWTAuthorMiddleware(), api.CreateProblemHandler)
	v1.GET("/problem/:id", middleware.JWTAuthorMiddleware(), api.GetProblemByIDHandler)
	v1.GET("/problem", middleware.JWTAuthorMiddleware(), api.GetProblemList)
	v1.POST("/submit", middleware.JWTAuthorMiddleware(), api.RunCodeHandler)

	// 文章相关
	v1.POST("/category", middleware.JWTAuthorMiddleware(), api.CreateCategory)

	return r
}
