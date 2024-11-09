// Package web config 应用配置
package web

import (
	"gintpl/pkg/config"
)

var Config = &configType{}

type configType struct {
	App   *config.App   `mapstructure:"app"`
	Db    *config.DB    `mapstructure:"db"`
	Redis *config.Redis `mapstructure:"redis"`
}

func init() {
	config.LoadConfig(Config)
}
