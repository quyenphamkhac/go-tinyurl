package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/models"
)

type JwtService struct {
	ttl    time.Duration
	secret string
	issuer string
}

func NewJwtService(ttl time.Duration, secret string, issuer string) *JwtService {
	return &JwtService{
		ttl:    ttl,
		secret: secret,
		issuer: issuer,
	}
}

func (j *JwtService) GenerateJwtToken(user *models.User) (*models.AccessTokenResponse, error) {
	var tokenResp *models.AccessTokenResponse
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(j.ttl).Unix()
	claims := models.AuthClaims{
		User: &models.UserClaims{
			Username: user.Username,
			UserID:   user.ID.String(),
		},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
			Issuer:    j.issuer,
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.GetConfig().Secret
	token, err := tokenWithClaims.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	tokenResp = &models.AccessTokenResponse{
		AccessToken: token,
		TTL:         int(j.ttl.Seconds()),
		ExpiredAt:   time.Now().Add(j.ttl),
		UserID:      user.ID.String(),
	}
	return tokenResp, nil
}

func (j *JwtService) VerifyToken(tokenString string) (*models.AuthClaims, error) {
	claims := &models.AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	claims, ok := token.Claims.(*models.AuthClaims)
	if !ok {
		return nil, errors.New("claims retrieve failed")
	}
	return claims, nil
}
