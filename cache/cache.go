package cache

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"
	
	"github.com/go-redis/redis/v8"
	
	"github.com/daarlabs/arcanum/cache/memory"
)

type Client interface {
	Exists(key string) bool
	Get(key string, data any) error
	Set(key string, data any, expiration time.Duration) error
	Destroy(key string) error
	
	MustGet(key string, data any)
	MustSet(key string, data any, expiration time.Duration)
	MustDestroy(key string)
}

type cache struct {
	ctx     context.Context
	adapter string
	memory  *memory.Client
	redis   *redis.Client
}

const (
	AdapterMemory = "memory"
	AdapterRedis  = "redis"
)

var (
	defaultMemoryCacheDir = os.TempDir() + "/.arcanum/cache/"
)

var (
	ErrorAdapterInstanceNotExist = errors.New("cache adapter instance not exist")
)

func New(ctx context.Context, mem *memory.Client, redis *redis.Client) Client {
	var adapter string
	if redis == nil {
		if mem == nil {
			mem = memory.New(defaultMemoryCacheDir)
		}
		adapter = AdapterMemory
	}
	if redis != nil {
		adapter = AdapterRedis
	}
	return &cache{
		ctx:     ctx,
		adapter: adapter,
		memory:  mem,
		redis:   redis,
	}
}

func (c cache) Exists(key string) bool {
	if c.isNil() {
		return false
	}
	switch c.adapter {
	case AdapterMemory:
		return c.memory.Exists(key)
	case AdapterRedis:
		cmd := c.redis.Exists(c.ctx, key)
		if cmd == nil {
			return false
		}
		return cmd.Val() > 0
	default:
		return false
	}
}

func (c cache) Get(key string, data any) error {
	if c.isNil() {
		return ErrorAdapterInstanceNotExist
	}
	var value string
	switch c.adapter {
	case AdapterMemory:
		value = c.memory.Get(key)
	case AdapterRedis:
		stored := c.redis.Get(c.ctx, key)
		if stored == nil {
			return nil
		}
		value = stored.Val()
	}
	if len(value) > 0 {
		return json.Unmarshal([]byte(value), data)
	}
	return nil
}

func (c cache) MustGet(key string, data any) {
	if err := c.Get(key, data); err != nil {
		panic(err)
	}
}

func (c cache) Set(key string, data any, expiration time.Duration) error {
	if c.isNil() {
		return ErrorAdapterInstanceNotExist
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	switch c.adapter {
	case AdapterMemory:
		return c.memory.Set(key, string(b), expiration)
	case AdapterRedis:
		if c.redis.Set(c.ctx, key, string(b), expiration).Err() != nil {
			return err
		}
		return nil
	}
	return nil
}

func (c cache) MustSet(key string, data any, expiration time.Duration) {
	if err := c.Set(key, data, expiration); err != nil {
		panic(err)
	}
}

func (c cache) Destroy(key string) error {
	switch c.adapter {
	case AdapterMemory:
		return c.memory.Destroy(key)
	case AdapterRedis:
		return c.Set(key, nil, time.Millisecond)
	}
	return nil
}

func (c cache) MustDestroy(key string) {
	if err := c.Destroy(key); err != nil {
		panic(err)
	}
}

func (c cache) isNil() bool {
	switch c.adapter {
	case AdapterMemory:
		return c.memory == nil
	case AdapterRedis:
		return c.redis == nil
	default:
		return true
	}
}
