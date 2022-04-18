package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client wraps child clients such as IamportClient
type Client interface {
	CreateBillingKey(billingKeyPayload interface{}) interface{}
	MakePayment(paymentPayload interface{}) interface{}
}

// Header includes header information of request instance
type Header struct {
	Authorization string
	ContentType   string
}

// HTTPInfo includes http information of request instance
type HTTPInfo struct {
	Method string
	URL    string
	Header Header
}

// RequestWithPayload makes request with a given payload
func RequestWithPayload(
	payload interface{},
	response interface{},
	httpInfo HTTPInfo,
) interface{} {
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest(
		httpInfo.Method,
		httpInfo.URL,
		bytes.NewBuffer(jsonPayload),
	)

	httpHeader := httpInfo.Header
	if httpHeader.Authorization != "" {
		req.Header.Add("Authorization", httpHeader.Authorization)
	}
	req.Header.Add("Content-Type", httpHeader.ContentType)

	client := &http.Client{}
	resp, _ := client.Do(req)

	json.NewDecoder(resp.Body).Decode(response)
	return response
}
