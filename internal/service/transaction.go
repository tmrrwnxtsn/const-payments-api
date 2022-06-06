//go:generate mockgen -source=transaction.go -destination=mocks/transaction_mock.go
package service

import (
	"errors"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	"math/rand"
	"time"
)

var (
	ErrIncorrectTransactionData  = errors.New("transaction data is incorrect")
	ErrTerminalTransactionStatus = errors.New("terminal transaction status")
)

// TransactionService представляет бизнес-логику работы с транзакциями.
type TransactionService interface {
	// Create создаёт транзакцию.
	Create(transaction model.Transaction) (uint64, error)
	// GetAllByUserID возвращает информацию о транзакциях пользователя по его ID.
	GetAllByUserID(userID uint64) ([]model.Transaction, error)
	// GetAllByUserEmail возвращает информацию о транзакциях пользователя по его email.
	GetAllByUserEmail(userEmail string) ([]model.Transaction, error)
	// GetStatus возвращает статус транзакции по её id.
	GetStatus(transactionID uint64) (model.Status, error)
	// ChangeStatus изменяет статус транзакции по её id.
	ChangeStatus(transactionID uint64, status model.Status) error
}

type transactionService struct {
	transactionRepository store.TransactionRepository
}

func NewTransactionService(transactionRepository store.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: transactionRepository}
}

func (s *transactionService) Create(transaction model.Transaction) (uint64, error) {
	if err := transaction.Validate(); err != nil {
		return 0, ErrIncorrectTransactionData
	}

	// случайное количество платежей при создании переходит в статус "ОШИБКА"
	rand.Seed(time.Now().UnixNano())
	if n := rand.Int(); n%8 == 0 {
		transaction.Status = model.StatusError
	}

	return s.transactionRepository.Create(transaction)
}

func (s *transactionService) GetAllByUserID(userID uint64) ([]model.Transaction, error) {
	return s.transactionRepository.GetAllByUserID(userID)
}

func (s *transactionService) GetAllByUserEmail(userEmail string) ([]model.Transaction, error) {
	return s.transactionRepository.GetAllByUserEmail(userEmail)
}

func (s *transactionService) GetStatus(transactionID uint64) (model.Status, error) {
	transaction, err := s.transactionRepository.GetByID(transactionID)
	if err != nil {
		return 0, err
	}
	return transaction.Status, nil
}

func (s *transactionService) ChangeStatus(transactionID uint64, status model.Status) error {
	transaction, err := s.transactionRepository.GetByID(transactionID)
	if err != nil {
		return err
	}

	// статусы УСПЕХ и НЕУСПЕХ являются терминальными: если платеж находится в них, его статус невозможно изменить
	if transaction.Status == model.StatusSuccess || transaction.Status == model.StatusFailure {
		return ErrTerminalTransactionStatus
	}

	return s.transactionRepository.ChangeStatus(transaction.ID, status)
}
