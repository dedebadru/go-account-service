package models

import (
    "time"
	"github.com/shopspring/decimal"
)

type Account struct {
    ID              int             `json:"id"`
    CustomerID      int             `json:"customer_id"`
    AccountNumber   string          `json:"account_number"`
    Balance         decimal.Decimal `json:"balance"`
    CreatedAt       time.Time       `json:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at"`
}
