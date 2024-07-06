package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Title       string
	Author      string
	Publisher   string
	Category    string
	Year        int
	Image       string
	ISBN        string
	Description string
	IsActive    bool
	IsDeleted   bool
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
