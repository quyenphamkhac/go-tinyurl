package repos

import (
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
}

func (c *CacheRepository) GetURL(hash string) *entities.URL {
	return nil
}
