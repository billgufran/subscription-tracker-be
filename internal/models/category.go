package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID            ULID   `gorm:"primaryKey;type:char(26)"`
	Name          string `gorm:"not null"`
	SystemDefined bool   `gorm:"not null;default:false"`
	// Use pointers because system-defined categories don't have a user
	UserID    *ULID `gorm:"type:char(26);index"` // Nullable for system-defined categories
	User      *User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var DefaultCategories = []Category{
	{Name: "Streaming", SystemDefined: true},
	{Name: "Gaming", SystemDefined: true},
	{Name: "Music", SystemDefined: true},
	{Name: "Cloud Storage", SystemDefined: true},
	{Name: "News", SystemDefined: true},
	{Name: "Fitness", SystemDefined: true},
	{Name: "Productivity Tools", SystemDefined: true},
}
