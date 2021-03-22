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
	issuer string
}

type UserClaims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
}

type authClaims struct {
	User *UserClaims `json:"user"`
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
		User: &UserClaims{
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

func (j *JwtService) VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %s", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
