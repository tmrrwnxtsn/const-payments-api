package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
)

// Store представляет базу данных, в которой находятся таблицы сущностей "Пользователь" и "Транзакция".
type Store interface {
	Users() UsersRepository
	Transactions() TransactionsRepository
}

type store struct {
	db                     *sqlx.DB
	logger                 log.Logger
	usersRepository        UsersRepository
	transactionsRepository TransactionsRepository
}

func New(db *sqlx.DB, logger log.Logger) Store {
	logger.Debug("a connection to the database has been established")
	return &store{db: db, logger: logger}
}

func (s *store) Users() UsersRepository {
	if s.usersRepository != nil {
		return s.usersRepository
	}

	s.usersRepository = NewUsersRepository(s)

	return s.usersRepository
}

func (s *store) Transactions() TransactionsRepository {
	if s.transactionsRepository != nil {
		return s.transactionsRepository
	}

	s.transactionsRepository = NewTransactionsRepository(s)

	return s.transactionsRepository
}
