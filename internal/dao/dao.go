package dao

import (
	"context"

	"github.com/ilaziness/gintpl/internal/ent"
	"github.com/ilaziness/gintpl/internal/ent/migrate"
	"github.com/ilaziness/gokit/hook"
	"github.com/ilaziness/gokit/log"
)

var client *ent.Client

func SetClient(c *ent.Client) {
	if c != nil {
		client = c
		hook.Exit.Register(func() {
			_ = client.Close()
		})
	}
}

// AutoMigration 自动迁移
func AutoMigration() {
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Logger.Errorf("failed creating schema resources: %v", err)
	}
}
