package store

// Store представляет слой хранения данных (база данных).
type Store interface {
	// Transactions представляет таблицу с информацией о транзакциях.
	Transactions() TransactionRepository
}
