package service

import (
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/domain"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/repository"
)


type TaskService struct {
	repo     repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo:     repo,
	}
}

func (s *TaskService) CreateTask(task domain.Task) (int, error) {
	return s.repo.Insert(task)
}

func (s *TaskService) GetById(id int) (domain.Task, error) {
	return s.repo.GetById(id)
}

func (s *TaskService) GetAll() ([]domain.Task, error) {
	return s.repo.GetAll()
}