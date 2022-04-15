package toss

type TossBillingKeyPayload struct {
	CardNumber          string `json:"cardNumber"`
	CardExpirationYear  string `json:"cardExpirationYear"`
	CardExpirationMonth string `json:"cardExpirationMonth"`
	CardPassword        string `json:"cardPassword"`
	CustomerBirthday    string `json:"customerBirthday"`
	CustomerKey         string `json:"customerKey"`
}

type TossBillingKeyResp struct {
	MID             string `json:"mId"`
	CustomerKey     string `json:"customerKey"`
	AuthenticatedAt string `json:"authenticatedAt"`
	Method          string `json:"method"`
	BillingKey      string `json:"billingKey"`
	Company         string `json:"cardCompany"`
	Number          string `json:"cardNumber"`
}

type TossPaymentPayload struct {
	OrderName   string `json:"orderName"`
	OrderID     string `json:"orderId"`
	OrderAmount string `json:"amount"`
	CustomerKey string `json:"customerKey"`
}

type TossPaymentResp struct {
	ApprovedAt    string `json:"approvedAt"`
	BalanceAmount string `json:"balanceAmount"`
	Cancels       string `json:"cancels"`
	Card          struct {
		AcquireStatus         string `json:"acquireStatus"`
		ApproveNo             string `json:"approveNo"`
		CardType              string `json:"cardType"`
		CardCompany           string `json:"company"`
		InstallmentPlanMonths int    `json:"installmentPlanMonths"`
		IsInterestFree        bool   `json:"isInterestFree"`
		CardNumber            string `json:"number"`
		OwnerType             string `json:"ownerType"`
		ReceiptUrl            string `json:"receiptUrl"`
		UseCardPoint          bool   `json:"useCardPoint"`
	} `json:"card"`
	CashReceipt         string `json:"cashReceipt"`
	Currency            string `json:"currency"`
	Discount            string `json:"discount"`
	DiscountAmount      int    `json:"discountAmount"`
	GiftCertificate     string `json:"giftCertificate"`
	IsPartialCancelable bool   `json:"isPartialCancelable"`
	MID                 string `json:"mId"`
	Method              string `json:"method"`
	MobilePhone         string `json:"mobilePhone"`
	OrderID             string `json:"orderId"`
	PaymentKey          string `json:"paymentKey"`
	RequestedAt         string `json:"requestedAt"`
	Secret              string `json:"secret"`
	Status              string `json:"status"`
	TotalAmount         int    `json:"totalAmount"`
	UseCashReceipt      bool   `json:"useCashReceipt"`
	UseDiscount         bool   `json:"useDiscount"`
	UseEscrow           bool   `json:"useEscrow"`
	VirtualAccount      string `json:"virtualAccount"`
}
