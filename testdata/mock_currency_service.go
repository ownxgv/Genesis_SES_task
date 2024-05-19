// mocks/mock_currency_service.go
package mocks

import (
	"github.com/ownxgv/Genesis_SES_task/models"
)

// MockCurrencyService представляет мок-сервис для тестирования.
type MockCurrencyService struct{}

// NewMockCurrencyService создает новый экземпляр MockCurrencyService.
func NewMockCurrencyService() *MockCurrencyService {
	return &MockCurrencyService{}
}

// GetCurrencyRate возвращает фиктивную ставку валюты для тестирования.
func (m *MockCurrencyService) GetCurrencyRate() (*models.CurrencyRate, error) {
	// Здесь вы можете вернуть фиктивные данные для тестирования
	return &models.CurrencyRate{USDRate: 27.5}, nil
}

// SubscribeEmail имитирует подписку на электронную почту.
func (m *MockCurrencyService) SubscribeEmail(email string) error {
	// Здесь вы можете имитировать успешную подписку на электронную почту для тестирования
	return nil
}

// GetSubscriptions возвращает фиктивные подписки для тестирования.
func (m *MockCurrencyService) GetSubscriptions() ([]models.Subscription, error) {
	// Здесь вы можете вернуть фиктивные подписки для тестирования
	return []models.Subscription{}, nil
}
