package service

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	"math/rand"
	"time"
)

// TransactionService представляет бизнес-логику работы с транзакциями.
type TransactionService interface {
	// Create создаёт транзакцию.
	Create(transaction model.Transaction) (uint64, model.Status, error)
	// GetAllByUserID возвращает информацию о транзакциях пользователя по его ID.
	GetAllByUserID(userID uint64) ([]model.Transaction, error)
	// GetAllByUserEmail возвращает информацию о транзакциях пользователя по его email.
	GetAllByUserEmail(userEmail string) ([]model.Transaction, error)
	// GetStatus возвращает статус транзакции по её ID.
	GetStatus(transactionID uint64) (model.Status, error)
	// ChangeStatus изменяет статус транзакции по её ID.
	ChangeStatus(transactionID uint64, status model.Status) error
	// Cancel отменяет транзакцию по её ID.
	Cancel(transactionID uint64) error
}

type transactionService struct {
	transactionRepository store.TransactionRepository
}

func NewTransactionService(transactionRepository store.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: transactionRepository}
}

func (s *transactionService) Create(transaction model.Transaction) (uint64, model.Status, error) {
	if err := transaction.Validate(); err != nil {
		return 0, 0, ErrIncorrectTransactionData
	}

	// случайное количество платежей при создании переходит в статус "ОШИБКА"
	rand.Seed(time.Now().UnixNano())
	if n := rand.Int(); n%6 == 0 {
		transaction.Status = model.StatusError
	}

	transactionID, err := s.transactionRepository.Create(transaction)
	if err != nil {
		return 0, 0, err
	}
	return transactionID, transaction.Status, nil
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

func (s *transactionService) Cancel(transactionID uint64) error {
	transaction, err := s.transactionRepository.GetByID(transactionID)
	if err != nil {
		return err
	}

	// статусы УСПЕХ и НЕУСПЕХ являются терминальными: если платеж находится в них, его статус невозможно отменить
	if transaction.Status == model.StatusSuccess || transaction.Status == model.StatusFailure {
		return ErrTerminalTransactionStatus
	}

	return s.transactionRepository.Delete(transactionID)
}
