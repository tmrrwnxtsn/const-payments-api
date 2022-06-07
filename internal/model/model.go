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
	ID           uint64    `json:"id" db:"id" example:"1"`
	UserID       uint64    `json:"user_id" db:"user_id" example:"1"`
	UserEmail    string    `json:"user_email" db:"user_email" example:"tmrrwnxtsn@gmail.com"`
	Amount       float64   `json:"amount" db:"amount" example:"123.456"`
	CurrencyCode string    `json:"currency_code" db:"currency_code" example:"RUB"`
	CreationTime time.Time `json:"creation_time" db:"creation_time" example:"2022-06-07T15:25:16.046823Z"`
	ModifiedTime time.Time `json:"modified_time" db:"modified_time" example:"2022-06-07T15:25:16.046823Z"`
	Status       Status    `json:"status" db:"status" example:"НОВЫЙ,string"`
}

// Validate проверяет информацию в транзакции на корректность.
func (t *Transaction) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.UserID, validation.Required),
		validation.Field(&t.UserEmail, validation.Required, is.Email),
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
	// StatusSuccess представляет статус "УСПЕХ". Терминальный.
	StatusSuccess
	// StatusFailure представляет статус "НЕУСПЕХ". Терминальный.
	StatusFailure
	// StatusError представляет статус "ОШИБКА". Назначается при создании транзакции.
	StatusError
)
