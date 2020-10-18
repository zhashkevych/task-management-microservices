package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"net/http"
)

func (h *Handler) token(c *gin.Context) {

}

type signUpResponse struct {
	Id int `json:"id"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		responseAndLogError(c, http.StatusBadRequest, "signUp", "invalid input body")
		return
	}

	id, err := h.userService.CreateUser(input)
	if err != nil {
		responseAndLogError(c, http.StatusBadRequest, "signUp", "invalid input body")
		return
	}

	c.JSON(http.StatusOK, signUpResponse{id})
}

func (h *Handler) profile(c *gin.Context) {

}
