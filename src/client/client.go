package client

type Client interface {
	CreateBillingKey(billingKeyPayload BillingKeyPayload) BillingKeyResp
	MakePayment(billingKey string, paymentPayload PaymentPayload) PaymentResp
}
