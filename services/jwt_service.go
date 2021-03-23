package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/entities"
)

type JwtService struct {
	ttl    time.Duration
	secret string
	issuer string
}

type authClaims struct {
	User *entities.UserClaims `json:"user"`
	jwt.StandardClaims
}

func NewJwtService(ttl time.Duration, secret string, issuer string) *JwtService {
	return &JwtService{
		ttl:    ttl,
		secret: secret,
		issuer: issuer,
	}
}

func (j *JwtService) GenerateJwtToken(user *entities.User) (*entities.AccessTokenResponse, error) {
	var tokenResp *entities.AccessTokenResponse
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(j.ttl).Unix()
	claims := authClaims{
		User: &entities.UserClaims{
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
	tokenResp = &entities.AccessTokenResponse{
		AccessToken: token,
		TTL:         int(j.ttl.Seconds()),
		ExpiredAt:   time.Now().Add(j.ttl),
		UserID:      user.ID.String(),
	}
	return tokenResp, nil
}

func (j *JwtService) VerifyToken(tokenString string) (*authClaims, error) {
	claims := &authClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	claims, ok := token.Claims.(*authClaims)
	if !ok {
		return nil, errors.New("claims retrieve failed")
	}
	return claims, nil
}
