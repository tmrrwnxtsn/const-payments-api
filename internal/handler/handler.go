package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Handler представляет обработчик запросов к API.
type Handler struct {
	service *service.Services
	logger  log.Logger
}

func NewHandler(services *service.Services, logger log.Logger) *Handler {
	return &Handler{service: services, logger: logger}
}

// InitRoutes инициализирует маршруты обработчика запросов.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		transactions := api.Group("/transactions")
		{
			transactions.POST("/", h.createTransaction)
			transactions.GET("/", h.getAllUserTransactions)
			transactions.GET("/:id/status", h.getTransactionStatus)
			transactions.PATCH("/:id/status", h.changeTransactionStatus)
			//transactions.DELETE("/:id/cancel", h.cancelTransaction)
		}
	}

	return router
}
