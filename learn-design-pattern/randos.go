package main

import (
	"context"
)

type paymentMethod string

const (
	Stripe paymentMethod = "STRIPE"
	Adyen  paymentMethod = "ADYEN"
	Wallet paymentMethod = "WALLET"
)

type PaymentReq struct {
	OrderId        string
	UserId         string
	AmountCents    int64
	Currency       string
	Method         paymentMethod
	IdempotencyKey string
}

type PaymentResult struct {
	ProviderTransactionId string
	Status                string
	Message               string
}

type PaymentStrategy interface {
	Authorize(ctx context.Context, req PaymentRequest) (PaymentResult, error)
}
