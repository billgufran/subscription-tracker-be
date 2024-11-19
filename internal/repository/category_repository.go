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

func (r *CategoryRepository) ExistsByNameAndUser(name string, userID models.ULID, excludeID *models.ULID) (bool, error) {
	query := r.db.Model(&models.Category{}).
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

func (r *CategoryRepository) GetAllForUser(userID models.ULID) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("user_id = $1 OR system_defined = $2", userID, true).
		Order("system_defined DESC, name ASC").
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id models.ULID) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = $1", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(category *models.Category) error {
	return r.db.Delete(category).Error
}
