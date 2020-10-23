package postgres

import (
	"database/sql"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/repository"
	"reflect"
	"testing"
)

func TestUserRepository_Insert(t *testing.T) {
	// Init DB and Repo
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewUserRepository(db)

	// Create Test Table
	tests := []struct {
		name    string
		s       repository.UserRepository
		user    domain.User
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			user: domain.User{
				FirstName: "first_name",
				LastName:  "last_name",
				Username:  "username",
				Password:  "password",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").WithArgs("first_name", "last_name", "username", "password").WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			s:    s,
			user: domain.User{
				FirstName: "",
				LastName:  "",
				Username:  "username",
				Password:  "password",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").WithArgs("first_name", "last_name", "username", "password").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	// Run Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Insert(tt.user)
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

func TestUserRepository_Get(t *testing.T) {
	// Init DB and Repo
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewUserRepository(db)

	// Create Test Table
	type credentials struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		s       repository.UserRepository
		creds   credentials
		mock    func()
		want    domain.User
		wantErr bool
	}{
		{
			name:  "Ok",
			s:     s,
			creds: credentials{"test", "qwerty"},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id", "first_name", "last_name", "username"}).AddRow(1, "test name", "test last name", "test")
				mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)
			},
			want: domain.User{
				Id:        1,
				FirstName: "test name",
				LastName:  "test last name",
				Username:  "test",
			},
		},
		{
			name:  "Not Found",
			s:     s,
			creds: credentials{"test", ""},
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM users").WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	// Run Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Get(tt.creds.username, tt.creds.password)
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

func TestUserRepository_GetById(t *testing.T) {
	// Init DB and Repo
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewUserRepository(db)

	// Create Test Table
	tests := []struct {
		name    string
		s       repository.UserRepository
		id      int
		mock    func()
		want    domain.User
		wantErr bool
	}{
		{
			name: "Ok",
			s:    s,
			id:   1,
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id", "first_name", "last_name", "username"}).AddRow(1, "test name", "test last name", "test")
				mock.ExpectQuery("SELECT (.+) FROM users WHERE id=?").WithArgs(1).WillReturnRows(rows)
			},
			want: domain.User{
				Id:        1,
				FirstName: "test name",
				LastName:  "test last name",
				Username:  "test",
			},
		},
		{
			name: "Not Found",
			s:    s,
			id:   404,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM users WHERE id=?").WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	// Run Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetById(tt.id)
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
