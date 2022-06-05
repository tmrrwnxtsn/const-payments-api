package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
