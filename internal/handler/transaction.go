package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"net/http"
	"strconv"
)

// createTransaction создаёт платёж (транзакцию).
func (h *Handler) createTransaction(c *gin.Context) {
	// принимает ID пользователя, email пользователя, сумму и валюту платежа
	type createTransactionRequest struct {
		UserID       uint64  `json:"user_id,string" binding:"required"`
		UserEmail    string  `json:"user_email" binding:"required"`
		Amount       float64 `json:"amount,string" binding:"required"`
		CurrencyCode string  `json:"currency_code" binding:"required"`
	}
	type createTransactionResponse struct {
		ID uint64 `json:"id"`
	}

	var request createTransactionRequest
	if err := c.BindJSON(&request); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, service.ErrIncorrectTransactionData)
		return
	}

	transactionID, err := h.service.TransactionService.Create(model.Transaction{
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

	c.JSON(http.StatusCreated, createTransactionResponse{
		ID: transactionID,
	})
}

type getAllUserTransactionsResponse struct {
	Data []model.Transaction `json:"data"`
}

// getAllUserTransactions возвращает список всех платежей (транзакций) пользователя по его ID или email.
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
	type getTransactionStatusResponse struct {
		Status model.Status `json:"status"`
	}
	transactionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, ErrInvalidTransactionID)
		return
	}

	status, err := h.service.TransactionService.GetStatus(transactionId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTransactionStatusResponse{
		Status: status,
	})
}

// changeTransactionStatus обновляет статус транзакции по её ID.
func (h *Handler) changeTransactionStatus(c *gin.Context) {
	type updateTransactionStatusRequest struct {
		Status model.Status `json:"status"`
	}

	var request updateTransactionStatusRequest
	if err := c.BindJSON(&request); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, service.ErrIncorrectTransactionData)
		return
	}

	transactionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, ErrInvalidTransactionID)
		return
	}

	if err = h.service.TransactionService.ChangeStatus(transactionId, request.Status); err != nil {
		if err == service.ErrTerminalTransactionStatus {
			h.newErrorResponse(c, http.StatusBadRequest, err)
			return
		}
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// cancelTransaction отменяет транзакцию (платеж) по его ID.
func (h *Handler) cancelTransaction(c *gin.Context) {
	transactionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, ErrInvalidTransactionID)
		return
	}

	if err = h.service.TransactionService.Cancel(transactionId); err != nil {
		if err == service.ErrTerminalTransactionStatus {
			h.newErrorResponse(c, http.StatusBadRequest, err)
			return
		}
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
