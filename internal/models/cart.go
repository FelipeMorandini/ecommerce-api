package models

import (
	"time"
)

type CartItem struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	UserID    int64     `json:"user_id" gorm:"index;not null"`
	ProductID int64     `json:"product_id" gorm:"index;not null"`
	VariantID *int64    `json:"variant_id" gorm:"index"`
	Quantity  int       `json:"quantity" gorm:"not null;default:1"`
	Price     float64   `json:"price" gorm:"not null"` // Price at the time of adding to cart
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User    User            `json:"-" gorm:"foreignKey:UserID"`
	Product Product         `json:"product" gorm:"foreignKey:ProductID"`
	Variant *ProductVariant `json:"variant,omitempty" gorm:"foreignKey:VariantID"`
}

// CalculateSubtotal calculates the subtotal for this cart item
func (ci *CartItem) CalculateSubtotal() float64 {
	return ci.Price * float64(ci.Quantity)
}
