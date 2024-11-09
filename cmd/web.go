package main

import (
	"gintpl/internal/app/web"
	"gintpl/internal/app/web/route"
	_ "gintpl/internal/timer"
	"gintpl/pkg/log"
	"gintpl/pkg/storage/cache"
	"gintpl/pkg/storage/mysql"
	"gintpl/pkg/storage/redis"
	"gintpl/pkg/websrv"

	"github.com/spf13/cobra"
)

var CmdWeb = &cobra.Command{
	Use:   "web",
	Short: "web server",
	Long:  `start web server listening`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	app := websrv.New(web.Config.App)
	loadComponent()

	// 初始化路由
	route.InitRoute(app.Gin)
	// 运行
	app.Run()
}

// loadComponent 初始化需要用到的组件
func loadComponent() {
	log.Init() // 必须
	mysql.Init(web.Config.Db)
	redis.Init(web.Config.Redis)
	cache.InitRedisCache(redis.Client)
}
