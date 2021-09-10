package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thewalkers2012/DOJ/logger"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/repository/redis"
	"github.com/Thewalkers2012/DOJ/routes"
	"github.com/Thewalkers2012/DOJ/settings"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Config.LogConfig); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success ...")

	// 3. 初始化 mysql 连接
	if err := mysql.Init(settings.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}

	// 4. 初始化 redis 连接
	if err := redis.Init(settings.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}

	// 5. 注册路由
	r := routes.Setup()

	// 6. 启动服务 （优雅关机操作）
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个 goroutine 启动服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
