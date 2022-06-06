package handler

import "errors"

var (
	// ErrInvalidQueryParams возникает при обработке некорректных параметров запроса.
	ErrInvalidQueryParams = errors.New("invalid query parameters data")
	// ErrInvalidTransactionID возникает при обработке некорректного ID транзакции.
	ErrInvalidTransactionID = errors.New("invalid transaction id")
)
