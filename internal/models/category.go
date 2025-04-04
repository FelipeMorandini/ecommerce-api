package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null"`
	Description string         `json:"description"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	ImageURL    string         `json:"image_url"`
	ParentID    *int64         `json:"parent_id" gorm:"index"`
	Level       int            `json:"level" gorm:"default:0"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Parent   *Category  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Products []Product  `json:"products,omitempty" gorm:"many2many:product_categories"`
}
