package cache

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cacher interface {
	Set(ctx context.Context, key string, val any, sec int) error
	Del(ctx context.Context, key string)
	Get(ctx context.Context, key string) string
	GetScan(ctx context.Context, key string, dst any) error
	GetOrSet(ctx context.Context, key string, sec int, fn func() any) string
}

var rc *redisCache
var _ Cacher = &redisCache{}

type redisCache struct {
	client *redis.Client
}

func InitRedisCache(client *redis.Client) {
	rc = &redisCache{
		client,
	}
}

func toDuration(sec int) time.Duration {
	if sec < 1 {
		panic("cache ttl最小是1秒")
	}
	return time.Second * time.Duration(sec)
}

func (rc *redisCache) Set(ctx context.Context, key string, val any, sec int) error {
	result := rc.client.Set(ctx, key, val, toDuration(sec))
	if result.Err() == nil {
		return nil
	}
	return fmt.Errorf("cache set error: %w", result.Err())
}

func (rc *redisCache) Del(ctx context.Context, key string) {
	_ = rc.client.Del(ctx, key)
}

func (rc *redisCache) Get(ctx context.Context, key string) string {
	result := rc.client.Get(ctx, key)
	return result.Val()
}

func (rc *redisCache) GetScan(ctx context.Context, key string, dst any) error {
	return rc.client.Get(ctx, key).Scan(&dst)
}

func (rc *redisCache) GetOrSet(ctx context.Context, key string, sec int, fn func() any) string {
	if rc.client.Exists(ctx, key).Val() == 1 {
		return rc.client.Get(ctx, key).Val()
	}
	var result *redis.StringCmd
	var val = fn()
	_, err := rc.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.Set(ctx, key, val, toDuration(sec))
		result = rdb.Get(ctx, key)
		return nil
	})
	if err != nil {
		slog.Info(fmt.Sprintf("GetOrSet error: %v", err))
	}
	return result.Val()
}
