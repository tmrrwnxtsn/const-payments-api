package model

import "time"

// User представляет сущность "Пользователь".
type User struct {
	ID    uint64
	Email string
}

// Status представляет статус платежа (НОВЫЙ, УСПЕХ, НЕУСПЕХ, ОШИБКА).
type Status int

const (
	New     Status = iota // НОВЫЙ
	Success               // УСПЕХ
	Failure               // НЕУСПЕХ
	Error                 // ОШИБКА
)

func (s Status) String() string {
	return [...]string{"НОВЫЙ", "УСПЕХ", "НЕУСПЕХ", "ОШИБКА"}[s]
}

// Transaction представляет сущность "Транзакция".
type Transaction struct {
	ID           uint64
	UserID       uint64
	Amount       uint64
	CurrencyCode string
	CreationTime time.Time
	ModifiedTime time.Time
	Status       Status
}
