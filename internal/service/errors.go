package service

import "errors"

var (
	// ErrIncorrectTransactionData возникает при попытке занести в базу данных невалидные данные транзакции.
	ErrIncorrectTransactionData = errors.New("transaction data is incorrect")
	// ErrTerminalTransactionStatus возникает при попытке изменения терминального статуса транзакции (model.StatusSuccess, model.StatusFailure).
	ErrTerminalTransactionStatus = errors.New("terminal transaction status")
)
