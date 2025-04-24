package configs

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     	string
	DBPort     	string
	DBUser     	string
	DBPassword 	string
	DBName     	string
	DBSchema   	string
	ServiceHost	string
	ServicePort	string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING: .env file not found, using system environment variables")
	}

	// CLI args
	host := flag.String("host", "0.0.0.0", "REST API host")
	port := flag.String("port", "8080", "REST API port")
	flag.Parse()

	// Env vars
	return &Config{
		DBHost:     	getEnv("DB_HOST", "localhost"),
		DBPort:     	getEnv("DB_PORT", "5432"),
		DBUser:     	getEnv("DB_USER", "postgres"),
		DBPassword: 	getEnv("DB_PASSWORD", ""),
		DBName:     	getEnv("DB_NAME", "accountdb"),
		DBSchema:   	getEnv("DB_SCHEMA", "public"),
		ServiceHost:	*host,
		ServicePort:	*port,
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (c *Config) DBConnection() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSchema,
	)
}

func (c *Config) ServiceAddress() string {
	return fmt.Sprintf("%s:%s",
		c.ServiceHost, c.ServicePort,
	)
}