package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateAccountNumber generates a unique account number with random 8 digits
// Format: 2-digit product code + 2-digit branch code + 8-digit random number
func GenerateAccountNumber(productCode, branchCode string) (string, error) {
	if len(productCode) != 2 {
		return "", fmt.Errorf("productCode must be 2 digits")
	}
	if len(branchCode) != 2 {
		return "", fmt.Errorf("branchCode must be 2 digits")
	}

	num, err := rand.Int(rand.Reader, big.NewInt(90000000))
	if err != nil {
		return "", fmt.Errorf("failed to generate random number: %v", err)
	}
	randomPart := fmt.Sprintf("%08d", num.Int64()+10000000)

	accountNumber := productCode + branchCode + randomPart
	return accountNumber, nil
}