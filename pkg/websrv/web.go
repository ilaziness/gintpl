// Package websrv provide web engine
package websrv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gintpl/pkg/config"
	"gintpl/pkg/log"
	"gintpl/pkg/middleware"
	"gintpl/pkg/timer"

	"github.com/gin-gonic/gin"
)

type App struct {
	Gin *gin.Engine
}

// New 创建一个web app
func New() *App {
	return &App{
		Gin: NewGin(),
	}
}

// NewGin gin engine
func NewGin() *gin.Engine {
	e := gin.New()
	e.ContextWithFallback = true
	e.Use(middleware.LogReq(), gin.CustomRecoveryWithWriter(nil, middleware.RecoveryHandle))
	return e
}

// Run 运行应用
func (a *App) Run(appCfg *config.App) {
	if appCfg.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	starup()

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", appCfg.Port),
		Handler: a.Gin,
	}
	go func() {
		log.Logger.Infof("app %s star on: %s", appCfg.ID, srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
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
	stopup()
	log.Logger.Info("Server exiting")
}

func starup() {
	timer.Run()
}

func stopup() {
	timer.Stop()
	log.FlushLogger()
}
