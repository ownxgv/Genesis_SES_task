// mocks/mock_currency_service.go

package mocks

import (
	"github.com/ownxgv/Genesis_SES_task/models"
	"github.com/stretchr/testify/mock"
)

type MockCurrencyService struct {
	mock.Mock
}

func (m *MockCurrencyService) GetCurrencyRate() (*models.CurrencyRate, error) {
	args := m.Called()
	return args.Get(0).(*models.CurrencyRate), args.Error(1)
}

func (m *MockCurrencyService) SubscribeEmail(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (m *MockCurrencyService) SendDailyRates() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCurrencyService) GetSubscriptions() ([]models.Subscription, error) {
	args := m.Called()
	return args.Get(0).([]models.Subscription), args.Error(1)
}
