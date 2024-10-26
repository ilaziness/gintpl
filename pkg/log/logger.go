// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

// Package log provide global singleton object access
package log

import (
	"log"

	"go.uber.org/zap"
)

var (
	Logger    *zap.SugaredLogger
	zapLogger *zap.Logger
)

func Init() {
	SetLogger()
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
	log.Println("zap logger sync")
	err := zapLogger.Sync()
	if err != nil {
		log.Println(err)
	}
}
