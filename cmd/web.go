package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilaziness/gintpl/internal/app/web"
	"github.com/ilaziness/gintpl/internal/app/web/route"
	"github.com/ilaziness/gintpl/internal/dao"
	"github.com/ilaziness/gintpl/internal/ent"
	_ "github.com/ilaziness/gintpl/internal/queue"
	_ "github.com/ilaziness/gintpl/internal/timer"
	"github.com/ilaziness/gokit/log"
	"github.com/ilaziness/gokit/middleware"
	"github.com/ilaziness/gokit/server"
	"github.com/ilaziness/gokit/storage/mysql"
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
	// 注册服务
	// server.Register(web.Config.Nacos, web.Config.App)
	// 运行
	webServer.Run()
}

// initComponent 初始化需要用到的组件
func initComponent() {
	// gormCmd
	// mysql.InitGORM(web.Config.Db)
	// ent
	dao.SetClient(ent.NewClient(ent.Driver(mysql.EntDriver(web.Config.Db))))

	//redis.Init(web.Config.Redis)
	//cache.InitRedisCache(redis.Client)
	// rocketmq.InitProducer(web.Config.RocketMq)
	// rocketmq.InitConsumer(web.Config.RocketMq)
	// otel.InitTracer(web.Config.WebApp.ID, web.Config.Otel)
}

func useMiddleware(g *gin.Engine) {
	g.Use(middleware.Test())
}
