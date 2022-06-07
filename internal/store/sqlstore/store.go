package sqlstore

import (
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
)

var _ store.Store = (*Store)(nil)

type Store struct {
	db                    *sqlx.DB
	transactionRepository store.TransactionRepository
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Transactions() store.TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = NewTransactionRepository(s)

	return s.transactionRepository
}
