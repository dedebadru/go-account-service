package routers

import (
    "github.com/labstack/echo/v4"
    "github.com/go-account-service/handlers"
)

func InitRoutes(e *echo.Echo, account *handlers.AccountHandler) {
    api := e.Group("/")

    api.POST("daftar", account.CustomerRegistration)
    api.POST("tabung", account.Credit)
    api.POST("tarik", account.Debit)
    api.GET("saldo/:no_rekening", account.GetBalance)
}
