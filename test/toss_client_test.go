package test

import (
	"luxpay/src/client"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	client.NewTossClient(tossSecret)
	code := m.Run()
	os.Exit(code)
}

func TestCreateBillingKey(t *testing.T) {
	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	tossClient := client.NewTossClient(tossSecret)
	billingKeyResp := tossClient.CreateBillingKey(
		os.Getenv("CARD_NUMBER"),     // CARD NUMBER
		os.Getenv("CARD_EXPR_YEAR"),  // YY
		os.Getenv("CARD_EXPR_MONTH"), // MM
		os.Getenv("CARD_PASSWORD"),   // DDDD
		os.Getenv("BIRTHDAY"),        // YYMMDD
		"test_customer_key",          // RANDOM STRING
	)
	assert.NotNil(t, billingKeyResp)
}
