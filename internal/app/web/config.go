// Package web config 应用配置
package web

import (
	config2 "github.com/ilaziness/gokit/config"
)

var Config = &configType{}

// configType 配置项按需组合
type configType struct {
	App      *config2.App      `mapstructure:"app"`
	Db       *config2.DB       `mapstructure:"db"`
	Redis    *config2.Redis    `mapstructure:"redis"`
	RocketMq *config2.RocketMq `mapstructure:"rocket_mq"`
	Otel     *config2.Otel     `mapstructure:"otel"`
	Nacos    *config2.Nacos    `mapstructure:"nacos"`
}

func init() {
	config2.LoadConfig(Config)
}
