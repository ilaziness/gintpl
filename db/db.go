package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/ilaziness/gintpl/db/ent"
	"github.com/ilaziness/gintpl/db/ent/migrate"
	"github.com/ilaziness/gokit/hook"
	"github.com/ilaziness/gokit/log"
)

var client *ent.Client

func GetClient() *ent.Client {
	return client
}

func SetClient(d *sql.Driver, debug bool) {
	if client != nil {
		return
	}
	options := make([]ent.Option, 0)
	options = append(options, ent.Driver(d))
	if debug {
		options = append(options, ent.Debug())
		options = append(options, ent.Log(func(a ...any) {
			log.Debug(context.Background(), fmt.Sprintf("%v", a))
		}))
	}
	client = ent.NewClient(options...)
	hook.Exit.Register(func() {
		_ = client.Close()
	})
	if debug {
		client = client.Debug()
	}
}

// AutoMigration 自动迁移
func AutoMigration() {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Logger.Errorf("failed creating schema resources: %v", err)
	}
}
