package service

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Services представляет слой бизнес-логики.
type Services struct {
	TransactionService TransactionService
	logger             log.Logger
}

func NewServices(store store.Store, logger log.Logger) *Services {
	return &Services{
		TransactionService: NewTransactionService(store.Transactions()),
		logger:             logger,
	}
}
