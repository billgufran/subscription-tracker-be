package models

import (
	"time"

	"gorm.io/gorm"
)

type BillingCycle struct {
	ID            ULID   `gorm:"primaryKey;type:char(26)"`
	Name          string `gorm:"not null"`
	Days          int    `gorm:"not null"`
	SystemDefined bool   `gorm:"not null;default:false"`
	UserID        *ULID  `gorm:"type:char(26);index"`
	User          *User  `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Default billing cycles
var DefaultBillingCycles = []BillingCycle{
	{Name: "Weekly", Days: 7, SystemDefined: true},
	{Name: "Monthly", Days: 30, SystemDefined: true},
	{Name: "Quarterly", Days: 90, SystemDefined: true},
	{Name: "Semi-Annual", Days: 180, SystemDefined: true},
	{Name: "Yearly", Days: 365, SystemDefined: true},
}

// CalculateNextBillingDate calculates the next billing date based on the cycle days
func (bc *BillingCycle) CalculateNextBillingDate(from time.Time) time.Time {
	return from.AddDate(0, 0, bc.Days)
}
