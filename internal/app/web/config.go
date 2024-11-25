// Package web config 应用配置
package web

import (
	"gintpl/pkg/config"
)

var Config = &configType{}

// configType 配置项按需组合
type configType struct {
	App      *config.App      `mapstructure:"app"`
	Db       *config.DB       `mapstructure:"db"`
	Redis    *config.Redis    `mapstructure:"redis"`
	RocketMq *config.RocketMq `mapstructure:"rocket_mq"`
	Otel     *config.Otel     `mapstructure:"otel"`
}

func init() {
	config.LoadConfig(Config)
}
