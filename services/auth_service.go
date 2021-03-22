package services

import (
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/entities"
	"github.com/quyenphamkhac/go-tinyurl/repos"
)

type AuthService struct {
	userRepo   *repos.UserRepository
	jwtService *JwtService
}

func NewAuthService(r *repos.UserRepository, j *JwtService) *AuthService {
	return &AuthService{
		userRepo:   r,
		jwtService: j,
	}
}

func (s *AuthService) SignUp(userDto *dtos.SignUpDto) (*entities.User, error) {
	return s.userRepo.CreateUser(userDto)
}

func (s *AuthService) Login(credentials *dtos.SignInDto) (*entities.AccessTokenResponse, error) {
	var token *entities.AccessTokenResponse
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
