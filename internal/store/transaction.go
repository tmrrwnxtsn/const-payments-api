package store

import (
	"fmt"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
)

const transactionTable = "transactions"

// TransactionRepository представляет таблицу транзакций в базе данных.
type TransactionRepository interface {
	// Create создаёт транзакцию.
	Create(transaction model.Transaction) (uint64, error)
	// GetAllByUserID возвращает список транзакций пользователя по его ID.
	GetAllByUserID(userID uint64) ([]model.Transaction, error)
	// GetAllByUserEmail возвращает список транзакций пользователя по его ID.
	GetAllByUserEmail(userEmail string) ([]model.Transaction, error)
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
