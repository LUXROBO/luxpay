package toss

import (
	"encoding/base64"

	"github.com/luxrobo/luxpay/client"
)

// TossClient is a client for toss API
type TossClient struct {
	apiURL     string
	header     client.Header
	billingKey *string
}

// SetBillingKey sets a billing key to a given tossClient
func (tc *TossClient) SetBillingKey(billingKey string) {
	tc.billingKey = &billingKey
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
		billingKey: nil,
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
	var billingKeyResp TossBillingKeyResp
	httpInfo := client.HTTPInfo{
		Method: "POST",
		URL:    tc.apiURL + "v1/billing/authorizations/card",
		Header: tc.header,
	}
	client.RequestWithPayload(
		billingKeyPayload,
		&billingKeyResp,
		httpInfo,
	)
	return billingKeyResp
}

// MakePayment makes a onetime payment using an issued billing key
func (tc TossClient) MakePayment(
	paymentPayload interface{},
) interface{} {
	var paymentResp TossPaymentResp
	httpInfo := client.HTTPInfo{
		Method: "POST",
		URL:    tc.apiURL + "v1/billing/" + *tc.billingKey,
		Header: tc.header,
	}
	client.RequestWithPayload(
		paymentPayload,
		&paymentResp,
		httpInfo,
	)
	return paymentResp
}
