package routes

import (
	"net/http"

	"github.com/Thewalkers2012/DOJ/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	return r
}
