package entities

import "time"

type AccessTokenResponse struct {
	AccessToken string    `json:"access_token"`
	TTL         int       `json:"ttl"`
	ExpiredAt   time.Time `json:"expired_at"`
	UserID      string    `json:"user_id"`
}

type UserClaims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
}

type RefreshToken struct {
}
