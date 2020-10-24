package service

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/zhashkevych/task-management-microservices/sidecar/jwt"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	mock_repository "github.com/zhashkevych/task-management-microservices/users-service/internal/repository/mocks"
	"reflect"
	"testing"
	"time"
)

const (
	testTokenTtl = 10 * time.Second
)

func TestUserService_GenerateToken(t *testing.T) {
	// Init Mock Data and Test Table
	type mockBehaviour struct {
		mockRepo func(r *mock_repository.MockUserRepository, username, password string)
	}
	type args struct {
		username string
		password string
	}

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		want          jwt.AccessToken
		wantErr       bool
	}{
		{
			name: "Ok",
			mockBehaviour: mockBehaviour{
				mockRepo: func(r *mock_repository.MockUserRepository, username, password string) {
					r.EXPECT().Get(username, password).Return(domain.User{Id: 1}, nil)
				},
			},
			args: args{"test", "test"},
			want: jwt.AccessToken{
				UserId: 1,
				StandardClaims: jwt2.StandardClaims{
					ExpiresAt: time.Now().Add(testTokenTtl).Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockUserRepository(c)

			s := NewUserService(UserServiceDeps{
				Repo:     repo,
				TokenTtl: testTokenTtl,
			})

			passwordHashed := s.getPasswordHash(tt.args.password)
			tt.mockBehaviour.mockRepo(repo, tt.args.username, passwordHashed)

			token, err := s.GenerateToken(tt.args.username, tt.args.password)
			if err != nil && !tt.wantErr {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(token ,tt.want) {
				t.Fatalf("expected value doesn't match the result")
			}
		})
	}
}
