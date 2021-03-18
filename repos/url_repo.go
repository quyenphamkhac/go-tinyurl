package repos

import (
	"time"

	"github.com/quyenphamkhac/go-tinyurl/entities"
)

type URLRespository struct {
}

func (r *URLRespository) GetAllURLs() []*entities.URL {
	return []*entities.URL{
		{
			UserID:         "1",
			OriginalURL:    "https://google.com",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(604800 * time.Second),
			Hash:           "4e932bc",
		},
	}
}
