package models

import "time"

type CurrencyRate struct {
	ID           uint      `gorm:"primaryKey"`
	CurrencyCode string    `gorm:"type:varchar(3);not null"`
	Rate         float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	USDRate      interface{}
}

type Subscription struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"type:varchar(100);not null"`
}
