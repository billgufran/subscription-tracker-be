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

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) Create(req *CreateCategoryRequest, userID models.ULID) (*models.Category, error) {
	// Check if category already exists for this user or as default
	exists, err := s.categoryRepo.ExistsByNameAndUser(req.Name, userID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("category with name '%s' already exists", req.Name)
	}

	category := &models.Category{
		ID:            models.NewULID(),
		Name:          req.Name,
		UserID:        &userID,
		SystemDefined: false,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetAll(userID models.ULID) ([]models.Category, error) {
	return s.categoryRepo.GetAllForUser(userID)
}

func (s *CategoryService) Update(id models.ULID, req *UpdateCategoryRequest, userID models.ULID) (*models.Category, error) {
	// Get the existing category
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("category not found")
	}

	// Check if it's a default category
	if category.SystemDefined {
		return nil, fmt.Errorf("cannot edit default category")
	}

	// Check if the category belongs to the user
	if category.UserID == nil || *category.UserID != userID {
		return nil, fmt.Errorf("category not found")
	}

	// Check if the new name conflicts with existing categories
	exists, err := s.categoryRepo.ExistsByNameAndUser(req.Name, userID, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("category with name '%s' already exists", req.Name)
	}

	// Update the category
	category.Name = req.Name
	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) Delete(id models.ULID, userID models.ULID) error {
	// Get the existing category
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("category not found")
	}

	// Check if it's a default category
	if category.SystemDefined {
		return fmt.Errorf("cannot delete default category")
	}

	// Check if the category belongs to the user
	if category.UserID == nil || *category.UserID != userID {
		return fmt.Errorf("category not found")
	}

	return s.categoryRepo.Delete(category)
}
