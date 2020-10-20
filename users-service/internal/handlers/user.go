package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/task-management-microservices/sidecar/jwt"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"net/http"
)

type signUpResponse struct {
	Id int `json:"id"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		responseAndLogError(c, http.StatusBadRequest, "signUp", "invalid input body: "+err.Error())
		return
	}

	id, err := h.userService.CreateUser(input)
	if err != nil {
		responseAndLogError(c, http.StatusInternalServerError, "signUp", err.Error())
		return
	}

	c.JSON(http.StatusOK, signUpResponse{id})
}

type tokenInput struct {
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type tokenResponse struct {
	AccessToken jwt.AccessToken `json:"access_token"`
}

func (h *Handler) token(c *gin.Context) {
	var input tokenInput
	if err := c.BindJSON(&input); err != nil {
		responseAndLogError(c, http.StatusBadRequest, "token", "invalid input body: "+err.Error())
		return
	}

	token, err := h.userService.GenerateToken(input.Username, input.Password)
	if err != nil {
		responseAndLogError(c, http.StatusInternalServerError, "token", err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{token})
}

func (h *Handler) profile(c *gin.Context) {
	userId, _ := c.Get(userIdCtx)

	profile, err := h.userService.GetProfile(userId.(int))
	if err != nil {
		responseAndLogError(c, http.StatusInternalServerError, "profile", err.Error())
		return
	}

	c.JSON(http.StatusOK, profile)
}
