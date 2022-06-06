package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
)

const transactionTable = "transactions"

var _ store.TransactionRepository = (*TransactionRepository)(nil)

type TransactionRepository struct {
	store *Store
}

func NewTransactionRepository(store *Store) *TransactionRepository {
	return &TransactionRepository{store: store}
}

func (r *TransactionRepository) Create(transaction model.Transaction) (uint64, error) {
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

func (r *TransactionRepository) GetAllByUserID(userID uint64) ([]model.Transaction, error) {
	getAllTransactionsQuery := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id = $1",
		transactionTable,
	)

	var transactions []model.Transaction
	if err := r.store.db.Select(&transactions, getAllTransactionsQuery, userID); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) GetAllByUserEmail(userEmail string) ([]model.Transaction, error) {
	getAllTransactionsQuery := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_email = $1",
		transactionTable,
	)

	var transactions []model.Transaction
	if err := r.store.db.Select(&transactions, getAllTransactionsQuery, userEmail); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) GetByID(transactionID uint64) (model.Transaction, error) {
	getTransactionByIdQuery := fmt.Sprintf(
		"SELECT * FROM %s WHERE id = $1",
		transactionTable)

	var transaction model.Transaction
	if err := r.store.db.Get(&transaction, getTransactionByIdQuery, transactionID); err != nil {
		if err == sql.ErrNoRows {
			return model.Transaction{}, store.ErrTransactionNotFound
		}
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepository) ChangeStatus(transactionID uint64, status model.Status) error {
	updateTransactionQuery := fmt.Sprintf(
		"UPDATE %s SET status = $1 WHERE id = $2",
		transactionTable,
	)

	_, err := r.store.db.Exec(updateTransactionQuery, status, transactionID)
	return err
}

func (r *TransactionRepository) Delete(transactionID uint64) error {
	deleteTransactionQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1",
		transactionTable,
	)

	_, err := r.store.db.Exec(deleteTransactionQuery, transactionID)
	return err
}
