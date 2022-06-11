package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tmrrwnxtsn/const-payments-api/docs"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
)

// Handler представляет маршрутизатор.
type Handler struct {
	service *service.Services
	logger  *logrus.Logger
}

func NewHandler(services *service.Services, logger *logrus.Logger) *Handler {
	return &Handler{service: services, logger: logger}
}

// InitRoutes инициализирует маршруты.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// middleware
	router.Use(
		setRequestID,
		h.logRequest,
		corsMiddleware,
	)

	api := router.Group("/api")
	{
		transactions := api.Group("/transactions")
		{
			transactions.POST("/", h.createTransaction)
			transactions.GET("/", h.getAllUserTransactions)
			transactions.GET("/:id/status", h.getTransactionStatus)
			transactions.PATCH("/:id/status", h.changeTransactionStatus)
			transactions.DELETE("/:id", h.cancelTransaction)
		}
	}

	// Swagger-документация
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
