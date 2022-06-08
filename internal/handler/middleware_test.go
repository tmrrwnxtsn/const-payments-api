package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_setRequestID(t *testing.T) {
	router := gin.New()
	router.Use(setRequestID)
	router.GET("/id", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		"GET",
		"/id",
		nil,
	)

	router.ServeHTTP(responseRecorder, request)

	requestID := responseRecorder.Header().Get(requestIDKey)
	assert.NotEqual(t, "", requestID)

	request = httptest.NewRequest(
		"GET",
		"/id",
		nil,
	)
	request.Header.Set(requestIDKey, "test request ID")

	router.ServeHTTP(responseRecorder, request)

	requestID = responseRecorder.Header().Get(requestIDKey)
	assert.Equal(t, "test request ID", requestID)
}
