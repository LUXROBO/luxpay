package test

import (
	"luxpay/src/client"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpMockEnvVars(t *testing.T) {
	t.Setenv("DEV_TOSS_SECRET", "test_sk_OALnQvDd2VJl2YzvdBa8Mj7X41mN")
	t.Setenv("CARD_NUMBER", "377989730301234")
	t.Setenv("CARD_EXPR_YEAR", "25")
	t.Setenv("CARD_EXPR_MONTH", "01")
	t.Setenv("CARD_PASSWORD", "1234")
	t.Setenv("BIRTHDAY", "990101")
}

func setUpTossClient(t *testing.T) *client.TossClient {
	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	tossClient := client.NewTossClient(tossSecret)
	return tossClient
}

func getCardInfo(t *testing.T) (string, string, string, string, string) {
	cardNumber := os.Getenv("CARD_NUMBER")
	cardExprYear := os.Getenv("CARD_EXPR_YEAR")
	cardExprMonth := os.Getenv("CARD_EXPR_MONTH")
	cardPassword := os.Getenv("CARD_PASSWORD")
	birthday := os.Getenv("BIRTHDAY")
	return cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday
}

func TestCreateBillingKey(t *testing.T) {
	setUpMockEnvVars(t)
	tossClient := setUpTossClient(t)
	cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday := getCardInfo(t)
	billingKeyResp := tossClient.CreateBillingKey(
		cardNumber,
		cardExprYear,
		cardExprMonth,
		cardPassword,
		birthday,
		"test_customer_key",
	)

	assert.Equal(t, "tvivarepublica4", billingKeyResp.MID)
	assert.Equal(t, "test_customer_key", billingKeyResp.CustomerKey)
	assert.IsType(t, "string", billingKeyResp.AuthenticatedAt)
	assert.Equal(t, "카드", billingKeyResp.Method)
	assert.IsType(t, "string", billingKeyResp.BillingKey)
	assert.Equal(t, "삼성", billingKeyResp.Company)
	assert.Equal(t, "377989******234", billingKeyResp.Number)
}

func TestMakePayment(t *testing.T) {
	setUpMockEnvVars(t)
	tossClient := setUpTossClient(t)
	// cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday := getCardInfo(t)
	// t.Run("[Test 1]", func(t *testing.T) {

	// }

	// t.Run("[Test 2]", func(t *testing.T) {

	// }

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
