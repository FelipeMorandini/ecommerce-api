package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductAttribute struct {
	ID           int64          `json:"id" gorm:"primaryKey"`
	ProductID    int64          `json:"product_id" gorm:"index;not null"`
	Name         string         `json:"name" gorm:"not null"`
	Value        string         `json:"value" gorm:"not null"`
	IsFilterable bool           `json:"is_filterable" gorm:"default:false"`
	IsSearchable bool           `json:"is_searchable" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Product Product `json:"-" gorm:"foreignKey:ProductID"`
}
