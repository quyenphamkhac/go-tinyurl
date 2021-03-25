package services

import (
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/models"
)

type URLService struct {
	urlRepo models.URLRepository
}

func NewUrlService(r models.URLRepository) *URLService {
	return &URLService{
		urlRepo: r,
	}
}

func (s *URLService) GetAllURLs() ([]models.URL, error) {
	return s.urlRepo.GetAllURLs()
}

func (s *URLService) CreateURL(createURLDto *dtos.CreateURLDto, user *models.User) (*models.URL, error) {
	return s.urlRepo.CreateURL(createURLDto, user)
}

func (s *URLService) GetUserURLByHash(hash string, user *models.UserClaims) (*models.URL, error) {
	return s.urlRepo.GetUserURLByHash(hash, user)
}

func (s *URLService) GetURLByHash(hash string) (*models.URL, error) {
	return s.urlRepo.GetURLByHash(hash)
}
