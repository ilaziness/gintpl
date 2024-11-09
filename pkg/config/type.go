package config

type Mode string

type App struct {
	ID string `mapstructure:"id"`
	// Mode debug, release
	Mode    string `mapstructure:"mode"`
	Port    uint16 `mapstructure:"port"`
	RootDir string `mapstructure:"root_dir"`
	Cors    *Cors  `mapstructure:"cors"`
}

type Cors struct {
	AllowOrigin      []string `mapstructure:"allowOrigin"`
	AllowMethods     []string `mapstructure:"allowMethods"`
	AllowHeaders     []string `mapstructure:"allowHeaders"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
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
