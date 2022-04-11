package main

import (
	"fmt"
	"os"

	"luxpay/src/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	tossClient := client.NewTossClient(tossSecret)
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

	impKey := os.Getenv("IAMPORT_KEY")
	impSecret := os.Getenv("IAMPORT_SECRET")

	impClient := client.NewIamportClient(impKey, impSecret)
	fmt.Println("IamportClient:", impClient)

	customerResp := impClient.CreateCustomer(
		"test_customer_uid",      // RANDOM STRING
		os.Getenv("CARD_NUMBER"), // CARD NUMBER
		"2025-07",                // CARD EXPIRY (YYYY-MM)
	)
	fmt.Println("CustomerResp:", customerResp)

	impPaymentResp := impClient.MakePayment(
		"test_customer_uid",
		"test_merchant_uid",
		1000,
		"test_payment_name",
	)

	fmt.Println("ImpPaymentResp:", impPaymentResp)
}
