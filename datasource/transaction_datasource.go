package datasource

import (
	"time"
)

type Transaction struct {
	Id                  int
	AccountNumber       string
	TransactionDate     time.Time
	Type                string
	FromToAccountNumber string
	InitialBalance      int64
	Amount              int64
	CreditOrDebit       string
}

var transactionHistories = []Transaction{}

type TransactionDatasources interface {
	GetTransactionHistoriesByAccountNumber(accountNumber string) (result []Transaction, err error)
	InsertTransactionHistory(transaction Transaction) (err error)
}

type transactionDatasource struct{}

func NewTransactionDatasource() *transactionDatasource {
	return &transactionDatasource{}
}

func (d transactionDatasource) GetTransactionHistoriesByAccountNumber(accountNumber string) (result []Transaction, err error) {
	for _, transaction := range transactionHistories {
		if transaction.AccountNumber == accountNumber {
			result = append(result, transaction)
		}
	}

	return
}

func (d transactionDatasource) InsertTransactionHistory(transaction Transaction) (err error) {
	id := 0
	if len(transactionHistories) > 0 {
		id = transactionHistories[len(transactionHistories)-1].Id + 1
	}

	transaction.Id = id
	transactionHistories = append(transactionHistories, transaction)

	return
}
