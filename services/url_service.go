package services

import (
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/entities"
	"github.com/quyenphamkhac/go-tinyurl/repos"
)

type URLService struct {
	repos *repos.URLRespository
}

func NewUrlService(r *repos.URLRespository) *URLService {
	return &URLService{
		repos: r,
	}
}

func (s *URLService) GetAllURLs() []entities.URL {
	return s.repos.GetAllURLs()
}

func (s *URLService) CreateURL(createURLDto *dtos.CreateURLDto, user *entities.User) (*entities.URL, error) {
	return s.repos.CreateURL(createURLDto, user)
}

func (s *URLService) GetUserURLByHash(hash string, user *entities.UserClaims) (*entities.URL, error) {
	return s.repos.GetUserURLByHash(hash, user)
}

func (s *URLService) GetURLByHash(hash string) (*entities.URL, error) {
	return s.repos.GetURLByHash(hash)
}
