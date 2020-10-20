package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/zhashkevych/task-management-microservices/sidecar/jwt"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/repository"
	"time"
)

type UserServiceDeps struct {
	Repo     repository.UserRepository
	Salt     string
	TokenTtl time.Duration
}

type UserService struct {
	repo     repository.UserRepository
	salt     string
	tokenTtl time.Duration
}

func NewUserService(deps UserServiceDeps) *UserService {
	return &UserService{
		repo:     deps.Repo,
		salt:     deps.Salt,
		tokenTtl: deps.TokenTtl,
	}
}

func (s *UserService) CreateUser(user domain.User) (int, error) {
	user.Password = s.getPasswordHash(user.Password)

	return s.repo.Insert(user)
}

func (s *UserService) GenerateToken(username, password string) (jwt.AccessToken, error) {
	user, err := s.repo.Get(username, s.getPasswordHash(password))
	if err != nil {
		return jwt.AccessToken{}, err
	}

	return jwt.New(jwt.TokenInput{
		UserId:    user.Id,
		ExpiresAt: time.Now().Add(s.tokenTtl).Unix(),
	}), nil
}

func (s *UserService) GetProfile(userId int) (domain.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(s.salt))

	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}
