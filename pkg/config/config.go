// Package config provide confg file read and unmarshal
package config

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// EnvConfigFile 配置文件环境变量名称
const EnvConfigFile = "ENV_CONFIG_FILE"

var (
	// 配置文件列表
	files       = make([]string, 0)
	defaultFile = "config"
	// 配置文件目录
	defaultDir  = "./config"
	defaultType = "toml"
	configFile  = ""
)

type WithFunc func()

func scanFile() {
	// 优先环境变量指定的配置文件
	cfgFile := os.Getenv(EnvConfigFile)
	if cfgFile != "" {
		configFile = cfgFile
		return
	}
	var err error
	dirfs := os.DirFS(defaultDir)
	files, err = fs.Glob(dirfs, fmt.Sprintf("*.%s", defaultType))
	if err != nil {
		panic(err)
	}
}

// LoadConfig 读取解析配置文件
func LoadConfig[T any](c T) {
	scanFile()

	v := viper.New()
	if configFile != "" {
		v.SetConfigFile(configFile)
	} else {
		v.SetConfigName(defaultFile)
		v.AddConfigPath(defaultDir)
		v.SetConfigType(defaultType)
	}
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file, fmt.Sprintf("%s.%s", defaultFile, defaultType)) {
			continue
		}
		v.SetConfigFile(fmt.Sprintf("%s/%s", defaultDir, file))
		if err := v.MergeInConfig(); err != nil {
			panic(err)
		}
	}
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	v.Set("app.root_dir", rootDir)
	v.OnConfigChange(func(e fsnotify.Event) {
		slog.Info("Config file changed:" + e.Name)
		if err = v.Unmarshal(&c); err != nil {
			panic(err)
		}
	})

	if err = v.Unmarshal(&c); err != nil {
		panic(err)
	}
}
