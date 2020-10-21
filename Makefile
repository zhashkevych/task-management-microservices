.SILENT:
.PHONY: run migrate-users migrate-tasks-drop migrate-tasks migrate-tasks-drop

help:
	bash ./help.sh

run:
	docker-compose up --remove-orphans --build

migrate-users:
	migrate -path ./users-service/schema -database 'postgres://postgres:qwerty@0.0.0.0:5433/users?sslmode=disable' up

migrate-users-drop:
	migrate -path ./users-service/schema -database 'postgres://postgres:qwerty@0.0.0.0:5433/users?sslmode=disable' drop

migrate-tasks:
	migrate -path ./tasks-service/schema -database 'postgres://postgres:qwerty@0.0.0.0:5433/tasks?sslmode=disable' up

migrate-tasks-drop:
	migrate -path ./tasks-service/schema -database 'postgres://postgres:qwerty@0.0.0.0:5433/tasks?sslmode=disable' drop