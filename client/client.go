package client

import (
	"bytes"
	"net/http"
)

type Client interface {
	CreateBillingKey(billingKeyPayload interface{}) interface{}
	MakePayment(billingKey string, paymentPayload interface{}) interface{}
}

type Header struct {
	Authorization string
	ContentType   string
}

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
