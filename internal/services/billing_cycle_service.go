package services

import (
	"fmt"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/repository"
	"subscription-tracker/internal/utils"

	"gorm.io/gorm"
)

type BillingCycleService struct {
	billingCycleRepo *repository.BillingCycleRepository
}

type CreateBillingCycleRequest struct {
	Name string `json:"name" binding:"required"`
	Days int    `json:"days" binding:"required,min=1"`
}

type UpdateBillingCycleRequest struct {
	Name string `json:"name" binding:"required"`
	Days int    `json:"days" binding:"required,min=1"`
}

func NewBillingCycleService(billingCycleRepo *repository.BillingCycleRepository) *BillingCycleService {
	return &BillingCycleService{
		billingCycleRepo: billingCycleRepo,
	}
}

func (s *BillingCycleService) Create(req *CreateBillingCycleRequest, userID models.ULID) (*models.BillingCycle, error) {
	exists, err := s.billingCycleRepo.ExistsByNameAndUser(req.Name, userID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, utils.NewValidationError("name", fmt.Sprintf("billing cycle with name '%s' already exists", req.Name))
	}

	billingCycle := &models.BillingCycle{
		Name:          req.Name,
		Days:          req.Days,
		UserID:        &userID,
		SystemDefined: false,
	}

	if err := s.billingCycleRepo.Create(billingCycle); err != nil {
		return nil, err
	}

	return billingCycle, nil
}

func (s *BillingCycleService) GetAll(userID models.ULID) ([]models.BillingCycle, error) {
	return s.billingCycleRepo.GetAllForUser(userID)
}

func (s *BillingCycleService) GetByID(id, userID models.ULID) (*models.BillingCycle, error) {
	billingCycle, err := s.billingCycleRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewNotFoundError("billing cycle")
		}
		return nil, err
	}

	if !billingCycle.SystemDefined && (billingCycle.UserID == nil || *billingCycle.UserID != userID) {
		return nil, utils.NewForbiddenError("billing cycle does not belong to user")
	}

	return billingCycle, nil
}

func (s *BillingCycleService) Update(id models.ULID, req *UpdateBillingCycleRequest, userID models.ULID) (*models.BillingCycle, error) {
	billingCycle, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if billingCycle.SystemDefined {
		return nil, utils.NewForbiddenError("system-defined billing cycles cannot be modified")
	}

	exists, err := s.billingCycleRepo.ExistsByNameAndUser(req.Name, userID, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("billing cycle with name '%s' already exists", req.Name)
	}

	billingCycle.Name = req.Name
	billingCycle.Days = req.Days
	if err := s.billingCycleRepo.Update(billingCycle); err != nil {
		return nil, err
	}

	return billingCycle, nil
}

func (s *BillingCycleService) Delete(id models.ULID, userID models.ULID) error {
	billingCycle, err := s.billingCycleRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("billing cycle not found")
	}

	if billingCycle.SystemDefined {
		return fmt.Errorf("cannot delete system-defined billing cycle")
	}

	if billingCycle.UserID == nil || *billingCycle.UserID != userID {
		return fmt.Errorf("billing cycle not found")
	}

	return s.billingCycleRepo.Delete(billingCycle)
}
