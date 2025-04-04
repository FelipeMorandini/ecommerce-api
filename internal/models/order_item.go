package models

import (
	"time"
)

type OrderItem struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	OrderID   int64     `json:"order_id" gorm:"index;not null"`
	ProductID int64     `json:"product_id" gorm:"index;not null"`
	VariantID *int64    `json:"variant_id" gorm:"index"`
	Name      string    `json:"name" gorm:"not null"`  // Product name at time of order
	SKU       string    `json:"sku" gorm:"not null"`   // Product SKU at time of order
	Price     float64   `json:"price" gorm:"not null"` // Price at time of order
	Quantity  int       `json:"quantity" gorm:"not null"`
	Subtotal  float64   `json:"subtotal" gorm:"not null"`
	Options   string    `json:"options"` // JSON string of selected options
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Order   Order           `json:"-" gorm:"foreignKey:OrderID"`
	Product Product         `json:"product" gorm:"foreignKey:ProductID"`
	Variant *ProductVariant `json:"variant,omitempty" gorm:"foreignKey:VariantID"`
}
