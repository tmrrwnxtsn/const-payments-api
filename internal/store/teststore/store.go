package teststore

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
)

var _ store.Store = (*Store)(nil)

type Store struct {
	transactionRepository store.TransactionRepository
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Transactions() store.TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store:        s,
		transactions: make(map[uint64]model.Transaction),
	}

	return s.transactionRepository
}
