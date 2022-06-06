package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"net/http"
	"strconv"
)

var (
	ErrInvalidQueryParams = errors.New("invalid query parameters data")
	ErrInvalidId          = errors.New("invalid id")
)

// createTransaction создаёт платёж (транзакцию).
func (h *Handler) createTransaction(c *gin.Context) {
	// принимает id пользователя, email пользователя, сумму и валюту платежа
	type createTransactionRequest struct {
		UserID       uint64  `json:"user_id,string" binding:"required"`
		UserEmail    string  `json:"user_email" binding:"required"`
		Amount       float64 `json:"amount,string" binding:"required"`
		CurrencyCode string  `json:"currency_code" binding:"required"`
	}

	var request createTransactionRequest
	if err := c.BindJSON(&request); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, service.ErrIncorrectTransactionData)
		return
	}

	transactionId, err := h.service.TransactionService.Create(model.Transaction{
		UserID:       request.UserID,
		UserEmail:    request.UserEmail,
		Amount:       request.Amount,
		CurrencyCode: request.CurrencyCode,
		Status:       model.StatusNew,
	})
	if err != nil {
		if err == service.ErrIncorrectTransactionData {
			h.newErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": transactionId,
	})
}

type getAllUserTransactionsResponse struct {
	Data []model.Transaction `json:"data"`
}

// getAllUserTransactions возвращает список всех платежей (транзакций) пользователя по его id или email.
func (h *Handler) getAllUserTransactions(c *gin.Context) {
	var transactions []model.Transaction

	userIDStr := c.Query("user_id")
	if userIDStr != "" {
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			h.newErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParams)
			return
		}

		transactions, err = h.service.TransactionService.GetAllByUserID(userID)
		if err != nil {
			h.newErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}

	userEmail := c.Query("user_email")
	if userIDStr == "" && userEmail != "" {
		var err error
		transactions, err = h.service.TransactionService.GetAllByUserEmail(userEmail)
		if err != nil {
			h.newErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, getAllUserTransactionsResponse{
		Data: transactions,
	})
}

// getTransactionStatus возвращает статус транзакции по её ID.
func (h *Handler) getTransactionStatus(c *gin.Context) {
	transactionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, ErrInvalidId)
		return
	}

	status, err := h.service.TransactionService.GetStatus(transactionId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
