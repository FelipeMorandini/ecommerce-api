package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID         int64          `json:"id" gorm:"primaryKey"`
	UserID     int64          `json:"user_id" gorm:"index;not null"`
	ProductID  int64          `json:"product_id" gorm:"index;not null"`
	OrderID    *int64         `json:"order_id" gorm:"index"` // Optional link to order
	Rating     int            `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Status     string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	IsVerified bool           `json:"is_verified" gorm:"default:false"` // Verified purchase
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Product Product `json:"-" gorm:"foreignKey:ProductID"`
	Order   *Order  `json:"-" gorm:"foreignKey:OrderID"`
}
