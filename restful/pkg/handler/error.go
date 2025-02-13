package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Log.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
