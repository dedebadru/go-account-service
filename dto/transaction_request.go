package dto

import (
	"github.com/shopspring/decimal"
)

type TransactionRequest struct {
	AccountNumber string  				`json:"no_rekening" validate:"required"`
	Amount  			decimal.Decimal `json:"nominal" validate:"required,gt=0"`
}
