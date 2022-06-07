package service

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
)

// Services представляет слой бизнес-логики.
type Services struct {
	TransactionService TransactionService
}

func NewServices(store store.Store) *Services {
	return &Services{
		TransactionService: NewTransactionService(store.Transactions()),
	}
}
