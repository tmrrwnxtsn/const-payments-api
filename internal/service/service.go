package service

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Service представляет слой бизнес-логики, связанный с сущностями "Пользователь" и "Транзакция".
type Service interface {
	TransactionService
}

type service struct {
	TransactionService
	logger log.Logger
}

func NewService(store store.Store, logger log.Logger) Service {
	return &service{
		TransactionService: NewTransactionService(
			store.Users(),
			store.Transactions(),
		),
		logger: logger,
	}
}
