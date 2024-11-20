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
	PaymentMethodTypeOther         PaymentMethodType = "other"
)

func IsValidPaymentMethodType(pmType PaymentMethodType) bool {
	switch pmType {
	case PaymentMethodTypeCreditCard,
		PaymentMethodTypeDebitCard,
		PaymentMethodTypeBankAccount,
		PaymentMethodTypeDigitalWallet,
		PaymentMethodTypeOther:
		return true
	}
	return false
}

type PaymentMethod struct {
	ID        ULID              `gorm:"primaryKey;type:char(26)"`
	UserID    ULID              `gorm:"type:char(26);not null"`
	User      User              `gorm:"foreignKey:UserID"`
	Name      string            `gorm:"not null"`
	Type      PaymentMethodType `gorm:"not null;type:varchar(20)"`
	LastFour  string            `gorm:"type:varchar(4)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
