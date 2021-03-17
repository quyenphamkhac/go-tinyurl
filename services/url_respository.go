package services

import (
	"github.com/quyenphamkhac/go-tinyurl/entities"
	"github.com/quyenphamkhac/go-tinyurl/respositories"
)

type UrlService struct {
	respository *respositories.UrlRespository
}

func NewUrlService(r *respositories.UrlRespository) *UrlService {
	return &UrlService{
		respository: r,
	}
}

func (u *UrlService) GetAllUrls() []*entities.Url {
	return []*entities.Url{}
}
