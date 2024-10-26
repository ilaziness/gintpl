package main

import (
	"gintpl/internal/app/web"
	"gintpl/internal/app/web/route"
	"gintpl/pkg/log"
	"gintpl/pkg/storage/mysql"
	"gintpl/pkg/storage/redis"
	"gintpl/pkg/webapp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdWeb = &cobra.Command{
	Use:   "web",
	Short: "web server",
	Long:  `web server`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var cfgFile string
var port uint16

func init() {
	CmdWeb.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	CmdWeb.PersistentFlags().Uint16VarP(&port, "port", "p", 8080, "port")
}

func run() {
	app := webapp.New(string(web.Config.App.Mode))
	if web.Config.App.Port == 0 {
		web.Config.App.Port = port
	}
	loadComponent(app.Gin)
	app.Run(web.Config.App)
}

func loadComponent(engine *gin.Engine) {
	route.InitRoute(engine)
	log.Init()
	mysql.Init(web.Config.Db)
	redis.Init(web.Config.Redis)
}
