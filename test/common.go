package test

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"testing"
)

func getCardInfo(t *testing.T) (string, string, string, string, string) {
	cardNumber := os.Getenv("CARD_NUMBER")
	cardExprYear := os.Getenv("CARD_EXPR_YEAR")
	cardExprMonth := os.Getenv("CARD_EXPR_MONTH")
	cardPassword := os.Getenv("CARD_PASSWORD")
	birthday := os.Getenv("BIRTHDAY")
	return cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
