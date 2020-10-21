package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/service"
	"net/http"
)

type Handler struct {
	taskService *service.TaskService
}

func NewHandler(taskService *service.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

func (h *Handler) Init() *gin.Engine {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initV1Routes(router)

	return router
}

func (h *Handler) initV1Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.POST("/tasks", userIdentity, h.createTask)
	v1.GET("/tasks", h.getAllTasks)
	v1.GET("/task/:id", h.getTaskById)
}
