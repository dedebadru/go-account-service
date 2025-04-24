package repositories

import (
    "database/sql"
    "time"
    "github.com/shopspring/decimal"
)

type AccountRepository struct {
    DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
    return &AccountRepository{DB: db}
}

func (r *AccountRepository) CreateAccount(accountNumber string, customerID int) (int, error) {
    var id int
    now := time.Now()
    err := r.DB.QueryRow(`
        INSERT INTO accounts (account_number, customer_id, created_at, updated_at)
        VALUES ($1, $2, $3, $3) RETURNING id
    `, accountNumber, customerID, now).Scan(&id)
    
    return id, err
}

func (r *AccountRepository) GetBalance(accountNumber string) (decimal.Decimal, error) {
    var balance decimal.Decimal
    err := r.DB.QueryRow("SELECT balance FROM accounts WHERE account_number = $1", accountNumber).Scan(&balance)
    return balance, err
}


func (r *AccountRepository) UpdateSaldo(accountID int, newSaldo decimal.Decimal) error {
    _, err := r.DB.Exec("UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3",
        newSaldo, time.Now(), accountID)
    return err
}

func (r *AccountRepository) GetAccountAndCustomerID(accountNumber string) (accountID, customerID int, err error) {
    err = r.DB.QueryRow(`SELECT id, customer_id FROM accounts WHERE account_number = $1`, accountNumber).Scan(&accountID, &customerID)
    return 
}