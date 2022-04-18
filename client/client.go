package client

import (
	"bytes"
	"net/http"
)

// Client wraps child clients such as IamportClient
type Client interface {
	CreateBillingKey(billingKeyPayload interface{}) interface{}
	MakePayment(billingKey string, paymentPayload interface{}) interface{}
}

// Header includes header information of request instance
type Header struct {
	Authorization string
	ContentType   string
}

// RequestWithPayload makes request with a given payload
func RequestWithPayload(
	jsonPayload []byte,
	httpMethod string,
	apiEndPoint string,
	header Header,
) *http.Response {
	req, _ := http.NewRequest(
		httpMethod,
		apiEndPoint,
		bytes.NewBuffer(jsonPayload),
	)

	if header.Authorization != "" {
		req.Header.Add("Authorization", header.Authorization)
	}
	req.Header.Add("Content-Type", header.ContentType)

	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}
