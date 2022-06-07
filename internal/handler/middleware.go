package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const requestIDKey = "X-Request-ID"

// setRequestID присваивает входящему HTTP-запросу ID (requestIDKey).
func setRequestID(c *gin.Context) {
	requestID := uuid.New().String()
	c.Header(requestIDKey, requestID)
	c.Next()
}

// logRequest логгирует входящий HTTP-запрос.
func (h *Handler) logRequest(c *gin.Context) {
	logger := h.logger.WithFields(logrus.Fields{
		"remote_addr": c.Request.RemoteAddr,
		"request_id":  c.Writer.Header().Get(requestIDKey),
	})
	logger.Infof("started %s %s", c.Request.Method, c.Request.RequestURI)

	now := time.Now()

	// обработка запроса
	c.Next()

	// после обработки запроса
	latency := time.Since(now)

	// получаем статус выполнения запроса
	status := c.Writer.Status()

	var level logrus.Level
	switch {
	case status >= 500:
		level = logrus.ErrorLevel
	case status >= 400:
		level = logrus.WarnLevel
	default:
		level = logrus.InfoLevel
	}

	logger.Logf(
		level,
		"completed with %d %s in %v",
		status,
		http.StatusText(status),
		latency,
	)
}

// corsMiddleware включает CORS.
func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
