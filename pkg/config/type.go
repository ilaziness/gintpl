package config

type Mode string

const (
	Dev  Mode = "dev"
	Prod Mode = "prod"
)

type App struct {
	ID string `mapstructure:"id"`
	// Mode debug, release
	Mode    Mode   `mapstructure:"mode"`
	Port    uint16 `mapstructure:"port"`
	RootDir string `mapstructure:"root_dir"`
}

type DB struct {
	DSN      string `mapstructure:"dsn"`
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"Port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
	Charset  string `mapstructure:"charset"`
	Timezone string `mapstructure:"timezone"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Db   uint   `mapstructure:"db"`
}
