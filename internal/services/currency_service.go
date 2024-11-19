package services

import (
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/repository"
)

type CurrencyService struct {
	currencyRepo *repository.CurrencyRepository
}

func NewCurrencyService(currencyRepo *repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{
		currencyRepo: currencyRepo,
	}
}

func (s *CurrencyService) GetAll() ([]models.Currency, error) {
	return s.currencyRepo.GetAll()
}
