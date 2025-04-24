package services

import (
	"errors"

	"github.com/go-account-service/dto"
	"github.com/go-account-service/repositories"
	"github.com/go-account-service/utils"
	"github.com/sirupsen/logrus"
)

type RegistrationService struct {
	customerRepo *repositories.CustomerRepository
	accountRepo  *repositories.AccountRepository
	logger       *logrus.Logger
}

func NewRegistrationService(
	customerRepo *repositories.CustomerRepository,
	accountRepo *repositories.AccountRepository,
	logger *logrus.Logger,
) *RegistrationService {
	return &RegistrationService{
		customerRepo: customerRepo,
		accountRepo:  accountRepo,
		logger:       logger,
	}
}

func (service *RegistrationService) RegisterCustomer(request dto.RegisterRequest) (string, error) {
	exists, err := service.customerRepo.ExistsByIdentityNumberOrPhoneNumber(request.IdentityNumber, request.PhoneNumber)
	if err != nil {
		service.logger.Error("Failed to check existing customer: ", err)
		return "", errors.New("internal error")
	}

	if exists {
		return "", errors.New("NIK atau No HP sudah terdaftar")
	}

	customerID, err := service.customerRepo.CreateCustomer(request.Name, request.IdentityNumber, request.PhoneNumber)
	if err != nil {
		service.logger.Warn("Failed to create customer: ", err)
		return "", errors.New("gagal menyimpan data nasabah")
	}

	accountNumber, err := utils.GenerateAccountNumber("10", "01")
	if err != nil {
		service.logger.Error("Failed to create accountNumber: ", err)
		return "", errors.New("gagal membuat akun")
	}

	accountID, err := service.accountRepo.CreateAccount(accountNumber, customerID)
	if err != nil {
		service.logger.Error("Failed to create account: ", err)
		return "", errors.New("gagal membuat akun")
	}

	service.logger.Infof("Nasabah %s berhasil didaftarkan dengan rekening %s - %d", request.Name, accountNumber, accountID)
	return accountNumber, nil
}
