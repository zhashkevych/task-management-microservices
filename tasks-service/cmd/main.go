package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/config"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/handlers"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/repository/postgres"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/server"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	// postgres driver import
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("failed to init config: %s\n", err.Error())
	}

	logrus.SetLevel(logrus.Level(cfg.LoggerLevel))
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			"service": cfg.ServiceName,
		},
	})

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	taskRepo := postgres.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	handler := handlers.NewHandler(taskService)

	srv := server.NewServer(cfg, handler.Init())
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Infof("%s started", cfg.ServiceName)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}