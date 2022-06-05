package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Handler представляет обработчик запросов к API.
type Handler struct {
	service service.Service
	logger  log.Logger
}

func NewHandler(services service.Service, logger log.Logger) *Handler {
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
			//transactions.GET("/:id", h.checkTransactionStatus)
			//transactions.PUT("/:id", h.updateTransactionStatus)
			//transactions.DELETE("/:id", h.cancelTransaction)
		}
	}

	return router
}
