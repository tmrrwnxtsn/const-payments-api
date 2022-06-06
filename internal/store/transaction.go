package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
)

const transactionTable = "transactions"

var ErrTransactionNotFound = errors.New("transaction not found")

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
}

type transactionRepository struct {
	store *store
}

func NewTransactionRepository(store *store) TransactionRepository {
	return &transactionRepository{store: store}
}

func (r *transactionRepository) Create(transaction model.Transaction) (uint64, error) {
	createTransactionQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, user_email, amount, currency_code, status) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		transactionTable,
	)

	var id uint64
	err := r.store.db.QueryRow(
		createTransactionQuery,
		transaction.UserID, transaction.UserEmail, transaction.Amount, transaction.CurrencyCode, transaction.Status,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *transactionRepository) GetAllByUserID(userID uint64) ([]model.Transaction, error) {
	getAllTransactionsQuery := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id = $1",
		transactionTable,
	)

	var transactions []model.Transaction
	err := r.store.db.Select(&transactions, getAllTransactionsQuery, userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) GetAllByUserEmail(userEmail string) ([]model.Transaction, error) {
	getAllTransactionsQuery := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_email = $1",
		transactionTable,
	)

	var transactions []model.Transaction
	err := r.store.db.Select(&transactions, getAllTransactionsQuery, userEmail)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) GetByID(transactionID uint64) (model.Transaction, error) {
	getTransactionByIdQuery := fmt.Sprintf(
		"SELECT * FROM %s WHERE id = $1",
		transactionTable)

	var transaction model.Transaction
	err := r.store.db.Get(&transaction, getTransactionByIdQuery, transactionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Transaction{}, ErrTransactionNotFound
		}
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (r *transactionRepository) ChangeStatus(transactionID uint64, status model.Status) error {
	updateTransactionQuery := fmt.Sprintf(
		"UPDATE %s SET status = $1 WHERE id = $2",
		transactionTable,
	)

	_, err := r.store.db.Exec(updateTransactionQuery, status, transactionID)
	return err
}
