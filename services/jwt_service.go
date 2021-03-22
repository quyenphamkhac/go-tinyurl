package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/entities"
)

type JwtService struct {
	ttl time.Duration
}

func NewJwtService(ttl time.Duration) *JwtService {
	return &JwtService{
		ttl: ttl,
	}
}

func (j *JwtService) GenerateJwtToken(user *entities.User) (*entities.AccessTokenResponse, error) {
	var tokenResp *entities.AccessTokenResponse
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(j.ttl).Unix()
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.GetConfig().Secret
	token, err := tokenWithClaims.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	tokenResp = &entities.AccessTokenResponse{
		AccessToken: token,
		TTL:         int(j.ttl.Seconds()),
		ExpiredAt:   time.Now().Add(j.ttl),
		UserID:      user.ID,
	}
	return tokenResp, nil
}
