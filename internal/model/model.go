package model

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"strings"
	"time"
)

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
func (t *Transaction) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Amount, validation.Required, validation.Min(0.0)),
		validation.Field(&t.CurrencyCode, validation.Required, is.CurrencyCode),
	)
}

// Status представляет статус транзакции.
type Status uint32

// String конвертирует Status в строковое представление. Например, StatusNew преобразуется в "НОВЫЙ".
func (s Status) String() string {
	if b, err := s.MarshalText(); err == nil {
		return string(b)
	} else {
		return "НЕИЗВЕСТНО"
	}
}

// ParseStatus принимает статус в строковой форме и возвращает соответствующую константу типа Status.
func ParseStatus(status string) (Status, error) {
	switch strings.ToLower(status) {
	case "новый":
		return StatusNew, nil
	case "успех":
		return StatusSuccess, nil
	case "неуспех":
		return StatusFailure, nil
	case "ошибка":
		return StatusError, nil
	}
	var l Status
	return l, fmt.Errorf("not a valid transaction status: %q", status)
}

// UnmarshalText реализует encoding.TextUnmarshaler.
func (s *Status) UnmarshalText(text []byte) error {
	status, err := ParseStatus(string(text))
	if err != nil {
		return err
	}
	*s = status
	return nil
}

// MarshalText реализует encoding.TextMarshaler.
func (s Status) MarshalText() ([]byte, error) {
	switch s {
	case StatusNew:
		return []byte("НОВЫЙ"), nil
	case StatusSuccess:
		return []byte("УСПЕХ"), nil
	case StatusFailure:
		return []byte("НЕУСПЕХ"), nil
	case StatusError:
		return []byte("ОШИБКА"), nil
	}
	return nil, fmt.Errorf("not a valid transaction status %d", s)
}

const (
	// StatusNew представляет статус "НОВЫЙ". Назначается при создании транзакции.
	StatusNew Status = iota
	// StatusSuccess представляет статус "УСПЕХ".
	StatusSuccess
	// StatusFailure представляет статус "НЕУСПЕХ".
	StatusFailure
	// StatusError представляет статус "ОШИБКА". Назначается при создании транзакции.
	StatusError
)
