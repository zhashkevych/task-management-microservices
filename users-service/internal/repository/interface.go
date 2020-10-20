package repository

import "github.com/zhashkevych/task-management-microservices/users-service/internal/domain"

type UserRepository interface {
	Insert(user domain.User) (int, error)
	GetById(id int) (domain.User, error)
	Get(username, password string) (domain.User, error)
}
