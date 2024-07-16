package models

import (
	"gorm.io/gorm"
	"time"
)

type Member struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	IdentityNo string    `gorm:"not null" json:"identity_no"`
	Address    string    `gorm:"not null" json:"address"`
	Phone      string    `gorm:"not null" json:"phone"`
	IsActive   bool      `gorm:"not null" json:"is_active"`
	IsDeleted  bool      `gorm:"not null" json:"is_deleted"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
}
