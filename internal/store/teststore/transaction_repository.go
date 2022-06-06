package teststore

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
)

var _ store.TransactionRepository = (*TransactionRepository)(nil)

type TransactionRepository struct {
	store        *Store
	transactions map[uint64]model.Transaction
}

func (r *TransactionRepository) Create(transaction model.Transaction) (uint64, error) {
	if err := transaction.Validate(); err != nil {
		return 0, err
	}

	transaction.ID = uint64(len(r.transactions) + 1)
	r.transactions[transaction.ID] = transaction

	return transaction.ID, nil
}

func (r *TransactionRepository) GetAllByUserID(userID uint64) ([]model.Transaction, error) {
	var transactions []model.Transaction

	for id, transaction := range r.transactions {
		if transaction.UserID == userID {
			transaction.ID = id
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

func (r *TransactionRepository) GetAllByUserEmail(userEmail string) ([]model.Transaction, error) {
	var transactions []model.Transaction

	for id, transaction := range r.transactions {
		if transaction.UserEmail == userEmail {
			transaction.ID = id
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

func (r *TransactionRepository) GetByID(transactionID uint64) (model.Transaction, error) {
	transaction, ok := r.transactions[transactionID]
	if !ok {
		return model.Transaction{}, store.ErrTransactionNotFound
	}

	return transaction, nil
}

func (r *TransactionRepository) ChangeStatus(transactionID uint64, status model.Status) error {
	transaction, ok := r.transactions[transactionID]
	if !ok {
		return store.ErrTransactionNotFound
	}

	transaction.Status = status
	r.transactions[transactionID] = transaction
	return nil
}

func (r *TransactionRepository) Delete(transactionID uint64) error {
	delete(r.transactions, transactionID)
	return nil
}
