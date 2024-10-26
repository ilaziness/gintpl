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

var (
	Gf *viper.Viper

	// 配置文件列表
	files = make([]string, 0)
	// 配置文件搜索路径
	paths = []string{"config"}

	mainFile = "config"
	// 配置文件目录
	configDir = "./config"
	fileType  = "toml"
)

type WithFunc func()

func scanFile() {
	var err error
	dirfs := os.DirFS(configDir)
	files, err = fs.Glob(dirfs, fmt.Sprintf("*.%s", fileType))
	if err != nil {
		panic(err)
	}
}

// AddPath 添加配置文件搜索路径
func AddPath(path string) {
	paths = append(paths, path)
}

// LoadConfig 读取解析配置文件
func LoadConfig[T any](c T) {
	scanFile()

	v := viper.New()
	v.SetConfigName(mainFile)
	v.AddConfigPath(configDir)
	v.SetConfigType(fileType)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file, fmt.Sprintf("%s.%s", mainFile, fileType)) {
			continue
		}
		v.SetConfigFile(fmt.Sprintf("%s/%s", configDir, file))
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
