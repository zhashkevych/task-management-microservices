package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/service"
	"net/http"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{userService: userService}
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
	{
		v1.POST("/sign-up", h.signUp)
		v1.GET("/token", h.token)
		v1.GET("/profile", h.profile)
	}
}
