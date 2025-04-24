package configs

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSchema    string
	ServiceHost string
	ServicePort string
}

func LoadConfig(logger *logrus.Logger) *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Warn("WARNING: .env file not found, using system environment variables")
	}

	// CLI args
	host := flag.String("host", "0.0.0.0", "REST API host")
	port := flag.String("port", "8080", "REST API port")
	flag.Parse()

	// Env vars
	return &Config{
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "accountdb"),
		DBSchema:    getEnv("DB_SCHEMA", "public"),
		ServiceHost: *host,
		ServicePort: *port,
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (config *Config) DBConnection() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.DBSchema,
	)
}

func (config *Config) ServiceAddress() string {
	return fmt.Sprintf("%s:%s",
		config.ServiceHost, config.ServicePort,
	)
}
