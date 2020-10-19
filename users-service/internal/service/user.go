package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/jwt"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/repository"
	"strconv"
)

type UserServiceDeps struct {
	Repo   repository.UserRepository
	Issuer *jwt.Issuer
	Salt   string
}

type UserService struct {
	repo   repository.UserRepository
	issuer *jwt.Issuer
	salt   string
}

func NewUserService(deps UserServiceDeps) *UserService {
	return &UserService{
		repo:   deps.Repo,
		salt:   deps.Salt,
		issuer: deps.Issuer,
	}
}

func (s *UserService) CreateUser(user domain.User) (int, error) {
	user.Password = s.getPasswordHash(user.Password)

	return s.repo.Insert(user)
}

func (s *UserService) GenerateToken(username, password string) (domain.AccessToken, error) {
	user, err := s.repo.Get(username, s.getPasswordHash(password))
	if err != nil {
		return domain.AccessToken{}, err
	}

	//TODO implement refresh tokens + jti

	accessToken := s.issuer.Issue(parseId(user.Id))
	return accessToken, nil
}

func (s *UserService) getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(s.salt))

	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}

func parseId(id int) string {
	return strconv.Itoa(id)
}