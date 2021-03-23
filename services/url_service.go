package services

import (
	"github.com/go-redis/cache/v8"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/entities"
	"github.com/quyenphamkhac/go-tinyurl/repos"
)

type URLService struct {
	repos *repos.URLRespository
	cache *cache.Cache
}

func NewUrlService(r *repos.URLRespository, c *cache.Cache) *URLService {
	return &URLService{
		repos: r,
		cache: c,
	}
}

func (s *URLService) GetAllURLs() []entities.URL {
	return s.repos.GetAllURLs()
}

func (s *URLService) CreateURL(createURLDto *dtos.CreateURLDto, user *entities.User) (*entities.URL, error) {
	return s.repos.CreateURL(createURLDto, user)
}

func (s *URLService) GetURLByHash(hash string, user *entities.UserClaims) (*entities.URL, error) {
	return s.repos.GetURLByHash(hash, user)
}
