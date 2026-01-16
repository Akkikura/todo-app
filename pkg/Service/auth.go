package Service

import (
	"crypto/sha1"
	"fmt"

	todo_app "github.com/akkikura/todo-app"
	"github.com/akkikura/todo-app/pkg/Repository"
)

const salt = "qkwjjkasdadqpwod21938"

type AuthService struct {
	repo Repository.Authorization
}

func NewAuthService(repo Repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
