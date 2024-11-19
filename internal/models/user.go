package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint            `gorm:"primaryKey"`
	Email          string          `gorm:"uniqueIndex;not null"`
	PasswordHash   string          `gorm:"not null"`
	Name           string          `gorm:"not null"`
	Categories     []Category      `gorm:"foreignKey:UserID"`
	Subscriptions  []Subscription  `gorm:"foreignKey:UserID"`
	PaymentMethods []PaymentMethod `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
