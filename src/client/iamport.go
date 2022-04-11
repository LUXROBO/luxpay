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
