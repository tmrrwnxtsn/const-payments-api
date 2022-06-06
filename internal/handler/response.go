package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, err error) {
	h.logger.Error(err.Error())
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: err.Error()})
}
