package mysql

import (
	"fmt"
	"gintpl/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitMysql 初始化MySQL连接
// dns refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
func Init(cfg *config.DB) {
	dsn := cfg.DSN
	if dsn == "" {
		dsn = buildDSN(cfg)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func buildDSN(cfg *config.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
}
