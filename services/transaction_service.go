package services

import (
	"errors"
	"github.com/go-account-service/repositories"
	"github.com/go-account-service/utils"
	"github.com/sirupsen/logrus"
  "github.com/shopspring/decimal"
)

type TransactionService struct {
	accountRepo		*repositories.AccountRepository
	transactionRepo	*repositories.TransactionRepository
	logger        *logrus.Logger
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

func (s *TransactionService) Saving(accountNumber string, amount decimal.Decimal, transactionType string) (decimal.Decimal, error) {
	var newBalance decimal.Decimal
	balance := decimal.NewFromInt(0)
	customerID, accountID, err := s.accountRepo.GetAccountAndCustomerID(accountNumber)
	if err != nil {
		return balance, errors.New("rekening tidak ditemukan")
	}

	prevBalance, err := s.accountRepo.GetBalance(accountNumber)
	if err != nil {
		return balance, errors.New("gagal mendapatkan saldo")
	}

	if transactionType == "CREDIT"{
		newBalance = utils.AddDecimal(prevBalance, amount)
	
	} else {	
		if prevBalance.LessThan(amount) {
			return balance, errors.New("saldo tidak mencukupi")
		}

		newBalance = utils.SubtractDecimal(prevBalance, amount)
	}

	err = s.transactionRepo.CreateTransaction(customerID, accountID, transactionType, amount)
	if err != nil {
		s.logger.Warn("Gagal mencatat transaksi: ", err)
	}

	err = s.accountRepo.UpdateSaldo(accountID, newBalance)
	if err != nil {
		return balance, errors.New("gagal menabung")
	}

	balance, _ = s.accountRepo.GetBalance(accountNumber)
	return balance, nil
}

