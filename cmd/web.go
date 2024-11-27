package main

import (
	"gintpl/internal/app/web"
	"gintpl/internal/app/web/route"
	_ "gintpl/internal/queue"
	_ "gintpl/internal/timer"
	"gintpl/pkg/log"
	"gintpl/pkg/middleware"
	"gintpl/pkg/server"
	"gintpl/pkg/storage/cache"
	"gintpl/pkg/storage/mysql"
	"gintpl/pkg/storage/redis"
	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

var CmdWeb = &cobra.Command{
	Use:   "httpd",
	Short: "http web server",
	Long:  `start a http web server listening`,
	Run: func(_ *cobra.Command, _ []string) {
		run()
	},
}

func run() {
	// config.InitNacos(web.Config.Nacos, web.Config)
	log.SetLevel(web.Config.App.Mode)
	initComponent()

	webServer := server.NewWeb(web.Config.App)
	// 设置自定义中间件
	useMiddleware(webServer.Gin)
	// 初始化路由
	route.InitRoute(webServer.Gin)
	// 运行
	webServer.Run()
}

// initComponent 初始化需要用到的组件
func initComponent() {
	mysql.Init(web.Config.Db)
	redis.Init(web.Config.Redis)
	cache.InitRedisCache(redis.Client)
	// rocketmq.InitProducer(web.Config.RocketMq)
	// rocketmq.InitConsumer(web.Config.RocketMq)
	// otel.InitTracer(web.Config.WebApp.ID, web.Config.Otel)
}

func useMiddleware(g *gin.Engine) {
	g.Use(middleware.Test())
}
