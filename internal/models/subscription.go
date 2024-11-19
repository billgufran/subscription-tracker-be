package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID              ULID          `gorm:"primaryKey;type:char(26)"`
	UserID          ULID          `gorm:"type:char(26);not null"`
	User            User          `gorm:"foreignKey:UserID"`
	CategoryID      ULID          `gorm:"type:char(26);not null"`
	Category        Category      `gorm:"foreignKey:CategoryID"`
	PaymentMethodID ULID          `gorm:"type:char(26);not null"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	CurrencyID      ULID          `gorm:"type:char(26);not null"`
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
