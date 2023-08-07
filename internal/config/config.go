// Package config 应用配置
package config

import (
	"fish/internal/bootstrap"
	"fish/internal/g"
	"github.com/spf13/viper"
	"os"
)

func init() {
	bootstrap.RegisterStartup(LoadConfig)
}

var (
	Gf  *viper.Viper
	App = &app{}
	DB  = &database{}
	// 默认配置文件名称
	defaultFileName = "config"
	// 配置文件列表
	files = make([]string, 0)
	// 配置文件搜索路径
	paths = []string{"config"}
)

// app 总体配置
type app struct {
	ID string `mapstructure:"id"`
	// App运行的根目录
	RootDir string
}

// database 数据库配置
// DSN和其他参数二选一配置
type database struct {
	DSN      string `mapstructure:"dsn"`
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Charset  string
}

// SetDefaultFile 设置默认配置文件名，不需要扩展名
func SetDefaultFile(file string) {
	defaultFileName = file
}

// AddFile 添加配置文件，不需要扩展名
func AddFile(file string) {
	files = append(files, file)
}

// AddPath 添加配置文件搜索路径
func AddPath(path string) {
	paths = append(paths, path)
}

func LoadConfig() {
	var err error
	App.RootDir, err = os.Getwd()
	if err != nil {
		g.Logger.Debug("set app run root directory error:", err)
	}
	g.Logger.Infoln("app run root directory:", App.RootDir)
	Gf = viper.New()
	Gf.SetConfigName(defaultFileName)
	for _, p := range paths {
		g.Logger.Infoln(p)
		Gf.AddConfigPath(p)
	}

	err = Gf.ReadInConfig()
	if err != nil {
		g.Logger.Errorf("read config error: %s", err)
	}
	for _, f := range files {
		Gf.SetConfigName(f)
		err = Gf.MergeInConfig()
		if err != nil {
			g.Logger.Errorf("merge config error: %s", err)
		}
	}
	err = Gf.UnmarshalKey("app", App)
	if err != nil {
		g.Logger.Errorf("decode config app part error: %s", err)
	}
	err = Gf.UnmarshalKey("database", DB)
	if err != nil {
		g.Logger.Errorf("decode config app part error: %v", err)
	}
}
