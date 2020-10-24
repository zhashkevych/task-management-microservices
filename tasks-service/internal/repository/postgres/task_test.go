package postgres

import (
	"github.com/zhashkevych/go-sqlxmock"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/domain"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/repository"
	"reflect"
	"testing"
	"time"
)

func TestTaskRepository_Insert(t *testing.T) {
	// Init DB and Repo
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%repo' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTaskRepository(db)

	// Create Test Table
	tests := []struct {
		name    string
		repo    repository.TaskRepository
		task    domain.Task
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			repo: repo,
			task: domain.Task{
				Title:  "test",
				UserId: 1,
			},
			mock: func() {
				//We added one row
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO tasks").WithArgs("test", 1).WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			repo: repo,
			task: domain.Task{},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").WithArgs("", 0).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	// Run Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.repo.Insert(tt.task)
			if err != nil && !tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepository_GetAll(t *testing.T) {
	// Init DB and Repo
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%repo' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTaskRepository(db)

	createdAt := time.Now()

	// Create Test Table
	tests := []struct {
		name    string
		repo    repository.TaskRepository
		mock    func()
		want    []domain.Task
		wantErr bool
	}{
		{
			name: "OK",
			repo: repo,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "user_id", "created_at"}).
					AddRow(1, "test1", 1, createdAt).
					AddRow(2, "test2", 2, createdAt)
				mock.ExpectQuery("SELECT (.+) FROM tasks").WillReturnRows(rows)
			},
			want: []domain.Task{
				{
					Id:        1,
					Title:     "test1",
					UserId:    1,
					CreatedAt: createdAt,
				},
				{
					Id:        2,
					Title:     "test2",
					UserId:    2,
					CreatedAt: createdAt,
				},
			},
		},
	}

	// Run Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.repo.GetAll()
			if err != nil && !tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepository_GetById(t *testing.T) {
	// Init DB and Repo
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%repo' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTaskRepository(db)

	createdAt := time.Now()

	// Create Test Table
	tests := []struct {
		name    string
		repo    repository.TaskRepository
		mock    func()
		id      int
		want    domain.Task
		wantErr bool
	}{
		{
			name: "OK",
			repo: repo,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "user_id", "created_at"}).
					AddRow(1, "test1", 1, createdAt)
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(1).WillReturnRows(rows)
			},
			id: 1,
			want: domain.Task{
				Id:        1,
				Title:     "test1",
				UserId:    1,
				CreatedAt: createdAt,
			},
		},
		{
			name: "Not Found",
			repo: repo,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "user_id", "created_at"})
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(404).WillReturnRows(rows)
			},
			id:      404,
			want:    domain.Task{},
			wantErr: true,
		},
	}

	// Run Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.repo.GetById(tt.id)
			if err != nil && !tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
