package repository

import (
	"subscription-tracker/internal/models"

	"gorm.io/gorm"
)

type BillingCycleRepository struct {
	db *gorm.DB
}

func NewBillingCycleRepository(db *gorm.DB) *BillingCycleRepository {
	return &BillingCycleRepository{db: db}
}

func (r *BillingCycleRepository) Create(billingCycle *models.BillingCycle) error {
	return r.db.Create(billingCycle).Error
}

func (r *BillingCycleRepository) GetByID(id models.ULID) (*models.BillingCycle, error) {
	var billingCycle models.BillingCycle
	err := r.db.Where("id = $1", id).First(&billingCycle).Error
	if err != nil {
		return nil, err
	}
	return &billingCycle, nil
}

func (r *BillingCycleRepository) GetAllForUser(userID models.ULID) ([]models.BillingCycle, error) {
	var billingCycles []models.BillingCycle
	err := r.db.Where("user_id = $1 OR system_defined = $2", userID, true).
		Order("system_defined DESC, name ASC").
		Find(&billingCycles).Error
	if err != nil {
		return nil, err
	}
	return billingCycles, nil
}

func (r *BillingCycleRepository) Update(billingCycle *models.BillingCycle) error {
	return r.db.Save(billingCycle).Error
}

func (r *BillingCycleRepository) Delete(billingCycle *models.BillingCycle) error {
	return r.db.Delete(billingCycle).Error
}

func (r *BillingCycleRepository) ExistsByNameAndUser(name string, userID models.ULID, excludeID *models.ULID) (bool, error) {
	query := r.db.Model(&models.BillingCycle{}).
		Where("name = $1 AND (user_id = $2 OR system_defined = $3)", name, userID, true)

	if excludeID != nil {
		query = query.Where("id != $4", *excludeID)
	}

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
