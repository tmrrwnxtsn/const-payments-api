package store

import "github.com/tmrrwnxtsn/const-payments-api/internal/model"

const transactionsTable = "transactions"

// TransactionsRepository представляет таблицу сущности "Транзакция" в базе данных.
type TransactionsRepository interface {
	// Create создаёт транзакцию в базе данных.
	Create(transaction model.Transaction) (uint64, error)
}

type transactionsRepository struct {
	store Store
}

func NewTransactionsRepository(store Store) TransactionsRepository {
	return &transactionsRepository{store: store}
}

func (r *transactionsRepository) Create(transaction model.Transaction) (uint64, error) {
	panic("implement me")
}
