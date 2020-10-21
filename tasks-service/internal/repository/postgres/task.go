package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/domain"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Insert(task domain.Task) (int, error) {
	var id int
	row := r.db.QueryRow("INSERT INTO tasks (title, user_id) VALUES ($1, $2) RETURNING id",
		task.Title, task.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Select(&tasks, "SELECT *  FROM tasks")

	return tasks, err
}

func (r *TaskRepository) GetById(id int) (domain.Task, error) {
	var task domain.Task
	err := r.db.Get(&task, "SELECT * FROM tasks WHERE id=$1", id)

	return task, err
}
