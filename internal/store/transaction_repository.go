package store

import "github.com/tmrrwnxtsn/const-payments-api/internal/model"

const transactionTable = "transactions"

// TransactionRepository представляет таблицу сущности "Транзакция" в базе данных.
type TransactionRepository interface {
	// Create создаёт транзакцию в базе данных.
	Create(transaction model.Transaction) (uint64, error)
}

type transactionRepository struct {
	store Store
}

func NewTransactionRepository(store Store) TransactionRepository {
	return &transactionRepository{store: store}
}

func (r *transactionRepository) Create(transaction model.Transaction) (uint64, error) {
	panic("implement me")
}
