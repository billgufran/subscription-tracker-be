package services

import (
	"fmt"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) Create(req *CreateCategoryRequest, userID uint) (*models.Category, error) {
	// Check if category already exists for this user or as default
	exists, err := s.categoryRepo.ExistsByNameAndUser(req.Name, userID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("category with name '%s' already exists", req.Name)
	}

	category := &models.Category{
		Name:      req.Name,
		UserID:    &userID,
		IsDefault: false,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}
