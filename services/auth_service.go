package services

import (
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/entities"
	"github.com/quyenphamkhac/go-tinyurl/repos"
)

type AuthService struct {
	userRepo *repos.UserRepository
}

func NewAuthService(r *repos.UserRepository) *AuthService {
	return &AuthService{
		userRepo: r,
	}
}

func (s *AuthService) SignUp(userDto *dtos.SignUpDto) (*entities.User, error) {
	return s.userRepo.CreateUser(userDto)
}
