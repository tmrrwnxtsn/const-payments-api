package service

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
)

// TransactionService представляет бизнес-логику, связанную с сущностью "Транзакция".
type TransactionService interface {
	Create(transaction model.Transaction) (uint64, error)
}

type transactionService struct {
	userRepository        store.UserRepository
	transactionRepository store.TransactionRepository
}

func NewTransactionService(userRepository store.UserRepository, transactionRepository store.TransactionRepository) TransactionService {
	return &transactionService{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

func (s *transactionService) Create(transaction model.Transaction) (uint64, error) {
	return s.transactionRepository.Create(transaction)
}
