package services

import (
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/models"
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

func (s *URLService) GetAllURLs() []models.URL {
	return s.repos.GetAllURLs()
}

func (s *URLService) CreateURL(createURLDto *dtos.CreateURLDto, user *models.User) (*models.URL, error) {
	return s.repos.CreateURL(createURLDto, user)
}

func (s *URLService) GetUserURLByHash(hash string, user *models.UserClaims) (*models.URL, error) {
	return s.repos.GetUserURLByHash(hash, user)
}

func (s *URLService) GetURLByHash(hash string) (*models.URL, error) {
	return s.repos.GetURLByHash(hash)
}
