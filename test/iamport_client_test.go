package test

import (
	"os"
	"testing"

	"github.com/luxrobo/luxpay/client/iamport"
	"github.com/stretchr/testify/assert"
)

func setUpIamportMockEnvVars(t *testing.T) {
	t.Setenv("IAMPORT_KEY", "imp_apikey")
	t.Setenv("IAMPORT_SECRET", "ekKoeW8RyKuT0zgaZsUtXXTLQ4AhPFW3ZGseDA6bkA5lamv9OqDMnxyeB9wqOsuO9W3Mx9YSJ4dTqJ3f")

	// Mock Iamport API Inputs
	t.Setenv("CARD_NUMBER", "4092023012341234")
	t.Setenv("CARD_EXPR_YEAR", "19")
	t.Setenv("CARD_EXPR_MONTH", "03")
	t.Setenv("CARD_PASSWORD", "37")
	t.Setenv("BIRTHDAY", "500203")
}

func setUpIamportClient() *iamport.IamportClient {
	iamportKey := os.Getenv("IAMPORT_KEY")
	iamportSecret := os.Getenv("IAMPORT_SECRET")
	iamportClient := iamport.NewIamportClient(iamportKey, iamportSecret)
	return iamportClient
}

func TestIamportCreateBillingKey(t *testing.T) {
	setUpIamportMockEnvVars(t)
	iamportClient := setUpIamportClient()
	cardNumber, cardExprYear, cardExprMonth, cardPassword, birthday := getCardInfo(t)
	billingKeyPayload := iamport.IamportBillingKeyPayload{
		CustomerUID: "test_customer_key",
		CardNumber:  cardNumber,
		Expiry:      "20" + cardExprYear + "-" + cardExprMonth,
		Birth:       birthday,
		Password:    cardPassword[:2],
	}
	billingKeyRespInterface := iamportClient.CreateBillingKey(billingKeyPayload)
	billingKeyResp := billingKeyRespInterface.(iamport.IamportBillingKeyResp)

	assert.Equal(t, -1, billingKeyResp.Code)
}

func TestIamportMakePayment(t *testing.T) {
	setUpIamportMockEnvVars(t)
	iamportClient := setUpIamportClient()

	// Create merchantUID in advance
	merchantUID, _ := generateRandomString(10)
	paymentPayload := iamport.IamportPaymentPayload{
		CustomerUID: "test_customer_key",
		MerchantUID: merchantUID,
		Amount:      100,
		Name:        "test_payment_name",
	}
	paymentRespInterface := iamportClient.MakePayment(
		"", paymentPayload,
	)
	paymentResp := paymentRespInterface.(iamport.IamportPaymentResp)

	assert.Equal(t, 1, paymentResp.Code)
}
