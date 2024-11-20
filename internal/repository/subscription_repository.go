package repository

import (
	"subscription-tracker/internal/models"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(subscription *models.Subscription) error {
	return r.db.Create(subscription).Error
}

func (r *SubscriptionRepository) GetAll(userID models.ULID) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := r.db.Where("user_id = $1", userID).
		Preload("Category").
		Preload("Currency").
		Preload("BillingCycle").
		Preload("PaymentMethod").
		Find(&subscriptions).Error
	return subscriptions, err
}

func (r *SubscriptionRepository) GetByID(id, userID models.ULID) (*models.Subscription, error) {
	var subscription models.Subscription
	err := r.db.Where("id = $1 AND user_id = $2", id, userID).
		First(&subscription).Error
	return &subscription, err
}

func (r *SubscriptionRepository) GetByCategory(categoryID, userID models.ULID) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := r.db.Where("category_id = $1 AND user_id = $2", categoryID, userID).
		Preload("Category").
		Find(&subscriptions).Error
	return subscriptions, err
}

func (r *SubscriptionRepository) GetByBillingCycle(billingCycleID, userID models.ULID) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := r.db.Where("billing_cycle_id = $1 AND user_id = $2", billingCycleID, userID).
		Preload("BillingCycle").
		Find(&subscriptions).Error
	return subscriptions, err
}

func (r *SubscriptionRepository) GetByPaymentMethod(paymentMethodID, userID models.ULID) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := r.db.Where("payment_method_id = $1 AND user_id = $2", paymentMethodID, userID).
		Preload("PaymentMethod").
		Find(&subscriptions).Error
	return subscriptions, err
}
