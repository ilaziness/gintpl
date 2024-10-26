package webapp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gintpl/pkg/config"
	"gintpl/pkg/log"

	"github.com/gin-gonic/gin"
)

type App struct {
	Gin *gin.Engine
}

func New(mode string) *App {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	return &App{
		Gin: gin.Default(),
	}
}

// Run 运行应用
func (a *App) Run(appCfg *config.App) {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", appCfg.Port),
		Handler: a.Gin,
	}
	go func() {
		log.Logger.Infof("app %s star on: %s", appCfg.ID, srv.Addr)
		err := srv.ListenAndServe()
		if err != nil {
			log.Logger.Fatal("Start Server error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logger.Infoln("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server Shutdown error:", err)
	}
	log.Logger.Info("Server exiting")
}
