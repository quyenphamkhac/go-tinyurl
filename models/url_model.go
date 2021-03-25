package models

import (
	"time"

	"github.com/quyenphamkhac/go-tinyurl/dtos"
)

type URL struct {
	Hash           string    `json:"hash,omitempty"`
	OriginalURL    string    `json:"original_url,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	UserID         string    `json:"user_id"`
}

type URLRepository interface {
	GetUserURLByHash(hash string, user *User) (*URL, error)
	GetURLByHash(hash string) (*URL, error)
	GetAllURLs() ([]URL, error)
	CreateURL(createURLDto *dtos.CreateURLDto) (*URL, error)
}

type URLService interface {
	GetUserURLByHash(hash string, user *User) (*URL, error)
	GetURLByHash(hash string) (*URL, error)
	GetAllURLs() ([]URL, error)
	CreateURL(createURLDto *dtos.CreateURLDto) (*URL, error)
}
