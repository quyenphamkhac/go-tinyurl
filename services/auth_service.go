package services

import (
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/models"
)

type AuthService struct {
	userRepo   models.UserRepository
	jwtService models.JwtService
}

func NewAuthService(r models.UserRepository, j models.JwtService) *AuthService {
	return &AuthService{
		userRepo:   r,
		jwtService: j,
	}
}

func (s *AuthService) SignUp(userDto *dtos.SignUpDto) (*models.User, error) {
	return s.userRepo.CreateUser(userDto)
}

func (s *AuthService) Login(credentials *dtos.SignInDto) (*models.AccessTokenResponse, error) {
	var token *models.AccessTokenResponse
	user, err := s.userRepo.ValidateUser(credentials)
	if err != nil {
		return nil, err
	}
	token, err = s.jwtService.GenerateJwtToken(user)
	if err != nil {
		return nil, err
	}
	return token, nil
}
