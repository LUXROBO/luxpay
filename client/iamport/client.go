package iamport

import (
	"encoding/json"

	"github.com/luxrobo/luxpay/client"
)

type IamportClient struct {
	apiUrl string
	header client.Header
}

func NewIamportClient(iamportKey string, iamportSecret string) *IamportClient {
	accessTokenPayload := AccessTokenPayload{
		ImpKey:    iamportKey,
		ImpSecret: iamportSecret,
	}
	accessTokenStruct := getAccessToken(accessTokenPayload)
	accessToken := accessTokenStruct.Response.AccessToken
	iamportClient := &IamportClient{
		apiUrl: "https://api.iamport.kr/",
		header: client.Header{
			Authorization: accessToken,
			ContentType:   "application/json",
		},
	}
	return iamportClient
}

func getAccessToken(accessTokenPayload AccessTokenPayload) AccessTokenResp {
	jsonPayload, _ := json.Marshal(accessTokenPayload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		"https://api.iamport.kr/users/getToken",
		client.Header{Authorization: "", ContentType: "application/json"},
	)
	var accessTokenResp AccessTokenResp
	json.NewDecoder(resp.Body).Decode(&accessTokenResp)
	return accessTokenResp
}

func (ic IamportClient) CreateBillingKey(
	billingKeyPayload interface{},
) interface{} {
	payload := billingKeyPayload.(IamportBillingKeyPayload)
	jsonPayload, _ := json.Marshal(payload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		ic.apiUrl+"subscribe/customers/"+payload.CustomerUID,
		ic.header,
	)
	var iamportBillingKeyResp IamportBillingKeyResp
	json.NewDecoder(resp.Body).Decode(&iamportBillingKeyResp)
	return iamportBillingKeyResp
}

func (ic IamportClient) MakePayment(
	billingKey string,
	paymentPayload interface{},
) interface{} {
	payload := paymentPayload.(IamportPaymentPayload)
	jsonPayload, _ := json.Marshal(payload)
	resp := client.RequestWithPayload(
		jsonPayload,
		"POST",
		ic.apiUrl+"subscribe/payments/again",
		ic.header,
	)
	var iamportPaymentResp IamportPaymentResp
	json.NewDecoder(resp.Body).Decode(&iamportPaymentResp)
	return iamportPaymentResp
}
