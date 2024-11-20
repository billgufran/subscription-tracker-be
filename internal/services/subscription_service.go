package services

import (
	"fmt"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/repository"
	"time"
)

type SubscriptionService struct {
	subscriptionRepo *repository.SubscriptionRepository
	categoryRepo     *repository.CategoryRepository
	currencyRepo     *repository.CurrencyRepository
	billingCycleRepo *repository.BillingCycleRepository
}

type CreateSubscriptionRequest struct {
	Name            string    `json:"name" binding:"required"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount" binding:"required,gt=0"`
	CategoryID      string    `json:"categoryId" binding:"required"`
	CurrencyID      string    `json:"currencyId" binding:"required"`
	BillingCycleID  string    `json:"billingCycleId" binding:"required"`
	PaymentMethodID string    `json:"paymentMethodId" binding:"required"`
	NextBillingDate time.Time `json:"nextBillingDate" binding:"required"`
	ReminderDays    int       `json:"reminderDays" binding:"gte=0"`
}

func NewSubscriptionService(
	subscriptionRepo *repository.SubscriptionRepository,
	categoryRepo *repository.CategoryRepository,
	currencyRepo *repository.CurrencyRepository,
	billingCycleRepo *repository.BillingCycleRepository,
) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
		categoryRepo:     categoryRepo,
		currencyRepo:     currencyRepo,
		billingCycleRepo: billingCycleRepo,
	}
}

func (s *SubscriptionService) Create(req *CreateSubscriptionRequest, userID models.ULID) (*models.Subscription, error) {
	// Parse IDs
	var categoryID, currencyID, billingCycleID, paymentMethodID models.ULID
	if err := categoryID.UnmarshalJSON([]byte(`"` + req.CategoryID + `"`)); err != nil {
		return nil, fmt.Errorf("invalid category ID")
	}
	if err := currencyID.UnmarshalJSON([]byte(`"` + req.CurrencyID + `"`)); err != nil {
		return nil, fmt.Errorf("invalid currency ID")
	}
	if err := billingCycleID.UnmarshalJSON([]byte(`"` + req.BillingCycleID + `"`)); err != nil {
		return nil, fmt.Errorf("invalid billing cycle ID")
	}
	if err := paymentMethodID.UnmarshalJSON([]byte(`"` + req.PaymentMethodID + `"`)); err != nil {
		return nil, fmt.Errorf("invalid payment method ID")
	}

	subscription := &models.Subscription{
		ID:              models.NewULID(),
		UserID:          userID,
		Name:            req.Name,
		Description:     req.Description,
		Amount:          req.Amount,
		CategoryID:      categoryID,
		CurrencyID:      currencyID,
		BillingCycleID:  billingCycleID,
		PaymentMethodID: paymentMethodID,
		NextBillingDate: req.NextBillingDate,
		ReminderDays:    req.ReminderDays,
		Active:          true,
	}

	if err := s.subscriptionRepo.Create(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}
