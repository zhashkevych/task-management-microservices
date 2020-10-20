package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(user domain.User) (int, error) {
	var id int
	row := r.db.QueryRow("INSERT INTO users (first_name, last_name, username, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) Get(username, password string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT id, first_name, last_name, username  FROM users WHERE username=$1 AND password=$2", username, password)

	return user, err
}

func (r *UserRepository) GetById(id int) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT id, first_name, last_name, username FROM users WHERE id=$1", id)

	return user, err
}
