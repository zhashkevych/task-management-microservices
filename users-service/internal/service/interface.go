package service

import (
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type Users interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string)
	GetProfile(userId int) (domain.User, error)
}
