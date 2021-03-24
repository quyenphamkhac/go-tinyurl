package repos

import (
	"context"
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
		return
	}
}

func (c *CacheRepository) GetURL(hash string) *entities.URL {
	var url entities.URL
	ctx := context.TODO()
	err := c.cache.Get(ctx, hash, &url)
	if err != nil {
		return nil
	}
	return &url
}
