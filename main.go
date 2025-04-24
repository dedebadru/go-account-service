package main

import (
    "database/sql"
    "log"
    "github.com/go-account-service/configs"
    "github.com/go-account-service/handlers"
    "github.com/go-account-service/repositories"
    "github.com/go-account-service/routers"
    "github.com/go-account-service/services"
    "github.com/go-account-service/utils"

    _ "github.com/lib/pq"
    "github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
    // "github.com/sirupsen/logrus"
)

func main() {
    // Load configuration
    cfg := configs.LoadConfig()

    // Open database connection
    db, err := sql.Open("postgres", cfg.DBConnection())
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer db.Close()

    // Create a new logger instance
    logger := utils.NewLogger()

    // Initialize repositories
    customerRepo := repositories.NewCustomerRepository(db)
    accountRepo := repositories.NewAccountRepository(db)
    transactionRepo := repositories.NewTransactionRepository(db)

	// Inisialisasi service
	registrationService := services.NewRegistrationService(customerRepo, accountRepo, logger)
	accountService := services.NewAccountService(accountRepo, logger)
	transactionService := services.NewTransactionService(accountRepo, transactionRepo, logger)

	// Inisialisasi handler
	handler := handlers.NewAccountHandler(registrationService, accountService, transactionService)

    // Create new Echo instance
    e := echo.New()

    // Add custom middleware for logging and error recovery
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Initialize routes
    routers.InitRoutes(e, handler)

    // Start the Echo server
    logger.Infof("Starting server at %s", cfg.ServiceAddress())
    if err := e.Start(cfg.ServiceAddress()); err != nil {
        logger.Fatalf("Failed to start server: %v", err)
    }
}