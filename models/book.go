package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"size:100"`
	Author      string    `gorm:"size:100"`
	Publisher   string    `gorm:"size:100"`
	Category    string    `gorm:"size:255"`
	Year        int       `gorm:"default:0"`
	Image       string    `gorm:"size:255"`
	ISBN        string    `gorm:"size:50"`
	Description string    `gorm:"size:255"`
	IsActive    bool      `gorm:"default:true"`
	IsDeleted   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	CreatedBy   string
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	UpdatedBy   string
}
