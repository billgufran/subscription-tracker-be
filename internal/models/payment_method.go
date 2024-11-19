package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard    PaymentMethodType = "credit_card"
	PaymentMethodTypeDebitCard     PaymentMethodType = "debit_card"
	PaymentMethodTypeBankAccount   PaymentMethodType = "bank_account"
	PaymentMethodTypeDigitalWallet PaymentMethodType = "digital_wallet"
)

type PaymentMethod struct {
	ID        uint              `gorm:"primaryKey"`
	UserID    uint              `gorm:"not null"`
	User      User              `gorm:"foreignKey:UserID"`
	Name      string            `gorm:"not null"`
	Type      PaymentMethodType `gorm:"not null;type:varchar(20)"`
	LastFour  string            `gorm:"type:varchar(4)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
