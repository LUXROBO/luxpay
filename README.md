# LUXPAY (럭스페이)
<div align="center">

[![Go Version](https://img.shields.io/github/go-mod/go-version/luxrobo/luxpay)](https://github.com/LUXROBO/luxpay)
[![Build Workflow Status (Github Actions)](https://img.shields.io/github/workflow/status/LUXROBO/luxpay/Go)](https://github.com/LUXROBO/luxpay/actions)
[![CodeClimate Maintainability](https://img.shields.io/codeclimate/maintainability/LUXROBO/luxpay)](https://github.com/LUXROBO/luxpay/tree/main)
[![CodeClimate Issues](https://img.shields.io/codeclimate/issues/LUXROBO/luxpay)](https://github.com/LUXROBO/luxpay/tree/main)
[![CodeClimate Coverage](https://img.shields.io/codeclimate/coverage/LUXROBO/luxpay)](https://github.com/LUXROBO/luxpay/tree/main/test)
[![Github LICENSE](https://img.shields.io/github/license/LUXROBO/luxpay)](https://github.com/LUXROBO/luxpay/blob/main/LICENSE)
[![Lines of Code](https://img.shields.io/tokei/lines/github/LUXROBO/luxpay)](https://github.com/LUXROBO/luxpay/tree/develop/src)

</div>

## Description
> A collection of payment functionalities written in Go

## Features
- Implement Iamport and Toss Client as a backend
- Subscription payment functions
- Self-certification via Iamport

## Installation
```bash
go get github.com/luxrobo/luxpay
```

## Usage
```
// TossClient Usage Example
tossClient := client.NewTossClient(tossSecret)
billingKeyResp := tossClient.CreateBillingKey(billingKeyPayload)
billingKey := billingKeyResp.BillingKey
paymentResp := tossClient.MakePayment(billingKey, paymentPayload)
fmt.Println("paymentResp.Status:", paymentResp.Status)

// IamportClient Usage Example
iamportClient := client.NewIamportClient(iamportKey, iamportSecret)
billingKeyResp := iamportClient.CreateBillingKey(billingKeyPayload)
billingKey := billingKeyResp.BillingKey
paymentResp := iamportClient.MakePayment(billingKey, paymentPayload)
fmt.Println("paymentResp.Status:", paymentResp.Status)
```
