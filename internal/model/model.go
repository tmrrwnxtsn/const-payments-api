package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

// Status представляет статус транзакции (НОВЫЙ, УСПЕХ, НЕУСПЕХ, ОШИБКА).
type Status uint32

const (
	New     Status = iota // НОВЫЙ
	Success               // УСПЕХ
	Failure               // НЕУСПЕХ
	Error                 // ОШИБКА
)

func (s Status) String() string {
	return [...]string{"НОВЫЙ", "УСПЕХ", "НЕУСПЕХ", "ОШИБКА"}[s]
}

// Transaction представляет транзакцию.
type Transaction struct {
	ID           uint64    `json:"id" db:"id"`
	UserID       uint64    `json:"user_id" db:"user_id"`
	UserEmail    string    `json:"user_email" db:"user_email"`
	Amount       float64   `json:"amount" db:"amount"`
	CurrencyCode string    `json:"currency_code" db:"currency_code"`
	CreationTime time.Time `json:"creation_time" db:"creation_time"`
	ModifiedTime time.Time `json:"modified_time" db:"modified_time"`
	Status       Status    `json:"status" db:"status"`
}

// Validate проверяет информацию в транзакции на корректность.
func (t Transaction) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Amount, validation.Required, validation.Min(0.0)),
		validation.Field(&t.CurrencyCode, validation.Required, is.CurrencyCode),
	)
}
