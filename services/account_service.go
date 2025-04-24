package services

import (
	"errors"

	"github.com/go-account-service/repositories"
	"github.com/go-account-service/utils"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	accountRepo *repositories.AccountRepository
	logger      *logrus.Logger
}

func NewAccountService(
	accountRepo *repositories.AccountRepository,
	logger *logrus.Logger,
) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		logger:      logger,
	}
}

func (service *AccountService) ShowBalance(accountNumber string) (decimal.Decimal, error) {
	account, error := service.accountRepo.GetAccountByAccountNumber(accountNumber)
	if error != nil {
		message := "rekening tidak ditemukan"
		service.logger.WithError(error).Error(message)
		return decimal.NewFromInt(0), errors.New(message)
	}

	return utils.FormatDecimal(account.Balance, 2), nil
}
