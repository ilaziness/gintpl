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
	// 初始化nacos配置中心
	// config.InitNacos(web.Config.Nacos, web.Config)

	// 设置运行模式
	log.SetLevel(web.Config.App.Mode)

	// 载入所需组件
	initComponent()

	webServer := server.NewWeb(web.Config.App)
	// 设置自定义中间件
	useMiddleware(webServer.Gin)
	// 初始化路由
	route.InitRoute(webServer.Gin)
	// nancos注册服务
	// server.Register(web.Config.Nacos, web.Config.App)

	// 运行服务
	webServer.Run()
}

// initComponent 初始化需要用到的组件
func initComponent() {
	// gormCmd
	// mysql.InitGORM(web.Config.Db)

	// ent
	dao.SetClient(ent.NewClient(ent.Driver(mysql.EntDriver(web.Config.Db))))
	dao.AutoMigration()

	// sqlx
	//mysql.InitSqlx(web.Config.Db)

	// 初始化redis
	//redis.Init(web.Config.Redis)

	// 初始化redis 缓存功能
	//cache.InitRedisCache(redis.Client)

	// 启用rocketmq 生产者
	// rocketmq.InitProducer(web.Config.RocketMq)

	// 启用rocketmq 消费者
	// rocketmq.InitConsumer(web.Config.RocketMq)

	// 开启链路追踪
	// otel.InitTracer(web.Config.WebApp.ID, web.Config.Otel)
}

func useMiddleware(g *gin.Engine) {
	g.Use(middleware.Test())
}
