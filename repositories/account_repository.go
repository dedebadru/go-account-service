package repositories

import (
	"database/sql"
	"time"

	"github.com/go-account-service/models"
	"github.com/shopspring/decimal"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (repository *AccountRepository) CreateAccount(accountNumber string, customerID int) (int, error) {
	var id int
	now := time.Now()
	err := repository.DB.QueryRow(`
        INSERT INTO accounts (account_number, customer_id, created_at, updated_at)
        VALUES ($1, $2, $3, $3) RETURNING id
    `, accountNumber, customerID, now).Scan(&id)

	return id, err
}

func (repository *AccountRepository) UpdateSaldo(accountID int, newSaldo decimal.Decimal) error {
	_, err := repository.DB.Exec("UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3",
		newSaldo, time.Now(), accountID)
	return err
}

func (repository *AccountRepository) GetAccountByAccountNumber(accountNumber string) (*models.Account, error) {
	account := &models.Account{}
	err := repository.DB.QueryRow(`SELECT id, customer_id, account_number, balance, created_at, updated_at FROM accounts WHERE account_number = $1`, accountNumber).Scan(&account.ID, &account.CustomerID, &account.AccountNumber, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}
