package models

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             gocql.UUID `json:"id,omitempty"`
	Username       string     `json:"username,omitempty"`
	Name           string     `json:"name,omitempty"`
	HashedPassword string     `json:"-"`
	Email          string     `json:"email,omitempty"`
	CreationDate   time.Time  `json:"creation_date,omitempty"`
	LastLogin      time.Time  `json:"last_login,omitempty"`
}

type UserRepository interface {
	CreateUser(createUserDto *dtos.SignUpDto) (*User, error)
	ValidateUser(credentials *dtos.SignInDto) (*User, error)
}

func (u *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
