package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, error error) {
	h.logger.Error(error.Error())
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: error.Error()})
}
