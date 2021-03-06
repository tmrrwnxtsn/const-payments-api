package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"io"
	"net/http"
	"strconv"
)

// createTransactionRequest принимает ID пользователя, email пользователя, сумму и валюту платежа
type createTransactionRequest struct {
	UserID       uint64  `json:"user_id,string" binding:"required" example:"1"`
	UserEmail    string  `json:"user_email" binding:"required" example:"tmrrwnxtsn@gmail.com"`
	Amount       float64 `json:"amount,string" binding:"required" example:"123.456"`
	CurrencyCode string  `json:"currency_code" binding:"required" example:"RUB"`
}
type createTransactionResponse struct {
	ID     uint64       `json:"id" example:"1"`
	Status model.Status `json:"status" example:"УСПЕХ,string"`
}

// createTransaction godoc
// @Summary      Создать платёж (транзакцию)
// @Description  Платёжная система создаёт платёж (транзакцию) и уведомляет, прошёл ли он в системе.
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        input  body      createTransactionRequest   true  "Информация о транзакции"
// @Success      201    {object}  createTransactionResponse  "ok"
// @Failure      400    {object}  errorResponse              "Некорректные данные транзакции"
// @Failure      500    {object}  errorResponse              "Ошибка на стороне сервера"
// @Router       /transactions/ [post]
func (h *Handler) createTransaction(c *gin.Context) {
	var request createTransactionRequest
	if err := c.BindJSON(&request); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, service.ErrIncorrectTransactionData)
		return
	}

	// пользователь создаёт платёж (транзакцию) в статусе "НОВЫЙ"
	// при создании случайным образом этот статус может поменяться на "ОШИБКА"
	transactionID, createdStatus, err := h.service.TransactionService.Create(model.Transaction{
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

	redirectURL := fmt.Sprintf("http://%s/api/transactions/%d/status", c.Request.Host, transactionID)

	// платёжная система делает запрос на наш эндпоинт на изменение статуса платежа (транзакции)
	appliedStatus, err := callbackChangeTransactionStatus(redirectURL, createdStatus, c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, createTransactionResponse{
		ID:     transactionID,
		Status: appliedStatus,
	})
}

// callbackChangeTransactionStatus уведомляет о том, прошёл платёж или нет,
// отправляя запрос на наш эндпоинт на изменение статуса (PATCH /api/transactions/:id/status).
func callbackChangeTransactionStatus(redirectURL string, createdStatus model.Status, c *gin.Context) (model.Status, error) {
	var (
		changeStatusRequestBody []byte
		appliedStatus           model.Status
	)
	if createdStatus == model.StatusNew {
		// если платёж (транзакция) создался со статусом "НОВЫЙ", то подтверждаем создание платежа статусом "УСПЕХ"
		changeStatusRequestBody = []byte(`{"status":"УСПЕХ"}`)
		appliedStatus = model.StatusSuccess
	} else {
		// если платёж (транзакция) создался со статусом "ОШИБКА", то отрицаем создание платежа статусом "НЕУСПЕХ"
		changeStatusRequestBody = []byte(`{"status":"НЕУСПЕХ"}`)
		appliedStatus = model.StatusFailure
	}

	changeStatusRequest, err := http.NewRequestWithContext(
		c,
		"PATCH",
		redirectURL,
		bytes.NewBuffer(changeStatusRequestBody),
	)
	if err != nil {
		return createdStatus, err
	}
	// устанавливаем ID запроса такой же, чтобы можно было отследить, каким образом выполнялось создание и уведломление
	changeStatusRequest.Header.Set(requestIDKey, c.Writer.Header().Get(requestIDKey))

	changeStatusResponse, err := http.DefaultClient.Do(changeStatusRequest)
	if err != nil {
		return createdStatus, err
	}
	defer changeStatusResponse.Body.Close()

	if changeStatusResponse.Status != "200 OK" {
		bodyBytes, err := io.ReadAll(changeStatusResponse.Body)
		if err != nil {
			return createdStatus, err
		}
		return createdStatus, fmt.Errorf("failed to verify payment creation: %s", string(bodyBytes))
	}
	return appliedStatus, nil
}

type getAllUserTransactionsResponse struct {
	Data []model.Transaction `json:"data"`
}

// getAllUserTransactions godoc
// @Summary      Получить список всех платежей (транзакций) пользователя
// @Description  Необходимо передать либо ID, либо email пользователя, чтобы получить его платежи (транзакции).
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        user_email  query     string                          false  "Email пользователя"
// @Param        user_id     query     number                          false  "ID пользователя"
// @Success      200         {object}  getAllUserTransactionsResponse  "ok"
// @Failure      400         {object}  errorResponse                   "Некорректные данные пользователя"
// @Failure      500         {object}  errorResponse                   "Ошибка на стороне сервера"
// @Router       /transactions/ [get]
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

type getTransactionStatusResponse struct {
	Status model.Status `json:"status" example:"НОВЫЙ,string"`
}

// getTransactionStatus godoc
// @Summary  Возвращает статус платежа (транзакции) по его ID
// @Tags     transactions
// @Accept   json
// @Produce  json
// @Param    id   path      string                        true  "ID платежа (транзакции)"
// @Success  200  {object}  getTransactionStatusResponse  "ok"
// @Failure  400  {object}  errorResponse                 "Некорректный ID платежа (транзакции)"
// @Failure  500  {object}  errorResponse                 "Ошибка на стороне сервера"
// @Router   /transactions/{id}/status/ [get]
func (h *Handler) getTransactionStatus(c *gin.Context) {
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

type changeTransactionStatusRequest struct {
	Status model.Status `json:"status" example:"УСПЕХ,string"`
}

// changeTransactionStatus godoc
// @Summary      Изменяет статус платежа (транзакции) по его ID
// @Description  Статусы "УСПЕХ" и "НЕУСПЕХ" являются терминальными - если платеж находится в них, его статус невозможно поменять.
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id     path      string                          true  "ID платежа (транзакции)"
// @Param        input  body      changeTransactionStatusRequest  true  "Новый статус транзакции"
// @Success      200    {object}  statusResponse                  "ok"
// @Failure      400    {object}  errorResponse                   "Некорректный ID платежа (транзакции) или терминальный статус платежа (транзакции)"
// @Failure      500    {object}  errorResponse                   "Ошибка на стороне сервера"
// @Router       /transactions/{id}/status/ [patch]
func (h *Handler) changeTransactionStatus(c *gin.Context) {
	var request changeTransactionStatusRequest
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

// cancelTransaction godoc
// @Summary  Отменяет платёж (транзакцию) по его ID
// @Tags     transactions
// @Accept   json
// @Produce  json
// @Param    id   path      string          true  "ID платежа (транзакции)"
// @Success  200  {object}  statusResponse  "ok"
// @Failure  400  {object}  errorResponse   "Некорректный ID платежа (транзакции)"
// @Failure  500  {object}  errorResponse   "Ошибка на стороне сервера"
// @Router   /transactions/{id}/ [delete]
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
