package repos

import (
	"context"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	session *gocql.Session
}

func NewUserRepository(s *gocql.Session) *UserRepository {
	return &UserRepository{
		session: s,
	}
}

func (r *UserRepository) CreateUser(userDto *dtos.SignUpDto) (*entities.User, error) {
	var user *entities.User
	var count int
	r.session.Query("SELECT COUNT(*) FROM users WHERER username = ? ALLOW FILTERING", userDto.Username).Iter().Scan(&count)
	if count > 0 {
		return user, errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user = &entities.User{
		Username:       userDto.Username,
		Name:           userDto.Name,
		HashedPassword: string(hashedPassword),
		Email:          userDto.Email,
		ID:             gocql.TimeUUID().String(),
		CreationDate:   time.Now(),
	}
	ctx := context.Background()
	if err := r.session.Query(`INSERT INTO users (username, name, hashed_password, email, id, creation_date) VALUES (?, ?, ?, ?, ?, ?)`,
		user.Username, user.Name, user.HashedPassword, user.Email, user.ID, user.CreationDate).WithContext(ctx).Exec(); err != nil {
		return nil, err
	}
	return user, nil
}
