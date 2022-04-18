package toss

import (
	"encoding/base64"
	"encoding/json"

	"github.com/luxrobo/luxpay/client"
)

// TossClient is a client for toss API
type TossClient struct {
	apiURL     string
	header     client.Header
	BillingKey *string
}

// NewTossClient creates a new TossClient
func NewTossClient(tossSecret string) *TossClient {
	authToken := getAuthToken(tossSecret)
	tossClient := &TossClient{
		apiURL: "https://api.tosspayments.com/",
		header: client.Header{
			Authorization: "Basic " + authToken,
			ContentType:   "application/json",
		},
		BillingKey: nil,
	}
	return tossClient
}

func getAuthToken(tossSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(tossSecret + ":"))
}

// CreateBillingKey requests a billing key
func (tc TossClient) CreateBillingKey(
	billingKeyPayload interface{},
) interface{} {
	payload := billingKeyPayload.(TossBillingKeyPayload)
	jsonPayload, _ := json.Marshal(payload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		tc.apiURL+"v1/billing/authorizations/card",
		tc.header,
	)
	var billingKeyResp TossBillingKeyResp
	json.NewDecoder(resp.Body).Decode(&billingKeyResp)
	return billingKeyResp
}

// MakePayment makes a onetime payment using an issued billing key
func (tc TossClient) MakePayment(
	paymentPayload interface{},
) interface{} {
	payload := paymentPayload.(TossPaymentPayload)
	jsonPayload, _ := json.Marshal(payload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		tc.apiURL+"v1/billing/"+*tc.BillingKey,
		tc.header,
	)
	var paymentResp TossPaymentResp
	json.NewDecoder(resp.Body).Decode(&paymentResp)
	return paymentResp
}
