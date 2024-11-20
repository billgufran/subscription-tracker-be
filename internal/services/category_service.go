package services

import (
	"fmt"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/repository"
	"subscription-tracker/internal/utils"

	"gorm.io/gorm"
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
	exists, err := s.categoryRepo.ExistsByNameAndUser(req.Name, userID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, utils.NewValidationError("name", fmt.Sprintf("category with name '%s' already exists", req.Name))
	}

	category := &models.Category{
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
	category, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if category.SystemDefined {
		return nil, utils.NewForbiddenError("system-defined categories cannot be modified")
	}

	exists, err := s.categoryRepo.ExistsByNameAndUser(req.Name, userID, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, utils.NewValidationError("name", fmt.Sprintf("category with name '%s' already exists", req.Name))
	}

	category.Name = req.Name
	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) Delete(id models.ULID, userID models.ULID) error {
	category, err := s.GetByID(id, userID)
	if err != nil {
		return err
	}

	if category.SystemDefined {
		return utils.NewForbiddenError("system-defined categories cannot be deleted")
	}

	return s.categoryRepo.Delete(category)
}

func (s *CategoryService) GetByID(id, userID models.ULID) (*models.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewNotFoundError("category")
		}
		return nil, err
	}

	if !category.SystemDefined && (category.UserID == nil || *category.UserID != userID) {
		return nil, utils.NewForbiddenError("category does not belong to user")
	}

	return category, nil
}
