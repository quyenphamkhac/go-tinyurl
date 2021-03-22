package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/entities"
)

type JwtService struct {
	ttl    time.Duration
	secret string
}

func NewJwtService(ttl time.Duration, secret string) *JwtService {
	return &JwtService{
		ttl:    ttl,
		secret: secret,
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
		UserID:      user.ID.String(),
	}
	return tokenResp, nil
}

func (j *JwtService) VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %s", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
