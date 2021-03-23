package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/quyenphamkhac/go-tinyurl/entities"
)

type CacheRepository struct {
	cache *cache.Cache
}

func NewCacheRepository(c *cache.Cache) *CacheRepository {
	return &CacheRepository{
		cache: c,
	}
}

func (c *CacheRepository) SetURL(url *entities.URL) {
	ctx := context.TODO()
	if err := c.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   url.Hash,
		Value: url,
		TTL:   time.Hour,
	}); err != nil {
		fmt.Println("[Error] Redis set URL fails: ", err.Error())
	}
}

func (c *CacheRepository) GetURL(hash string) *entities.URL {
	var url entities.URL
	ctx := context.TODO()
	err := c.cache.Get(ctx, hash, &url)
	if err != nil {
		fmt.Println("[Error] Redis get URL fails: ", err.Error())
		return nil
	}
	return &url
}
