package client

import (
	"encoding/base64"
	"encoding/json"
)

type TossClient struct {
	apiUrl string
	header Header
}

func NewTossClient(tossSecret string) *TossClient {
	authToken := getAuthToken(tossSecret)
	tossClient := &TossClient{
		apiUrl: "https://api.tosspayments.com/",
		header: Header{
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
	billingKeyPayload BillingKeyPayload,
) BillingKeyResp {
	jsonPayload, _ := json.Marshal(billingKeyPayload)
	resp := requestWithPayload(
		jsonPayload,
		"POST",
		tc.apiUrl+"v1/billing/authorizations/card",
		tc.header,
	)
	var billingKeyResp BillingKeyResp
	json.NewDecoder(resp.Body).Decode(&billingKeyResp)
	return billingKeyResp
}

func (tc TossClient) MakePayment(
	billingKey string,
	paymentPayload PaymentPayload,
) PaymentResp {
	jsonPayload, _ := json.Marshal(paymentPayload)
	resp := requestWithPayload(
		jsonPayload,
		"POST",
		tc.apiUrl+"v1/billing/"+billingKey,
		tc.header,
	)
	var paymentResp PaymentResp
	json.NewDecoder(resp.Body).Decode(&paymentResp)
	return paymentResp
}
