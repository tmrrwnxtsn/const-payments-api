package store

import "errors"

// ErrTransactionNotFound возникает, когда искомая транзакция не найдена в базе данных.
var ErrTransactionNotFound = errors.New("transaction not found")
