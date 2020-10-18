package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func responseAndLogError(c *gin.Context, statusCode int, handler, message string) {
	logrus.WithField("handler", handler).Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
