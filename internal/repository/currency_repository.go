package repository

import (
	"subscription-tracker/internal/models"

	"gorm.io/gorm"
)

type CurrencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (r *CurrencyRepository) GetAll() ([]models.Currency, error) {
	var currencies []models.Currency
	err := r.db.Order("code ASC").Find(&currencies).Error
	if err != nil {
		return nil, err
	}
	return currencies, nil
}
