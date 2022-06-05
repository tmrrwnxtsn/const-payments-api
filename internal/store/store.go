package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Store представляет слой хранения данных (база данных).
type Store interface {
	// Transactions представляет таблицу с информацией о транзакциях.
	Transactions() TransactionRepository
}

type store struct {
	db                    *sqlx.DB
	logger                log.Logger
	transactionRepository TransactionRepository
}

func NewStore(db *sqlx.DB, logger log.Logger) Store {
	logger.Debug("a connection to the database has been established")
	return &store{db: db, logger: logger}
}

func (s *store) Transactions() TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = NewTransactionRepository(s)

	return s.transactionRepository
}
