package services

import (
	"github.com/ownxgv/Genesis_SES_task/models"
	"github.com/ownxgv/Genesis_SES_task/repositories"
	"github.com/ownxgv/Genesis_SES_task/utils"
)

type CurrencyService struct {
	Repo *repositories.CurrencyRepository
}

func NewCurrencyService(repo *repositories.CurrencyRepository) *CurrencyService {
	return &CurrencyService{Repo: repo}
}

func (s *CurrencyService) GetCurrencyRate() (*models.CurrencyRate, error) {
	rate, err := utils.FetchUSDRate()
	if err != nil {
		return nil, err
	}

	currencyRate := &models.CurrencyRate{
		CurrencyCode: "USD",
		Rate:         rate,
	}

	err = s.Repo.SaveCurrencyRate(currencyRate)
	if err != nil {
		return nil, err
	}

	return currencyRate, nil
}

func (s *CurrencyService) SubscribeEmail(email string) error {
	rate, err := s.Repo.GetCurrencyRate()
	if err != nil {
		return err
	}

	err = utils.SendEmail(email, rate.Rate)
	if err != nil {
		return err
	}

	return s.Repo.SubscribeEmail(email)
}

func (s *CurrencyService) SendDailyRates() error {
	// Ваша логика для отправки ежедневных курсов валют
	// Возможно, вам нужно получить курсы валют и отправить их на почту или сохранить в базу данных
	return nil
}
