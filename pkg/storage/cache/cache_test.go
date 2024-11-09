package cache

import (
	"context"
	"encoding"
	"encoding/json"
	"gintpl/pkg/config"
	"gintpl/pkg/storage/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cfg = &config.Redis{
	Host: "127.0.0.1",
}

var _ encoding.BinaryMarshaler = &tests{}

type tests struct {
	Name string
	Age  int
}

func (ts *tests) MarshalBinary() (data []byte, err error) {
	return json.Marshal(ts)
}

func (ts *tests) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ts)
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

func TestCache(t *testing.T) {
	ttl := 5
	ctx := context.Background()

	k1 := "k1"
	v1 := "abcr"
	err := Set(ctx, k1, v1, ttl)
	assert.Equal(t, err, nil, "set return error")
	var v1s string
	GetScan(ctx, k1, &v1s)
	assert.Equal(t, v1, v1s)

	k2 := "k2"
	v2 := 2
	err = Set(ctx, k2, v2, ttl)
	assert.Equal(t, err, nil, "set return error")
	var v2s int
	GetScan(ctx, k2, &v2s)
	assert.Equal(t, v2, v2s)

	k3 := "k3"
	v3 := &tests{"namek3", 15}
	assert.Equal(t, Set(ctx, k3, v3, ttl), nil, "set return error")
	v3s := &tests{}
	GetScan(ctx, k3, v3s)
	assert.Equal(t, v3, v3s)

	k4 := "k4"
	v4 := "k4k4k4"
	assert.Equal(t, v4, GetOrSet(ctx, k4, ttl, func() any {
		return v4
	}))
}
