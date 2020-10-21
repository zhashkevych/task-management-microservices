package repository

import "github.com/zhashkevych/task-management-microservices/tasks-service/internal/domain"

type TaskRepository interface {
	Insert(task domain.Task) (int, error)
	GetById(id int) (domain.Task, error)
	GetAll() ([]domain.Task, error)
}
