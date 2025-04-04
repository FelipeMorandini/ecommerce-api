package models

import (
	"time"

	"gorm.io/gorm"
)

type ShippingMethod struct {
	ID            int64          `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"not null"`
	Description   string         `json:"description"`
	Price         float64        `json:"price" gorm:"not null"`
	EstimatedDays string         `json:"estimated_days"` // e.g., "3-5 days"
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
