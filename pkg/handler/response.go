package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
)

type errorMsg struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger := logging.GetLogger()
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorMsg{message})
}
