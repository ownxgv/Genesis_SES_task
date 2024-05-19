package models

type CurrencyRate struct {
	ID      uint    `gorm:"primaryKey"`
	USDRate float64 `gorm:"not null"`
}

type Subscription struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"not null;unique"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SubscriptionRequest struct {
	Email string `json:"email" binding:"required,email"`
}
