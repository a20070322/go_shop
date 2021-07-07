package initialize

import (
	"context"
	"github.com/a20070322/shop-go/controller/auth"
	"github.com/a20070322/shop-go/controller/basic"
	"github.com/a20070322/shop-go/controller/customer"
	"github.com/a20070322/shop-go/controller/goods"
	"github.com/a20070322/shop-go/controller/order"
	"github.com/a20070322/shop-go/controller/pay"
	"github.com/a20070322/shop-go/controller/test"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/middleware"
	"github.com/a20070322/shop-go/types/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Router() {
	router := gin.New()
	// 本地开发引入自带中间件比较美观
	if global.AppSetting.Env == config.EnvLOCAL {
		//router.Use(gin.Logger(), gin.Recovery())
	}
	if global.AppSetting.Env != config.EnvLOCAL {
		//router.Use(middleware.GinRecovery(global.Logger, true))
	}
	// 日志记录至文件
	router.Use(middleware.GinLogger())
	router.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})
	api := router.Group("api")
	// api注入
	apiRegisterApis(api)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(global.AppSetting.Server.Port),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("数据库连接关闭")
	global.Db.Close()
	global.RabbitMq.Close()
	log.Println("Server exiting")
}

func apiRegisterApis(api *gin.RouterGroup) {
	auth.Router(api)
	test.TestRouter(api)
	goods.Router(api)
	order.Router(api)
	customer.Router(api)
	basic.Router(api)
	pay.Router(api)
}
