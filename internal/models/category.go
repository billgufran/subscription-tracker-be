package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        ULID   `gorm:"primaryKey;type:char(26)"`
	Name      string `gorm:"not null"`
	IsDefault bool   `gorm:"default:false"`
	// Use pointers because default categories don't have a user
	UserID    *ULID `gorm:"type:char(26);index"` // Nullable for default categories
	User      *User `gorm:"foreignKey:UserID"`
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
