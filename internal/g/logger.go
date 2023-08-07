// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

// Package g provide global singleton object access
package g

import (
	"fish/internal/bootstrap"
	"go.uber.org/zap"
	"log"
)

var (
	Logger    *zap.SugaredLogger
	zapLogger *zap.Logger
)

func init() {
	bootstrap.RegisterStartup(SetLogger)
	bootstrap.RegisterDestructor(FlushLogger)
	log.Println("register logger")
}

func SetLogger() {
	var err error
	zapLogger, err = zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	Logger = zapLogger.Sugar()
	Logger.Infoln("zap logger created")
}

func FlushLogger() {
	err := zapLogger.Sync()
	if err != nil {
		log.Println("zap logger sync")
		log.Println(err)
	}
}
