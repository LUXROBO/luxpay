package client

import (
	"bytes"
	"net/http"
)

func requestWithPayload(
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
