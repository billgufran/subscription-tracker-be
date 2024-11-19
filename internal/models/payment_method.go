package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Name      string `gorm:"not null"`
	Type      string `gorm:"not null"`
	LastFour  string `gorm:"type:varchar(4)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
