package model

import (
	"testing"
	"time"
)

// TestTransaction создаёт экземпляр валидной транзакции для тестов.
func TestTransaction(t *testing.T) *Transaction {
	t.Helper()

	return &Transaction{
		UserID:       10,
		UserEmail:    "tmrrwnxtsn@gmail.com",
		Amount:       255.24,
		CurrencyCode: "RUB",
		CreationTime: time.Now(),
		ModifiedTime: time.Now(),
		Status:       StatusNew,
	}
}
