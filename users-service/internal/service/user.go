package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
	salt string
}

func NewUserService(r repository.UserRepository, salt string) *UserService {
	return &UserService{
		repo: r,
		salt: salt,
	}
}

func (s *UserService) CreateUser(user domain.User) (int, error) {
	user.Password = s.getPasswordHash(user.Password)

	return s.repo.Insert(user)
}

func (s *UserService) getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(s.salt))

	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}
