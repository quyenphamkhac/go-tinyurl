package repos

import (
	"time"

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
	m := map[string]interface{}{}
	query := "SELECT * FROM urls"
	iterable := r.session.Query(query).Iter()
	for iterable.MapScan(m) {
		urls = append(urls, entities.URL{
			Hash:           m["hash"].(string),
			OriginalURL:    m["original_url"].(string),
			CreationDate:   m["creation_date"].(time.Time),
			ExpirationDate: m["expiration_date"].(time.Time),
			UserID:         m["user_id"].(string),
		})
	}
	return urls
}
