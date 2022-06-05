package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Store представляет базу данных, в которой находятся таблицы сущностей "Пользователь" и "Транзакция".
type Store interface {
	Users() UserRepository
	Transactions() TransactionRepository
}

type store struct {
	db                    *sqlx.DB
	logger                log.Logger
	userRepository        UserRepository
	transactionRepository TransactionRepository
}

func NewStore(db *sqlx.DB, logger log.Logger) Store {
	logger.Debug("a connection to the database has been established")
	return &store{db: db, logger: logger}
}

func (s *store) Users() UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = NewUserRepository(s)

	return s.userRepository
}

func (s *store) Transactions() TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = NewTransactionRepository(s)

	return s.transactionRepository
}
