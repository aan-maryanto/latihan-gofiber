package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null; unique"`
	Password  string    `gorm:"not null"`
	Salt      string    `gorm:"not null"`
	IsActive  bool      `gorm:"default:true"`
	IsDeleted bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string
}
