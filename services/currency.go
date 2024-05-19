package services

import (
	"github.com/ownxgv/Genesis_SES_task/models"
	"github.com/ownxgv/Genesis_SES_task/repositories"
	"github.com/ownxgv/Genesis_SES_task/utils"
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
		usdRate, err := utils.FetchUSDRate()
		if err != nil {
			return nil, err
		}

		newRate := &models.CurrencyRate{
			USDRate: usdRate,
		}

		err = s.Repo.SaveCurrencyRate(newRate)
		if err != nil {
			return nil, err
		}

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

func (s *CurrencyService) GetSubscriptions() ([]models.Subscription, error) {
	subscriptions, err := s.Repo.GetSubscriptions()
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}
