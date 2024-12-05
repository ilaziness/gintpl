// Package web config 应用配置
package web

import (
	pkgconfig "github.com/ilaziness/gokit/config"
)

// Config 应用配置对象
var Config = &configType{}

// configType 配置项按需组合
type configType struct {
	App      *pkgconfig.App      `mapstructure:"app"`
	Db       *pkgconfig.DB       `mapstructure:"db"`
	Redis    *pkgconfig.Redis    `mapstructure:"redis"`
	RocketMq *pkgconfig.RocketMq `mapstructure:"rocket_mq"`
	Otel     *pkgconfig.Otel     `mapstructure:"otel"`
	Nacos    *pkgconfig.Nacos    `mapstructure:"nacos"`
}

func init() {
	pkgconfig.LoadConfig(Config)
}
