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
	v1.GET("/user", middleware.JWTAuthorMiddleware(), api.GetUserList)
	v1.PUT("/user", middleware.JWTAuthorMiddleware(), api.UpdateUser)
	v1.GET("/user_detail", middleware.JWTAuthorMiddleware(), api.GetUserDetails)
	v1.DELETE("/user", middleware.JWTAuthorMiddleware(), api.DeleteUserHandler)

	// 题目相关的业务
	v1.POST("/problem", middleware.JWTAuthorMiddleware(), api.CreateProblemHandler)
	v1.GET("/problem/:id", middleware.JWTAuthorMiddleware(), api.GetProblemByIDHandler)
	v1.GET("/problem", middleware.JWTAuthorMiddleware(), api.GetProblemList)
	v1.DELETE("/problem", middleware.JWTAuthorMiddleware(), api.DeleteProblem)

	// 文章相关
	v1.POST("/category", middleware.JWTAuthorMiddleware(), api.CreateCategory)
	v1.GET("/category", middleware.JWTAuthorMiddleware(), api.GetCategoryByProblem)
	v1.GET("/category_list", middleware.JWTAuthorMiddleware(), api.GetAllCategories)
	v1.GET("/category_details", middleware.JWTAuthorMiddleware(), api.GetCategoryDetails)
	v1.DELETE("/category", middleware.JWTAuthorMiddleware(), api.DeleteCategory)

	// 提交题目相关路由
	v1.POST("/submit", middleware.JWTAuthorMiddleware(), api.RunCodeHandler)
	v1.GET("/submission", middleware.JWTAuthorMiddleware(), api.GetAllSubmissionsByIdAndProblem)
	v1.GET("/solved", middleware.JWTAuthorMiddleware(), api.GetPersonSolved) // 获取个人 AC 题目数量
	v1.GET("/submit_count", middleware.JWTAuthorMiddleware(), api.GetPersonSubmission)

	// 比赛相关的路由
	v1.POST("/context", middleware.JWTAuthorMiddleware(), api.CreateContextHandler)
	v1.GET("/context_list", middleware.JWTAuthorMiddleware(), api.GetContextList)
	v1.GET("/context/:id", middleware.JWTAuthorMiddleware(), api.GetContext)
	v1.DELETE("/context", middleware.JWTAuthorMiddleware(), api.DeleteContext)
	v1.PUT("/context", middleware.JWTAuthorMiddleware(), api.UpdateContext)

	// 比赛题目相关
	v1.POST("/context_problem", middleware.JWTAuthorMiddleware(), api.AddProblemToContext)
	v1.GET("/context_problem", middleware.JWTAuthorMiddleware(), api.ProblemInContext) // 判断该题目是否在 context 中
	v1.DELETE("/context_problem", middleware.JWTAuthorMiddleware(), api.DeletePorblemInContext)
	v1.GET("/context_problem_list", middleware.JWTAuthorMiddleware(), api.ContextProblemList)
	v1.GET("/context_problems", middleware.JWTAuthorMiddleware(), api.GetAllContextProblem)

	return r
}
