package models

import "gorm.io/gorm"

type CurrencyRate struct {
	gorm.Model
	USDRate float64 `json:"usd_rate"`
}

type Subscription struct {
	gorm.Model
	Email string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SubscriptionRequest struct {
	Email string `json:"email" binding:"required,email"`
}
