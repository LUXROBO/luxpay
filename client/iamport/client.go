package iamport

import (
	"github.com/luxrobo/luxpay/client"
)

// IamportClient is a client for iamport API
type IamportClient struct {
	apiURL string
	header client.Header
}

// NewIamportClient creates a new IamportClient
func NewIamportClient(iamportKey string, iamportSecret string) *IamportClient {
	accessTokenPayload := AccessTokenPayload{
		ImpKey:    iamportKey,
		ImpSecret: iamportSecret,
	}
	accessTokenStruct := getAccessToken(accessTokenPayload)
	accessToken := accessTokenStruct.Response.AccessToken
	iamportClient := &IamportClient{
		apiURL: "https://api.iamport.kr/",
		header: client.Header{
			Authorization: accessToken,
			ContentType:   "application/json",
		},
	}
	return iamportClient
}

func getAccessToken(accessTokenPayload AccessTokenPayload) AccessTokenResp {
	var accessTokenResp AccessTokenResp
	client.RequestWithPayload(
		accessTokenPayload,
		&accessTokenResp,
		"POST",
		"https://api.iamport.kr/users/getToken",
		client.Header{Authorization: "", ContentType: "application/json"},
	)
	return accessTokenResp
}

// CreateBillingKey requests a billing key
func (ic IamportClient) CreateBillingKey(
	billingKeyPayload interface{},
) interface{} {
	var billingKeyResp IamportBillingKeyResp
	payload := billingKeyPayload.(IamportBillingKeyPayload)
	client.RequestWithPayload(
		billingKeyPayload,
		&billingKeyResp,
		"POST",
		ic.apiURL+"subscribe/customers/"+payload.CustomerUID,
		ic.header,
	)
	return billingKeyResp
}

// MakePayment makes a onetime payment using an issued billing key
func (ic IamportClient) MakePayment(
	paymentPayload interface{},
) interface{} {
	var paymentResp IamportPaymentResp
	client.RequestWithPayload(
		paymentPayload,
		&paymentResp,
		"POST",
		ic.apiURL+"subscribe/payments/again",
		ic.header,
	)
	return paymentResp
}
