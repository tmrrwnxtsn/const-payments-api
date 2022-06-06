package store

import "github.com/tmrrwnxtsn/const-payments-api/internal/model"

// TransactionRepository представляет таблицу транзакций в базе данных.
type TransactionRepository interface {
	// Create создаёт транзакцию.
	Create(transaction model.Transaction) (uint64, error)
	// GetAllByUserID возвращает список транзакций пользователя по его ID.
	GetAllByUserID(userID uint64) ([]model.Transaction, error)
	// GetAllByUserEmail возвращает список транзакций пользователя по его ID.
	GetAllByUserEmail(userEmail string) ([]model.Transaction, error)
	// GetByID возвращает транзакцию по её ID.
	GetByID(transactionID uint64) (model.Transaction, error)
	// ChangeStatus изменяет статус транзакции по её ID.
	ChangeStatus(transactionID uint64, status model.Status) error
	// Delete удаляет транзакцию из БД по её ID.
	Delete(transactionID uint64) error
}
