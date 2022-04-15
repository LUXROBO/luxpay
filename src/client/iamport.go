package client

import (
	"encoding/json"
)

type IamportClient struct {
	apiUrl string
	header Header
}

type AccessTokenPayload struct {
	ImpKey    string `json:"imp_key"`
	ImpSecret string `json:"imp_secret"`
}

type AccessTokenResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		AccessToken string `json:"access_token"`
		ExpiredAt   int    `json:"expired_at"`
		Now         int    `json:"now"`
	} `json:"response"`
}

type CustomerPayload struct {
	CustomerUid string `json:"customer_uid"`
	CardNumber  string `json:"card_number"`
	Expiry      string `json:"expiry"`
	Birth       string `json:"birth"`
	Password    string `json:"pwd_2digit"`
}

type CustomerResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		CustomerUid      string `json:"customer_uid"`
		PgProvider       string `json:"pg_provider"`
		PgId             string `json:"pg_id"`
		CardName         string `json:"card_name"`
		CardCode         string `json:"card_code"`
		CardNumber       string `json:"card_number"`
		CardType         string `json:"card_type"`
		CustomerName     string `json:"customer_name"`
		CustomerTel      string `json:"customer_tel"`
		CustomerEmail    string `json:"customer_email"`
		CustomerAddr     string `json:"customer_addr"`
		CustomerPostcode string `json:"customer_postcode"`
		Inserted         int    `json:"inserted"`
		Updated          int    `json:"updated"`
	} `json:"response"`
}

type MakePaymentPayload struct {
	CustomerUid string `json:"customer_uid"`
	MerchantUid string `json:"merchant_uid"`
	Amount      int    `json:"amount"`
	Name        string `json:"name"`
}

type MakePaymentResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		ImpUid        string `json:"imp_uid"`
		MerchantUid   string `json:"merchant_uid"`
		PayMethod     string `json:"pay_method"`
		Channel       string `json:"channel"`
		PgProvider    string `json:"pg_provider"`
		EmbPgProvider string `json:"emb_pg_provider"`
		PgTid         string `json:"pg_tid"`
		PgId          string `json:"pg_id"`
		Escrow        bool   `json:"escrow"`
		ApplyNum      string `json:"apply_num"`
		BankCode      string `json:"bank_code"`
		BankName      string `json:"bank_name"`
		CardCode      string `json:"card_code"`
		CardName      string `json:"card_name"`
		CardQuota     int    `json:"card_quota"`
		CardNumber    string `json:"card_number"`
		CardType      string `json:"card_type"`
	} `json:"response"`
}

func NewIamportClient(iamportKey string, iamportSecret string) *IamportClient {
	accessTokenPayload := AccessTokenPayload{
		ImpKey:    iamportKey,
		ImpSecret: iamportSecret,
	}
	accessTokenStruct := getAccessToken(accessTokenPayload)
	accessToken := accessTokenStruct.Response.AccessToken
	return &IamportClient{
		apiUrl: "https://api.iamport.kr/",
		header: Header{
			Authorization: accessToken,
			ContentType:   "application/json",
		},
	}
}

func getAccessToken(accessTokenPayload AccessTokenPayload) AccessTokenResp {
	jsonPayload, _ := json.Marshal(accessTokenPayload)
	resp := requestWithPayload(
		jsonPayload,
		"POST",
		"https://api.iamport.kr/users/getToken",
		Header{Authorization: "", ContentType: "application/json"},
	)
	var accessTokenResp AccessTokenResp
	json.NewDecoder(resp.Body).Decode(&accessTokenResp)
	return accessTokenResp
}

func (ic IamportClient) CreateCustomer(
	customerPayload CustomerPayload,
) CustomerResp {
	jsonPayload, _ := json.Marshal(customerPayload)
	resp := requestWithPayload(
		jsonPayload,
		"POST",
		ic.apiUrl+"subscribe/customers/"+customerPayload.CustomerUid,
		ic.header,
	)
	var customerResp CustomerResp
	json.NewDecoder(resp.Body).Decode(&customerResp)
	return customerResp
}

func (ic IamportClient) MakePayment(
	makePaymentPayload MakePaymentPayload,
) MakePaymentResp {
	jsonPayload, _ := json.Marshal(makePaymentPayload)
	resp := requestWithPayload(
		jsonPayload,
		"POST",
		ic.apiUrl+"subscribe/payments/again",
		ic.header,
	)
	var makePaymentResp MakePaymentResp
	json.NewDecoder(resp.Body).Decode(&makePaymentResp)
	return makePaymentResp
}
