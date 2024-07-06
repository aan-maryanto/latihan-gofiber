package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string
	Password  string
	Salt      string
	IsActive  bool
	IsDeleted bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
