package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 写入允许访问的域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		// 缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置可以访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 设置可以访问的 Header
		ctx.Writer.Header().Set("Access-Control-Allow-Header", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
