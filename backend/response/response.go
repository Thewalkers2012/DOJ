package response

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg interface{}) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
