// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package web

import (
	"context"
	"errors"
	"fish/internal/bootstrap"
	"fish/internal/config"
	"fish/internal/g"
	"fish/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Web struct {
	Gin *gin.Engine
}

func NewWeb() *Web {
	return &Web{
		Gin: gin.Default(),
	}
}

// Run 运行应用
func (a *Web) Run(addr string) {
	a.init()
	srv := http.Server{
		Addr:    addr,
		Handler: a.Gin,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			g.Logger.Infof("app %s started", config.App.ID)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	g.Logger.Infoln("Shutdown Server ...")
	bootstrap.RunDestructor()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		g.Logger.Fatal("Server Shutdown:", err)
	}
	g.Logger.Info("Server exiting")
}

// init 应用初始化
// 注册中间件， 注册路由
func (a *Web) init() {
	bootstrap.RunInit()
	a.useMiddleware()
	a.SetRoute()
}

func (a *Web) useMiddleware() {
	a.Gin.Use(middleware.Response)
}

func (a *Web) Get(path string, h middleware.Handler) {
	a.Gin.GET(path, middleware.Handle(h))
}

func (a *Web) Post(path string, h middleware.Handler) {
	a.Gin.POST(path, middleware.Handle(h))
}

func (a *Web) Put(path string, h middleware.Handler) {
	a.Gin.PUT(path, middleware.Handle(h))
}

func (a *Web) Delete(path string, h middleware.Handler) {
	a.Gin.DELETE(path, middleware.Handle(h))
}

func (a *Web) Options(path string, h middleware.Handler) {
	a.Gin.OPTIONS(path, middleware.Handle(h))
}
