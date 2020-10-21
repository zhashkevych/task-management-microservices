package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/task-management-microservices/sidecar/jwt"
	"net/http"
	"strings"
)

const userIdCtx = "userId"

func userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		responseAndLogError(c, http.StatusUnauthorized, "userIdentity", "invalid Authorization value")
		return
	}

	if headerParts[0] != "Bearer" {
		responseAndLogError(c, http.StatusUnauthorized, "userIdentity", "invalid Authorization value")
		return
	}

	token, err := jwt.ParseToken(headerParts[1])
	if err != nil {
		responseAndLogError(c, http.StatusUnauthorized, "userIdentity", err.Error())
		return
	}

	c.Set(userIdCtx, token.UserId)
}