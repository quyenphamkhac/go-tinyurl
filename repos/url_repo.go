package repos

import (
	"context"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/models"
	"github.com/teris-io/shortid"
)

type URLRespository struct {
	session   *gocql.Session
	cacheRepo *CacheRepository
}

func NewURLRepository(s *gocql.Session, c *CacheRepository) *URLRespository {
	return &URLRespository{
		session:   s,
		cacheRepo: c,
	}
}

func (r *URLRespository) GetURLByHash(hash string) (*models.URL, error) {
	var url *models.URL
	url = r.cacheRepo.GetURL(hash)
	if url != nil {
		return url, nil
	}
	m := map[string]interface{}{}
	var found bool = false
	query := "SELECT * FROM urls WHERE hash = ? LIMIT 1 ALLOW FILTERING"
	iterable := r.session.Query(query, hash).Iter()
	for iterable.MapScan(m) {
		found = true
		url = &models.URL{
			Hash:           m["hash"].(string),
			OriginalURL:    m["original_url"].(string),
			CreationDate:   m["creation_date"].(time.Time),
			ExpirationDate: m["expiration_date"].(time.Time),
			UserID:         m["user_id"].(string),
		}
	}
	if !found {
		return nil, errors.New("url not found")
	}
	r.cacheRepo.SetURL(url)
	return url, nil
}

func (r *URLRespository) GetUserURLByHash(hash string, user *models.UserClaims) (*models.URL, error) {
	var url *models.URL
	url = r.cacheRepo.GetURL(hash)
	if url != nil {
		return url, nil
	}
	m := map[string]interface{}{}
	var found bool = false
	query := "SELECT * FROM urls WHERE user_id = ? AND hash = ? LIMIT 1 ALLOW FILTERING"
	iterable := r.session.Query(query, user.UserID, hash).Iter()
	for iterable.MapScan(m) {
		found = true
		url = &models.URL{
			Hash:           m["hash"].(string),
			OriginalURL:    m["original_url"].(string),
			CreationDate:   m["creation_date"].(time.Time),
			ExpirationDate: m["expiration_date"].(time.Time),
			UserID:         m["user_id"].(string),
		}
	}
	if !found {
		return nil, errors.New("url not found")
	}
	r.cacheRepo.SetURL(url)
	return url, nil
}

func (r *URLRespository) GetAllURLs() []models.URL {
	var urls []models.URL
	m := map[string]interface{}{}
	query := "SELECT * FROM urls"
	iterable := r.session.Query(query).Iter()
	for iterable.MapScan(m) {
		urls = append(urls, models.URL{
			Hash:           m["hash"].(string),
			OriginalURL:    m["original_url"].(string),
			CreationDate:   m["creation_date"].(time.Time),
			ExpirationDate: m["expiration_date"].(time.Time),
			UserID:         m["user_id"].(string),
		})
	}
	return urls
}

func (r *URLRespository) CreateURL(createURLDto *dtos.CreateURLDto, user *models.User) (*models.URL, error) {
	hash, err := shortid.Generate()
	if err != nil {
		return nil, errors.New("can't generate new hash")
	}
	var tinyurl *models.URL
	var count int
	r.session.Query("SELECT COUNT(*) FROM urls WHERE user_id = ? AND original_url = ? ALLOW FILTERING", user.ID.String(), createURLDto.OriginalURL).Iter().Scan(&count)
	if count > 0 {
		return nil, errors.New("url already hashed")
	}
	tinyurl = &models.URL{
		Hash:           hash,
		OriginalURL:    createURLDto.OriginalURL,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().Add(14 * time.Hour * 24),
		UserID:         user.ID.String(),
	}
	ctx := context.Background()
	if err := r.session.Query(`INSERT INTO urls (hash, original_url, creation_date, expiration_date, user_id) VALUES (?, ?, ?, ?, ?)`,
		tinyurl.Hash, tinyurl.OriginalURL, tinyurl.CreationDate, tinyurl.ExpirationDate, tinyurl.UserID).WithContext(ctx).Exec(); err != nil {
		return nil, err
	}
	return tinyurl, nil
}
