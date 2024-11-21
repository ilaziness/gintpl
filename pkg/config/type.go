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
	AllowOrigin      []string `mapstructure:"allow_origin"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
}

type DB struct {
	DSN      string `mapstructure:"dsn"`
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"Port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
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

type RocketMq struct {
	Endpoint      string `mapstructure:"endpoint"`
	AccessKey     string `mapstructure:"access_key"`
	SecretKey     string `mapstructure:"secret_key"`
	Transaction   bool   `mapstructure:"transaction"`
	ProducerTopic string `mapstructure:"producer_topic"`
}
