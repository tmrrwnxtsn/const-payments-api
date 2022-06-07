package handler

import (
	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status string `json:"status" example:"ok"`
}

type errorResponse struct {
	Message string `json:"message" example:"invalid transaction id"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: err.Error()})
}
