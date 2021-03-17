package entities

import "time"

type Url struct {
	Hash           string    `json:"hash,omitempty"`
	OriginalURL    string    `json:"original_url,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	UserID         string    `json:"user_id"`
}
