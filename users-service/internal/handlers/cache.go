package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	cacheControlHeaderValue = "max-age=%d"
)

func setCacheHeader(c *gin.Context, seconds int) {
	c.Header("Cache-Control", fmt.Sprintf(cacheControlHeaderValue, seconds))
}
