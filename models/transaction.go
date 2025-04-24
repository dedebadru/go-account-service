package models

import (
    "time"
	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"
)

type Transaction struct {
    ID         int             `json:"id"`
    CustomerID int             `json:"customer_id"`
    AccountID  int             `json:"account_id"`
    Type       TransactionType `json:"type"`
    Amount     decimal.Decimal `json:"amount"`
    CreatedAt  time.Time       `json:"created_at"`
    UpdatedAt  time.Time       `json:"updated_at"`
}