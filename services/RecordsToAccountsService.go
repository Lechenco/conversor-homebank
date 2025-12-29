package services

import (
	"strconv"
	"time"

	"github.com/Lechenco/conversor-homebank/models"
)

const (
	COLUMN_ACCOUNT     = 2
	COLUMN_TRANSACTION = 0
	COLUMN_CATEGORY    = 1
	DATE_FORMAT        = "2006/01/02"
)

func RecordsToAccounts(records [][]string) []models.Account {
	columnAccount := COLUMN_ACCOUNT
	accounts := []models.Account{}

	for columnAccount < len(records[0]) {
		acc := models.Account{
			Name:         records[0][columnAccount],
			Transactions: recordsToTransactions(records, columnAccount),
		}

		accounts = append(accounts, acc)
		columnAccount++
	}

	return accounts
}

func recordsToTransactions(records [][]string, columnAccount int) []*models.Transaction {
	transactions := []*models.Transaction{}

	for i := 1; i < len(records); i++ {
		transaction := recordToTransaction(records[i], columnAccount)

		transactions = append(transactions, &transaction)
	}

	return transactions
}

func recordToTransaction(record []string, columnAccount int) models.Transaction {
	d := time.Now().Format(DATE_FORMAT)
	isTransfer := record[COLUMN_TRANSACTION] == "1"
	category := record[COLUMN_CATEGORY]
	value, _ := strconv.ParseFloat(record[columnAccount], 32)

	transaction := models.Transaction{
		Date:   d,
		Value:  float32(value),
		Status: models.None,
	}

	if isTransfer {
		transaction.Transfer = category
	} else {
		transaction.Category = category
	}

	return transaction
}
