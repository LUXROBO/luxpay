package test

import (
	"luxpay/src/client"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpTossClient() *client.TossClient {
	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	tossClient := client.NewTossClient(tossSecret)
	return tossClient
}

func TestCreateBillingKey(t *testing.T) {
	tossClient := setUpTossClient()
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

func TestMakePayment(t *testing.T) {
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

	paymentResp := tossClient.MakePayment(
		billingKeyResp.BillingKey,
		"test_order_name", // FAKE ORDER NAME
		"test_order_id",
		"1000",
		"test_customer_key",
	)
	assert.NotNil(t, paymentResp)
}
