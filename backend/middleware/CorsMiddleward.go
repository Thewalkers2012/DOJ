package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 写入允许访问的域名和方法
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		// 缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置可以访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 可以访问的 headers
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
