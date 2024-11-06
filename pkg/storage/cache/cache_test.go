package cache

import (
	"context"
	"gintpl/pkg/config"
	"gintpl/pkg/storage/redis"
	"testing"
)

var cfg = &config.Redis{
	Host: "127.0.0.1",
}

func init() {
	redis.Init(cfg)
	InitRedisCache(redis.Client)
}

func TestMain(t *testing.T) {
	ctx := context.Background()
	ttl := 5

	t.Run("int", func(t *testing.T) {
		key := "keyint"
		if err := Set(ctx, key, 1, ttl); err != nil {
			t.Error(err.Error())
		}
		get := Get(ctx, key)
		if get != "1" {
			t.Errorf(`expect <"1">, got <%s>`, get)
		}
	})

	t.Run("set_or_get", func(t *testing.T) {
		v1 := GetOrSet(ctx, "stg1", ttl, func() any {
			return 2
		})
		if v1 != "2" {
			t.Errorf(`expect <"2">, got <%s>`, v1)
		}
	})

	t.Run("get_scan", func(t *testing.T) {
		k1 := "gc1"
		v1 := Set(ctx, k1, "abc", ttl)
		var vg1 string
		if err := GetScan(ctx, k1, &vg1); err != nil {
			t.Error(err)
		}
		if vg1 != "abc" {
			t.Errorf(`expect <"abc">, got <%s>`, v1)
		}
	})

}
