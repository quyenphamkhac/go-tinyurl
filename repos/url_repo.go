package repos

import (
	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/entities"
)

type URLRespository struct {
	session *gocql.Session
}

func NewURLRepository(s *gocql.Session) *URLRespository {
	return &URLRespository{
		session: s,
	}
}

func (r *URLRespository) GetAllURLs() []entities.URL {
	var urls []entities.URL
	return urls
}
