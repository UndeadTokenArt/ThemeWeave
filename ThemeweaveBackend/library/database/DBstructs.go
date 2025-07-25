package database

import (
	"time"
)

// User struct to match the database table
type Website struct {
	ID            uint     `gorm:"primaryKey"`
	UserID        uint     `gorm:"not null"` // Foreign key to User
	Name          string   `gorm:"not null"`
	Theme         string   `gorm:"not null;default:'default'"`
	ColorScheme   string   `gorm:"not null;default:'#ffffff,#000000,#3498db'"` // Primary,Secondary,Accent
	MainBody      []string `gorm:"type:text"`
	CallToAction  []string `gorm:"type:text"`
	HeaderContent string   `gorm:"type:text"`
	FooterContent string   `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
