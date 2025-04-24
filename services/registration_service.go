package services

import (
	"errors"
	"fmt"
	"github.com/go-account-service/dto"
	"github.com/go-account-service/repositories"
	"github.com/go-account-service/utils"
	"github.com/sirupsen/logrus"
)

type RegistrationService struct {
	customerRepo	*repositories.CustomerRepository
	accountRepo		*repositories.AccountRepository
	logger        *logrus.Logger
}

func NewRegistrationService(
	customerRepo *repositories.CustomerRepository,
	accountRepo *repositories.AccountRepository,
	logger *logrus.Logger,
) *RegistrationService {
	return &RegistrationService{
		customerRepo:    customerRepo,
		accountRepo:     accountRepo,
		logger:          logger,
	}
}

func (s *RegistrationService) RegisterCustomer(req dto.RegisterRequest) (string, error) {
	exists, err := s.customerRepo.ExistsByIdentityNumberOrPhoneNumber(req.IdentityNumber, req.PhoneNumber)
	if err != nil {
		s.logger.Error("Failed to check existing customer: ", err)
		return "", errors.New("internal error")
	}

	if exists {
		return "", errors.New("NIK atau No HP sudah terdaftar")
	}

	customerID, err := s.customerRepo.CreateCustomer(req.Name, req.IdentityNumber, req.PhoneNumber)
	if err != nil {
		s.logger.Error("Failed to create customer: ", err)
		return "", errors.New("gagal menyimpan data nasabah")
	}
	
	accountNumber, err := utils.GenerateAccountNumber("10", "01")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Generated Rekening:", accountNumber)
	}

	accountID, err := s.accountRepo.CreateAccount(accountNumber, customerID)
	if err != nil {
		s.logger.Error("Failed to create account: ", err)
		return "", errors.New("gagal membuat akun")
	}

	s.logger.Infof("Nasabah %s berhasil didaftarkan dengan rekening %s - %s", req.Name, accountNumber, accountID)
	return accountNumber, nil
}
