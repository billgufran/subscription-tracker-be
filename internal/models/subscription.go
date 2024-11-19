package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID              uint          `gorm:"primaryKey"`
	UserID          uint          `gorm:"not null"`
	User            User          `gorm:"foreignKey:UserID"`
	CategoryID      uint          `gorm:"not null"`
	Category        Category      `gorm:"foreignKey:CategoryID"`
	PaymentMethodID uint          `gorm:"not null"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	CurrencyID      uint          `gorm:"not null"`
	Currency        Currency      `gorm:"foreignKey:CurrencyID"`
	Name            string        `gorm:"not null"`
	Description     string
	Amount          float64   `gorm:"type:decimal(10,2);not null"`
	BillingCycle    string    `gorm:"not null"`
	NextBillingDate time.Time `gorm:"not null"`
	ReminderDays    int       `gorm:"default:7"`
	Active          bool      `gorm:"default:true"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
