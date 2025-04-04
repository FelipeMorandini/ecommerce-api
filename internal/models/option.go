package models

import (
	"time"

	"gorm.io/gorm"
)

type Option struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	DisplayName string         `json:"display_name" gorm:"not null"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Values []OptionValue `json:"values,omitempty" gorm:"foreignKey:OptionID"`
}

type OptionValue struct {
	ID           int64          `json:"id" gorm:"primaryKey"`
	OptionID     int64          `json:"option_id" gorm:"index;not null"`
	Value        string         `json:"value" gorm:"not null"`
	DisplayValue string         `json:"display_value" gorm:"not null"`
	SortOrder    int            `json:"sort_order" gorm:"default:0"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Option   Option           `json:"-" gorm:"foreignKey:OptionID"`
	Variants []ProductVariant `json:"variants,omitempty" gorm:"many2many:variant_option_values"`
}
