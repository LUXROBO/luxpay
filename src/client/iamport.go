package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const impApiUrl string = "https://api.iamport.kr/"

type IamportClient struct {
	apiUrl        string
	Authorization string
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
	accessTokenStruct := getAccessToken(iamportKey, iamportSecret)
	accessToken := accessTokenStruct.Response.AccessToken
	return &IamportClient{
		apiUrl:        "https://api.iamport.kr/",
		Authorization: accessToken,
	}
}

func getAccessToken(iamportKey string, iamportSecret string) AccessTokenResp {
	payload := AccessTokenPayload{
		ImpKey:    iamportKey,
		ImpSecret: iamportSecret,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		impApiUrl+"users/getToken",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data AccessTokenResp
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func (ic IamportClient) CreateCustomer(
	customerUid string,
	cardNumber string,
	expiry string,
) CustomerResp {
	payload := CustomerPayload{
		CustomerUid: customerUid,
		CardNumber:  cardNumber,
		Expiry:      expiry,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		impApiUrl+"subscribe/customers/"+customerUid,
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", ic.Authorization)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data CustomerResp
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func (ic IamportClient) MakePayment(
	customerUid string,
	merchantUid string,
	amount int,
	paymentName string,
) MakePaymentResp {
	payload := MakePaymentPayload{
		CustomerUid: customerUid,
		MerchantUid: merchantUid,
		Amount:      amount,
		Name:        paymentName,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		impApiUrl+"subscribe/payments/again",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", ic.Authorization)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data MakePaymentResp
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}
