package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	ProductID int64          `json:"product_id" gorm:"index;not null"`
	URL       string         `json:"url" gorm:"not null"`
	AltText   string         `json:"alt_text"`
	SortOrder int            `json:"sort_order" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Product Product `json:"-" gorm:"foreignKey:ProductID"`
}
