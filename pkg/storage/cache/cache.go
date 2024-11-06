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

var cacheInst Cacher
var _ Cacher = &redisCache{}

func Set(ctx context.Context, key string, val any, sec int) error {
	return cacheInst.Set(ctx, key, val, sec)
}
func Del(ctx context.Context, key string) {
	cacheInst.Del(ctx, key)
}
func Get(ctx context.Context, key string) string {
	return cacheInst.Get(ctx, key)
}
func GetScan(ctx context.Context, key string, dst any) error {
	return cacheInst.GetScan(ctx, key, dst)
}
func GetOrSet(ctx context.Context, key string, sec int, fn func() any) string {
	return cacheInst.GetOrSet(ctx, key, sec, fn)
}

// redisCache redis缓存实现
type redisCache struct {
	client *redis.Client
}

// InitRedisCache 创建redis缓存
func InitRedisCache(client *redis.Client) {
	cacheInst = &redisCache{
		client,
	}
}

// toDuration 转换时间为duration类型
func toDuration(sec int) time.Duration {
	if sec < 1 {
		panic("cache ttl最小是1秒")
	}
	return time.Second * time.Duration(sec)
}

// Set 设置有效时长是sec秒的缓存
func (rc *redisCache) Set(ctx context.Context, key string, val any, sec int) error {
	result := rc.client.Set(ctx, key, val, toDuration(sec))
	if result.Err() == nil {
		return nil
	}
	return fmt.Errorf("cache set error: %w", result.Err())
}

// Del 删除缓存key
func (rc *redisCache) Del(ctx context.Context, key string) {
	_ = rc.client.Del(ctx, key)
}

// Get 获取缓存Key
func (rc *redisCache) Get(ctx context.Context, key string) string {
	result := rc.client.Get(ctx, key)
	return result.Val()
}

// GetScan 获取缓存key，并把获取到的值反序列化为dst
func (rc *redisCache) GetScan(ctx context.Context, key string, dst any) error {
	return rc.client.Get(ctx, key).Scan(dst)
}

// GetOrSet 获取缓存，不存在则使用fn函数获取值，再写入缓存
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
