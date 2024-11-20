package repository

import (
	"subscription-tracker/internal/models"

	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

func (r *PaymentMethodRepository) Create(paymentMethod *models.PaymentMethod) error {
	return r.db.Create(paymentMethod).Error
}

func (r *PaymentMethodRepository) GetByID(id models.ULID) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := r.db.Where("id = $1", id).First(&paymentMethod).Error
	if err != nil {
		return nil, err
	}
	return &paymentMethod, nil
}

func (r *PaymentMethodRepository) GetAllForUser(userID models.ULID) ([]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod
	err := r.db.Where("user_id = $1", userID).
		Order("name ASC").
		Find(&paymentMethods).Error
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

func (r *PaymentMethodRepository) Update(paymentMethod *models.PaymentMethod) error {
	return r.db.Save(paymentMethod).Error
}

func (r *PaymentMethodRepository) Delete(paymentMethod *models.PaymentMethod) error {
	return r.db.Delete(paymentMethod).Error
}

func (r *PaymentMethodRepository) ExistsByNameTypeAndUser(name string, pmType models.PaymentMethodType, userID models.ULID, excludeID *models.ULID) (bool, error) {
	var count int64
	query := r.db.Model(&models.PaymentMethod{}).
		Where("name = $1 AND type = $2 AND user_id = $3", name, pmType, userID)

	if excludeID != nil {
		query = query.Where("id != $4", excludeID)
	}

	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
