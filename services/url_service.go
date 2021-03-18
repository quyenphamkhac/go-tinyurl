package services

import (
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

func (u *URLService) GetAllURLs() []entities.URL {
	return u.repos.GetAllURLs()
}
