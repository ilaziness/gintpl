使用方法

```go
import (
    "context"
    "gintpl/pkg/storage/cache"
)

// 先初始化
cache.InitRedisCache(redis.Client)

//使用
ctx := context.Background()
ttl := 5 // 秒
cache.Set(ctx, "key1", 1, ttl)
cache.Get(ctx, "key1")
var v int
cache.GetScan(ctx, "key1", &v)
cache.Del(ctx, "key1")
cache.GetOrSet(ctx, "key1", ttl, function() any {
    return "val"
})
```