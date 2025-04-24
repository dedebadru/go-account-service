package repositories

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (repository *TransactionRepository) CreateTransaction(customerID, accountID int, transactionType string, amount decimal.Decimal) error {
	_, err := repository.DB.Exec(`
        INSERT INTO transactions (customer_id, account_id, type, amount, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $5)
    `, customerID, accountID, transactionType, amount, time.Now())
	return err
}
