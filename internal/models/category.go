package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	IsDefault bool   `gorm:"default:false"`
	UserID    *uint  `gorm:"index"` // Nullable for default categories
	User      *User  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var DefaultCategories = []Category{
	{Name: "Streaming", IsDefault: true},
	{Name: "Software", IsDefault: true},
	{Name: "Gaming", IsDefault: true},
	{Name: "Music", IsDefault: true},
	{Name: "Cloud Storage", IsDefault: true},
	{Name: "News", IsDefault: true},
	{Name: "Fitness", IsDefault: true},
	{Name: "Productivity Tools", IsDefault: true},
}
