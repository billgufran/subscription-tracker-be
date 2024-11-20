package services

import (
	"fmt"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/repository"
)

type PaymentMethodService struct {
	paymentMethodRepo *repository.PaymentMethodRepository
}

type CreatePaymentMethodRequest struct {
	Name     string                   `json:"name" binding:"required"`
	Type     models.PaymentMethodType `json:"type" binding:"required"`
	LastFour string                   `json:"lastFour,omitempty"`
}

type UpdatePaymentMethodRequest struct {
	Name     string                   `json:"name" binding:"required"`
	Type     models.PaymentMethodType `json:"type" binding:"required"`
	LastFour string                   `json:"lastFour" binding:"required,len=4"`
}

func NewPaymentMethodService(paymentMethodRepo *repository.PaymentMethodRepository) *PaymentMethodService {
	return &PaymentMethodService{
		paymentMethodRepo: paymentMethodRepo,
	}
}

func (s *PaymentMethodService) Create(req *CreatePaymentMethodRequest, userID models.ULID) (*models.PaymentMethod, error) {
	if !models.IsValidPaymentMethodType(req.Type) {
		return nil, fmt.Errorf("invalid payment method type: %s", req.Type)
	}

	exists, err := s.paymentMethodRepo.ExistsByNameTypeAndUser(req.Name, req.Type, userID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("payment method with name '%s' and type '%s' already exists", req.Name, req.Type)
	}

	paymentMethod := &models.PaymentMethod{
		UserID:   userID,
		Name:     req.Name,
		Type:     req.Type,
		LastFour: req.LastFour,
	}

	if err := s.paymentMethodRepo.Create(paymentMethod); err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (s *PaymentMethodService) GetAll(userID models.ULID) ([]models.PaymentMethod, error) {
	return s.paymentMethodRepo.GetAllForUser(userID)
}

func (s *PaymentMethodService) Update(id models.ULID, req *UpdatePaymentMethodRequest, userID models.ULID) (*models.PaymentMethod, error) {
	if !models.IsValidPaymentMethodType(req.Type) {
		return nil, fmt.Errorf("invalid payment method type: %s", req.Type)
	}

	paymentMethod, err := s.paymentMethodRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("payment method not found")
	}

	if paymentMethod.UserID != userID {
		return nil, fmt.Errorf("payment method not found")
	}

	exists, err := s.paymentMethodRepo.ExistsByNameTypeAndUser(req.Name, req.Type, userID, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("payment method with name '%s' and type '%s' already exists", req.Name, req.Type)
	}

	paymentMethod.Name = req.Name
	paymentMethod.Type = req.Type
	paymentMethod.LastFour = req.LastFour

	if err := s.paymentMethodRepo.Update(paymentMethod); err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (s *PaymentMethodService) Delete(id models.ULID, userID models.ULID) error {
	paymentMethod, err := s.paymentMethodRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("payment method not found")
	}

	if paymentMethod.UserID != userID {
		return fmt.Errorf("payment method not found")
	}

	return s.paymentMethodRepo.Delete(paymentMethod)
}
