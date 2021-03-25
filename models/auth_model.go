package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
)

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

type AuthClaims struct {
	User *UserClaims `json:"user"`
	jwt.StandardClaims
}

type AuthService interface {
	SignUp() (*User, error)
	Login(credentials *dtos.SignInDto) (*AccessTokenResponse, error)
}

type JwtService interface {
	GenerateJwtToken(user *User) (*AccessTokenResponse, error)
	VerifyToken(tokenString string) (*AuthClaims, error)
}
