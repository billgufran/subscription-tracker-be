package repository

import (
	"subscription-tracker/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) ExistsByNameAndUser(name string, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Category{}).
		Where("name = $1 AND (user_id = $2 OR is_default = $3)", name, userID, true).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
