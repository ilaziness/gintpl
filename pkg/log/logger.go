// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

// Package log provide global singleton object access
package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger    *zap.SugaredLogger
	zapLogger *zap.Logger
	logLevel  = zapcore.DebugLevel
)

func Init(mode ...string) {
	if len(mode) > 0 {
		switch mode[0] {
		case "debug":
			logLevel = zapcore.DebugLevel
		case "release":
			logLevel = zapcore.InfoLevel
		}
	}
	setLogger()
}

func setLogger() {
	// 配置日志写入文件
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/app.log", // 日志文件名
		MaxSize:    10,            // 每个日志文件的最大大小（MB）
		MaxBackups: 10,            // 保留的旧日志文件的最大数量
		MaxAge:     60,            // 保留旧日志文件的最大天数
		Compress:   true,          // 是否压缩旧日志文件
	})
	// 配置日志写入控制台
	consoleWriter := zapcore.Lock(os.Stdout)
	// 配置日志级别
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel
	})
	// 创建一个核心（Core），将日志同时写入文件和控制台
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // 使用 JSON 编码器
			fileWriter,
			highPriority,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(buildEncoderConsoleConfig()), // 使用控制台编码器
			consoleWriter,
			highPriority,
		),
	)

	// 创建日志记录器
	zapLogger = zap.New(core, zap.AddCaller())
	Logger = zapLogger.Sugar()
	Logger.Infoln("zap logger created")
}

func FlushLogger() {
	if zapLogger == nil {
		return
	}
	err := zapLogger.Sync()
	if err != nil {
		log.Println(err)
	}
}

// buildEncoderConsole 自定义控制台编码器
func buildEncoderConsoleConfig() zapcore.EncoderConfig {
	consoleEncoderConfig := zap.NewProductionEncoderConfig()
	consoleEncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000")) // 时间格式化到毫秒
	}
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 添加颜色
	consoleEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder      // 简短的调用者信息
	return consoleEncoderConfig
}

// Debug 增加了记录trace id
func Debug(ctx context.Context, tpl string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().HasTraceID() {
		Logger.Debugw(fmt.Sprintf(tpl, args...), "trace_id", span.SpanContext().TraceID().String())
		return
	}
	Logger.Debugf(tpl, args...)
}

// Info 增加了记录trace id
func Info(ctx context.Context, tpl string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().HasTraceID() {
		Logger.Infow(fmt.Sprintf(tpl, args...), "trace_id", span.SpanContext().TraceID().String())
		return
	}
	Logger.Infof(tpl, args...)
}

// Warn 增加了记录trace id
func Warn(ctx context.Context, tpl string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().HasTraceID() {
		Logger.Warnw(fmt.Sprintf(tpl, args...), "trace_id", span.SpanContext().TraceID().String())
		return
	}
	Logger.Warnf(tpl, args...)
}

// Error 增加了记录trace id
func Error(ctx context.Context, tpl string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().HasTraceID() {
		Logger.Errorw(fmt.Sprintf(tpl, args...), "trace_id", span.SpanContext().TraceID().String())
		return
	}
	Logger.Errorf(tpl, args...)
}
