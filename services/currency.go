package services

import (
	"github.com/your-username/currency-service/models"
	"github.com/your-username/currency-service/repositories"
	"github.com/your-username/currency-service/utils"
)

type CurrencyService struct {
	Repo repositories.CurrencyRepository
}

func NewCurrencyService(repo repositories.CurrencyRepository) CurrencyService {
	return CurrencyService{Repo: repo}
}

func (s *CurrencyService) GetCurrencyRate() (*models.CurrencyRate, error) {
	rate, err := s.Repo.GetCurrencyRate()
	if err != nil {
		return nil, err
	}

	if rate.ID == 0 {
		newRate, err := utils.FetchCurrencyRate()
		if err != nil {
			return nil, err
		}

		err = s.Repo.SaveCurrencyRate(newRate)
		if err != nil {
			return nil, err
			rate = newRate
		}

		return rate, nil
	}

	func (s *CurrencyService) SubscribeEmail(email string) error {
		return s.Repo.SubscribeEmail(email)
	}

	func (s *CurrencyService) SendDailyRates() error {
		subscriptions, err := s.Repo.GetSubscriptions()
		if err != nil {
		return err
	}

		rate, err := s.GetCurrencyRate()
		if err != nil {
		return err
	}

		for _, sub := range subscriptions {
		err = utils.SendEmail(sub.Email, rate.USDRate)
		if err != nil {
		return err
	}
	}

		return nil
	}