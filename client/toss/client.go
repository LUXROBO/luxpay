package toss

import (
	"encoding/base64"
	"encoding/json"

	"github.com/luxrobo/luxpay/client"
)

type TossClient struct {
	apiUrl string
	header client.Header
}

func NewTossClient(tossSecret string) *TossClient {
	authToken := getAuthToken(tossSecret)
	tossClient := &TossClient{
		apiUrl: "https://api.tosspayments.com/",
		header: client.Header{
			Authorization: "Basic " + authToken,
			ContentType:   "application/json",
		},
	}
	return tossClient
}

func getAuthToken(tossSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(tossSecret + ":"))
}

func (tc TossClient) CreateBillingKey(
	billingKeyPayload interface{},
) interface{} {
	payload := billingKeyPayload.(TossBillingKeyPayload)
	jsonPayload, _ := json.Marshal(payload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		tc.apiUrl+"v1/billing/authorizations/card",
		tc.header,
	)
	var billingKeyResp TossBillingKeyResp
	json.NewDecoder(resp.Body).Decode(&billingKeyResp)
	return billingKeyResp
}

func (tc TossClient) MakePayment(
	billingKey string,
	paymentPayload interface{},
) interface{} {
	payload := paymentPayload.(TossPaymentPayload)
	jsonPayload, _ := json.Marshal(payload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		tc.apiUrl+"v1/billing/"+billingKey,
		tc.header,
	)
	var paymentResp TossPaymentResp
	json.NewDecoder(resp.Body).Decode(&paymentResp)
	return paymentResp
}
