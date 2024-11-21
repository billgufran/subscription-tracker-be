package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             ULID            `gorm:"primaryKey;type:char(26)"`
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Email = strings.ToLower(u.Email)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.Email = strings.ToLower(u.Email)
	return nil
}
