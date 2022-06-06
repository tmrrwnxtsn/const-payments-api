package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	mockservice "github.com/tmrrwnxtsn/const-payments-api/internal/service/mocks"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestHandler_createTransaction(t *testing.T) {
	type mockBehavior func(s *mockservice.MockTransactionService, transaction model.Transaction)

	logger := log.New()

	tests := []struct {
		name                 string
		inputBody            string
		inputTransaction     model.Transaction
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"user_id":"11","user_email":"tmrrwnxtsn@gmail.com","amount":"123.456","currency_code":"RUB"}`,
			inputTransaction: model.Transaction{
				UserID:       11,
				UserEmail:    "tmrrwnxtsn@gmail.com",
				Amount:       123.456,
				CurrencyCode: "RUB",
			},
			mockBehavior: func(s *mockservice.MockTransactionService, transaction model.Transaction) {
				s.EXPECT().Create(transaction).Return(uint64(1), nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "empty fields",
			inputBody:            `{"user_id":"11","amount":"123.456"}`,
			mockBehavior:         func(s *mockservice.MockTransactionService, transaction model.Transaction) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"transaction data is incorrect"}`,
		},
		{
			name:      "invalid fields",
			inputBody: `{"user_id":"11","user_email":"tmrrwnxtsn@gmail.com","amount":"-654.321","currency_code":"ruble"}`,
			inputTransaction: model.Transaction{
				UserID:       11,
				UserEmail:    "tmrrwnxtsn@gmail.com",
				Amount:       -654.321,
				CurrencyCode: "ruble",
			},
			mockBehavior: func(s *mockservice.MockTransactionService, transaction model.Transaction) {
				s.EXPECT().Create(transaction).Return(uint64(0), service.ErrIncorrectTransactionData)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"transaction data is incorrect"}`,
		},
		{
			name:      "service failure",
			inputBody: `{"user_id":"11","user_email":"tmrrwnxtsn@gmail.com","amount":"123.456","currency_code":"RUB"}`,
			inputTransaction: model.Transaction{
				UserID:       11,
				UserEmail:    "tmrrwnxtsn@gmail.com",
				Amount:       123.456,
				CurrencyCode: "RUB",
			},
			mockBehavior: func(s *mockservice.MockTransactionService, transaction model.Transaction) {
				s.EXPECT().Create(transaction).Return(uint64(0), errors.New("service failure"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"service failure"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockTransactionService := mockservice.NewMockTransactionService(c)
			tt.mockBehavior(mockTransactionService, tt.inputTransaction)

			services := &service.Services{TransactionService: mockTransactionService}
			handler := NewHandler(services, logger)

			router := gin.New()
			router.POST("/api/transactions/", handler.createTransaction)

			responseRecorder := httptest.NewRecorder()
			request := httptest.NewRequest(
				"POST",
				"/api/transactions/",
				bytes.NewBufferString(tt.inputBody),
			)

			router.ServeHTTP(responseRecorder, request)

			assert.Equal(t, tt.expectedStatusCode, responseRecorder.Code)
			assert.Equal(t, tt.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}

func TestHandler_getAllUserTransactions(t *testing.T) {
	type mockBehavior func(s *mockservice.MockTransactionService, queryParamValue string)

	testData := []model.Transaction{
		{
			ID:           3,
			UserID:       13,
			UserEmail:    "tmrrwnxtsn@gmail.com",
			Amount:       255,
			CurrencyCode: "USD",
			CreationTime: time.Date(2022, 6, 5, 12, 12, 12, 12, time.UTC),
			ModifiedTime: time.Date(2022, 6, 5, 12, 12, 12, 12, time.UTC),
			Status:       model.StatusNew,
		},
		{
			ID:           4,
			UserID:       13,
			UserEmail:    "tmrrwnxtsn@gmail.com",
			Amount:       123,
			CurrencyCode: "USD",
			CreationTime: time.Date(2022, 6, 5, 12, 12, 12, 12, time.UTC),
			ModifiedTime: time.Date(2022, 6, 5, 12, 12, 12, 12, time.UTC),
			Status:       model.StatusError,
		},
	}

	logger := log.New()

	tests := []struct {
		name                 string
		queryParamName       string
		queryParamValue      string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:            "ok with user_id",
			queryParamName:  "user_id",
			queryParamValue: "13",
			mockBehavior: func(s *mockservice.MockTransactionService, queryParamValue string) {
				userID, _ := strconv.ParseUint(queryParamValue, 10, 64)
				s.EXPECT().GetAllByUserID(userID).Return(testData, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"data":[{"id":3,"user_id":13,"user_email":"tmrrwnxtsn@gmail.com","amount":255,"currency_code":"USD","creation_time":"2022-06-05T12:12:12.000000012Z","modified_time":"2022-06-05T12:12:12.000000012Z","status":0},{"id":4,"user_id":13,"user_email":"tmrrwnxtsn@gmail.com","amount":123,"currency_code":"USD","creation_time":"2022-06-05T12:12:12.000000012Z","modified_time":"2022-06-05T12:12:12.000000012Z","status":3}]}`,
		},
		{
			name:            "ok with user_email",
			queryParamName:  "user_email",
			queryParamValue: "tmrrwnxtsn@gmail.com",
			mockBehavior: func(s *mockservice.MockTransactionService, queryParamValue string) {
				s.EXPECT().GetAllByUserEmail(queryParamValue).Return(testData, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"data":[{"id":3,"user_id":13,"user_email":"tmrrwnxtsn@gmail.com","amount":255,"currency_code":"USD","creation_time":"2022-06-05T12:12:12.000000012Z","modified_time":"2022-06-05T12:12:12.000000012Z","status":0},{"id":4,"user_id":13,"user_email":"tmrrwnxtsn@gmail.com","amount":123,"currency_code":"USD","creation_time":"2022-06-05T12:12:12.000000012Z","modified_time":"2022-06-05T12:12:12.000000012Z","status":3}]}`,
		},
		{
			name:                 "ok without params",
			mockBehavior:         func(s *mockservice.MockTransactionService, queryParamValue string) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"data":null}`,
		},
		{
			name:                 "invalid user_id",
			queryParamName:       "user_id",
			queryParamValue:      "-ABS",
			mockBehavior:         func(s *mockservice.MockTransactionService, queryParamValue string) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"incorrect query parameters data"}`,
		},
		{
			name:            "service failure with user_id",
			queryParamName:  "user_id",
			queryParamValue: "13",
			mockBehavior: func(s *mockservice.MockTransactionService, queryParamValue string) {
				userID, _ := strconv.ParseUint(queryParamValue, 10, 64)
				s.EXPECT().GetAllByUserID(userID).Return(nil, errors.New("service failure"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"service failure"}`,
		},
		{
			name:            "service failure with user_email",
			queryParamName:  "user_email",
			queryParamValue: "tmrrwnxtsn@gmail.com",
			mockBehavior: func(s *mockservice.MockTransactionService, queryParamValue string) {
				s.EXPECT().GetAllByUserEmail(queryParamValue).Return(nil, errors.New("service failure"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"service failure"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockTransactionService := mockservice.NewMockTransactionService(c)
			tt.mockBehavior(mockTransactionService, tt.queryParamValue)

			services := &service.Services{TransactionService: mockTransactionService}
			handler := NewHandler(services, logger)

			router := gin.New()
			router.GET("/api/transactions/", handler.getAllUserTransactions)

			responseRecorder := httptest.NewRecorder()
			request := httptest.NewRequest(
				"GET",
				fmt.Sprintf("/api/transactions/?%s=%s", tt.queryParamName, tt.queryParamValue),
				nil,
			)

			router.ServeHTTP(responseRecorder, request)

			assert.Equal(t, tt.expectedStatusCode, responseRecorder.Code)
			assert.Equal(t, tt.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
