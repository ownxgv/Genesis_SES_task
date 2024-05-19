package repositories

import (
	"github.com/your-username/currency-service/models"
	"gorm.io/gorm"
)

type CurrencyRepository struct {
	DB *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{DB: db}
}

func (r *CurrencyRepository) GetCurrencyRate() (*models.CurrencyRate, error) {
	var rate models.CurrencyRate
	err := r.DB.Last(&rate).Error
	return &rate, err
}

func (r *CurrencyRepository) SaveCurrencyRate(rate *models.CurrencyRate) error {
	return r.DB.Create(rate).Error
}

func (r *CurrencyRepository) SubscribeEmail(email string) error {
	subscription := models.Subscription{Email: email}
	return r.DB.Create(&subscription).Error
}

func (r *CurrencyRepository) GetSubscriptions() ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := r.DB.Find(&subscriptions).Error
	return subscriptions, err
}
