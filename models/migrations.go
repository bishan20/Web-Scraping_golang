package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	UserName  string    `gorm:"unique; not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}

type Scrape struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	ScrappingUrl  string    `gorm:"not null; index"`
	ScrappingData string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"index; default:now()"`
}
