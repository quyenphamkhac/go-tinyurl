package db

import (
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/quyenphamkhac/go-tinyurl/config"
)

var (
	initRedisOnce sync.Once
	redisClient   *redis.Client
)

func GetRedisClient() *redis.Client {
	initRedisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.GetConfig().Addr,
		})
	})
	return redisClient
}

var (
	initCacheOnce sync.Once
	redisCache    *cache.Cache
)

func GetRedisCache() *cache.Cache {
	initCacheOnce.Do(func() {
		redisCache = cache.New(&cache.Options{
			Redis:      GetRedisClient(),
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
	})
	return redisCache
}
