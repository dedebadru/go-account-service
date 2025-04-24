package services

import (
	"errors"

	"github.com/go-account-service/repositories"
	"github.com/go-account-service/utils"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type TransactionService struct {
	accountRepo     *repositories.AccountRepository
	transactionRepo *repositories.TransactionRepository
	logger          *logrus.Logger
}

func NewTransactionService(
	accountRepo *repositories.AccountRepository,
	transactionRepo *repositories.TransactionRepository,
	logger *logrus.Logger,
) *TransactionService {
	return &TransactionService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
		logger:          logger,
	}
}

func (service *TransactionService) Saving(accountNumber string, amount decimal.Decimal, transactionType string) (decimal.Decimal, error) {
	var newBalance decimal.Decimal
	balance := decimal.NewFromInt(0)
	account, err := service.accountRepo.GetAccountByAccountNumber(accountNumber)
	if err != nil {
		return balance, errors.New("rekening tidak ditemukan")
	}

	if transactionType == "CREDIT" {
		newBalance = utils.AddDecimal(account.Balance, amount)

	} else {
		if account.Balance.LessThan(amount) {
			return balance, errors.New("saldo tidak mencukupi")
		}

		newBalance = utils.SubtractDecimal(account.Balance, amount)
	}

	err = service.transactionRepo.CreateTransaction(account.CustomerID, account.ID, transactionType, amount)
	if err != nil {
		service.logger.Warn("Gagal mencatat transaksi: ", err)
	}

	err = service.accountRepo.UpdateSaldo(account.ID, newBalance)
	if err != nil {
		return balance, errors.New("gagal menabung")
	}

	account, _ = service.accountRepo.GetAccountByAccountNumber(accountNumber)
	return account.Balance, nil
}
