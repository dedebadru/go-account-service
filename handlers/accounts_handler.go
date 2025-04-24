package handlers

import (
	"net/http"

	"github.com/go-account-service/utils"
	"github.com/go-account-service/services"
	"github.com/go-account-service/dto"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	registrationService 	*services.RegistrationService
	accountService 				*services.AccountService
	transactionService 		*services.TransactionService
}

func NewAccountHandler(registrationService *services.RegistrationService, accountService *services.AccountService, transactionService *services.TransactionService) *AccountHandler {
	return &AccountHandler{registrationService, accountService, transactionService}
}

func (handler *AccountHandler) CustomerRegistration(context echo.Context) error {
	var req dto.RegisterRequest
	if err := context.Bind(&req); err != nil {
		return utils.BadRequestError(context, "Invalid request format")
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.BadRequestError(context, "Invalid request " + err.Error())
	}

	accountNumber, err := handler.registrationService.RegisterCustomer(req)
	if err != nil {
		return utils.BadRequestError(context, err.Error())
	}

	return context.JSON(http.StatusOK, echo.Map{"no_rekening": accountNumber})
}

func (handler *AccountHandler) Credit(context echo.Context) error {
	var req dto.TransactionRequest
	if err := context.Bind(&req); err != nil {
		return utils.BadRequestError(context, "Invalid request format")
	}

	balance, err := handler.transactionService.Saving(req.AccountNumber, req.Amount, "CREDIT")
	if err != nil {
		return utils.BadRequestError(context, err.Error())
	}

	return context.JSON(http.StatusOK, echo.Map{"saldo": balance})
}

func (handler *AccountHandler) Debit(context echo.Context) error {
	var req dto.TransactionRequest
	if err := context.Bind(&req); err != nil {
		return utils.BadRequestError(context, "Invalid request format")
	}

	balance, err := handler.transactionService.Saving(req.AccountNumber, req.Amount, "DEBIT")
	if err != nil {
		return utils.BadRequestError(context, err.Error())
	}

	return context.JSON(http.StatusOK, echo.Map{"saldo": balance})
}

func (handler *AccountHandler) GetBalance(context echo.Context) error {
	accountNumber := context.Param("no_rekening")
	balance, err := handler.accountService.ShowBalance(accountNumber)
	if err != nil {
		return utils.BadRequestError(context, err.Error())
	}

	return context.JSON(http.StatusOK, echo.Map{"saldo": balance})
}
