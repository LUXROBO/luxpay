package test

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"testing"

	"github.com/luxrobo/luxpay/client/toss"
	"github.com/stretchr/testify/assert"
)

func setUpTossMockEnvVars(t *testing.T) {
	t.Setenv("DEV_TOSS_SECRET", "test_sk_OALnQvDd2VJl2YzvdBa8Mj7X41mN")
	t.Setenv("CARD_NUMBER", "377989730301234")
	t.Setenv("CARD_EXPR_YEAR", "25")
	t.Setenv("CARD_EXPR_MONTH", "01")
	t.Setenv("CARD_PASSWORD", "1234")
	t.Setenv("BIRTHDAY", "990101")
}

func setUpTossClient(t *testing.T) *toss.TossClient {
	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	tossClient := toss.NewTossClient(tossSecret)
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

func TestCreateBillingKey(t *testing.T) {
	setUpTossMockEnvVars(t)
	tossClient := setUpTossClient(t)
	cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday := getCardInfo(t)
	billingKeyPayload := toss.TossBillingKeyPayload{
		CardNumber:          cardNumber,
		CardExpirationYear:  cardExprYear,
		CardExpirationMonth: cardExprMonth,
		CardPassword:        cardPassword,
		CustomerBirthday:    birthday,
		CustomerKey:         "test_customer_key",
	}
	billingKeyRespInterface := tossClient.CreateBillingKey(billingKeyPayload)
	billingKeyResp := billingKeyRespInterface.(toss.TossBillingKeyResp)

	assert.Equal(t, "tvivarepublica4", billingKeyResp.MID)
	assert.Equal(t, "test_customer_key", billingKeyResp.CustomerKey)
	assert.IsType(t, "string", billingKeyResp.AuthenticatedAt)
	assert.Equal(t, "카드", billingKeyResp.Method)
	assert.IsType(t, "string", billingKeyResp.BillingKey)
	assert.Equal(t, "삼성", billingKeyResp.Company)
	assert.Equal(t, "377989******234", billingKeyResp.Number)
}

func TestMakePayment(t *testing.T) {
	setUpTossMockEnvVars(t)
	tossClient := setUpTossClient(t)
	TestCreateBillingKey(t)

	cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday := getCardInfo(t)
	billingKeyPayload := toss.TossBillingKeyPayload{
		CardNumber:          cardNumber,
		CardExpirationYear:  cardExprYear,
		CardExpirationMonth: cardExprMonth,
		CardPassword:        cardPassword,
		CustomerBirthday:    birthday,
		CustomerKey:         "test_customer_key",
	}
	billingKeyRespInterface := tossClient.CreateBillingKey(billingKeyPayload)
	billingKeyResp := billingKeyRespInterface.(toss.TossBillingKeyResp)

	// Create unique orderID in advance
	uniqueOrderID, _ := generateRandomString(10)
	paymentPayload := toss.TossPaymentPayload{
		OrderName:   "test_order_name",
		OrderID:     uniqueOrderID,
		OrderAmount: "1000",
		CustomerKey: "test_customer_key",
	}
	paymentRespInterface := tossClient.MakePayment(
		billingKeyResp.BillingKey,
		paymentPayload,
	)
	paymentResp := paymentRespInterface.(toss.TossPaymentResp)
	assert.Equal(t, "DONE", paymentResp.Status)
	assert.Equal(t, uniqueOrderID, paymentResp.OrderID)
}
