package repositories

import (
	"database/sql"
	"time"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (repository *CustomerRepository) ExistsByIdentityNumberOrPhoneNumber(identityNumber, phoneNumber string) (bool, error) {
	var count int
	err := repository.DB.QueryRow("SELECT COUNT(*) FROM customers WHERE identity_number = $1 OR phone_number = $2", identityNumber, phoneNumber).Scan(&count)
	return count > 0, err
}

func (repository *CustomerRepository) CreateCustomer(name, identityNumber, phoneNumber string) (int, error) {
	var id int
	now := time.Now()
	err := repository.DB.QueryRow(`
        INSERT INTO customers (name, identity_number, phone_number, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $4) RETURNING id
    `, name, identityNumber, phoneNumber, now).Scan(&id)
	return id, err
}
