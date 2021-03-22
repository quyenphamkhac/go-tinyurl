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

func (s *URLService) CreateURL(createURLDto *dtos.CreateURLDto) (*entities.URL, error) {
	return s.repos.CreateURL(createURLDto)
}
