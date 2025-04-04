package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductVariant struct {
	ID             int64          `json:"id" gorm:"primaryKey"`
	ProductID      int64          `json:"product_id" gorm:"index;not null"`
	SKU            string         `json:"sku" gorm:"uniqueIndex;not null"`
	Name           string         `json:"name" gorm:"not null"`
	Price          float64        `json:"price" gorm:"not null"`
	CompareAtPrice *float64       `json:"compare_at_price"`
	Stock          int            `json:"stock" gorm:"not null;default:0"`
	Weight         *float64       `json:"weight"`
	ImageURL       string         `json:"image_url"`
	IsDefault      bool           `json:"is_default" gorm:"default:false"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Product      Product       `json:"-" gorm:"foreignKey:ProductID"`
	OptionValues []OptionValue `json:"option_values,omitempty" gorm:"many2many:variant_option_values"`
}
