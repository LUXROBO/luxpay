package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type BillingKeyPayload struct {
	CardNumber          string `json:"cardNumber"`
	CardExpirationYear  string `json:"cardExpirationYear"`
	CardExpirationMonth string `json:"cardExpirationMonth"`
	CardPassword        string `json:"cardPassword"`
	CustomerBirthday    string `json:"customerBirthday"`
	CustomerKey         string `json:"customerKey"`
}

type BillingKeyResp struct {
	MID             string `json:"mId"`
	CustomerKey     string `json:"customerKey"`
	AuthenticatedAt string `json:"authenticatedAt"`
	Method          string `json:"method"`
	BillingKey      string `json:"billingKey"`
	Company         string `json:"cardCompany"`
	Number          string `json:"cardNumber"`
}

type PaymentPayload struct {
	OrderName   string `json:"orderName"`
	OrderId     string `json:"orderId"`
	OrderAmount string `json:"amount"`
	CustomerKey string `json:"customerKey"`
}

type PaymentResp struct {
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
	MId                 string `json:"mId"`
	Method              string `json:"method"`
	MobilePhone         string `json:"mobilePhone"`
	OrderId             string `json:"orderId"`
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

const apiUrl string = "https://api.tosspayments.com/"
const contentType string = "application/json"

type Header struct {
	Authorization string
	ContentType   string
}

type TossClient struct {
	apiUrl string
	header Header
}

func NewTossClient(tossSecret string) *TossClient {
	authToken := getAuthToken(tossSecret)
	return &TossClient{
		apiUrl: apiUrl,
		header: Header{
			Authorization: "Basic " + authToken,
			ContentType:   contentType,
		},
	}
}

func getAuthToken(tossSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(tossSecret + ":"))
}

func (tc TossClient) CreateBillingKey(
	cardNumber string,
	cardExpirationYear string,
	cardExpirationMonth string,
	cardPassword string,
	customerBirthday string,
	customerUid string,
) BillingKeyResp {

	payload := BillingKeyPayload{
		CardNumber:          cardNumber,
		CardExpirationYear:  cardExpirationYear,
		CardExpirationMonth: cardExpirationMonth,
		CardPassword:        cardPassword,
		CustomerBirthday:    customerBirthday,
		CustomerKey:         customerUid,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		apiUrl+"v1/billing/authorizations/card",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", tc.header.Authorization)
	req.Header.Add("Content-Type", tc.header.ContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data BillingKeyResp
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func (tc TossClient) MakePayment(
	billingKey string,
	orderName string,
	orderId string,
	orderAmount string,
	customerUid string,
) PaymentResp {
	payload := PaymentPayload{
		OrderName:   orderName,
		OrderId:     orderId,
		OrderAmount: orderAmount,
		CustomerKey: customerUid,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		apiUrl+"v1/billing/"+billingKey,
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", tc.header.Authorization)
	req.Header.Add("Content-Type", tc.header.ContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data PaymentResp
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	tossClient := NewTossClient(tossSecret)
	fmt.Println(tossClient)
	billingKeyResp := tossClient.CreateBillingKey(
		os.Getenv("CARD_NUMBER"),     // CARD NUMBER
		os.Getenv("CARD_EXPR_YEAR"),  // YY
		os.Getenv("CARD_EXPR_MONTH"), // MM
		os.Getenv("CARD_PASSWORD"),   // DDDD
		os.Getenv("BIRTHDAY"),        // YYMMDD
		"test_customer_key",          // RANDOM STRING
	)
	fmt.Println("BillingKeyResp:", billingKeyResp)
	paymentResp := tossClient.MakePayment(
		billingKeyResp.BillingKey,
		"test_order_name",
		"test_order_id",
		"1000",
		"test_customer_key",
	)
	fmt.Println("PaymentResp:", paymentResp)
}
