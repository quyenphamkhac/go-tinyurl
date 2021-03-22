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
	r.session.Query("SELECT COUNT(*) FROM users WHERE username = ? ALLOW FILTERING", userDto.Username).Iter().Scan(&count)
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
		LastLogin:      time.Now(),
	}
	ctx := context.Background()
	if err := r.session.Query(`INSERT INTO users (username, name, hashed_password, email, id, creation_date, last_login) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		user.Username, user.Name, user.HashedPassword, user.Email, user.ID, user.CreationDate, user.LastLogin).WithContext(ctx).Exec(); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ValidateUser(credentials *dtos.SignInDto) (*entities.User, error) {
	var user *entities.User
	var found bool = false
	m := map[string]interface{}{}
	query := `SELECT * FROM users WHERE username = ? LIMIT 1`
	iterable := r.session.Query(query, credentials.Username).Consistency(gocql.One).Iter()
	for iterable.MapScan(m) {
		found = true
		user = &entities.User{
			ID:             m["id"].(string),
			HashedPassword: m["hashed_password"].(string),
			Username:       m["username"].(string),
			Name:           m["name"].(string),
			Email:          m["email"].(string),
			CreationDate:   m["creation_date"].(time.Time),
			LastLogin:      m["last_login"].(time.Time),
		}
	}
	if !found {
		return nil, errors.New("user not found")
	}
	isValid, err := user.ComparePassword(credentials.Password)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
