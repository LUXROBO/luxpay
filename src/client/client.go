package client

type Client interface {
	CreateBillingKey(billingKeyPayload BillingKeyPayload) BillingKeyResp
	MakePayment(billingKey string, paymentPayload PaymentPayload) PaymentResp
}

type Header struct {
	Authorization string
	ContentType   string
}

func NewClient(pgService string, tossSecret string) Client {
	client := NewTossClient(tossSecret)
	return client
}
