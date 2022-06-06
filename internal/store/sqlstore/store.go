package sqlstore

import (
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

var _ store.Store = (*Store)(nil)

type Store struct {
	db                    *sqlx.DB
	logger                log.Logger
	transactionRepository store.TransactionRepository
}

func NewStore(db *sqlx.DB, logger log.Logger) *Store {
	return &Store{db: db, logger: logger}
}

func (s *Store) Transactions() store.TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = NewTransactionRepository(s)

	return s.transactionRepository
}
