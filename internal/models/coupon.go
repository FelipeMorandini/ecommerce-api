package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponType string

const (
	CouponTypePercentage   CouponType = "percentage"
	CouponTypeFixed        CouponType = "fixed"
	CouponTypeFreeShipping CouponType = "free_shipping"
)

type Coupon struct {
	ID                 int64          `json:"id" gorm:"primaryKey"`
	Code               string         `json:"code" gorm:"uniqueIndex;not null"`
	Type               CouponType     `json:"type" gorm:"type:varchar(20);not null"`
	Value              float64        `json:"value"` // Percentage or fixed amount
	MinimumSpend       *float64       `json:"minimum_spend"`
	MaximumDiscount    *float64       `json:"maximum_discount"`
	Description        string         `json:"description"`
	StartDate          *time.Time     `json:"start_date"`
	EndDate            *time.Time     `json:"end_date"`
	UsageLimit         *int           `json:"usage_limit"`                  // Total number of times this coupon can be used
	UsageCount         int            `json:"usage_count" gorm:"default:0"` // Number of times this coupon has been used
	PerUserLimit       *int           `json:"per_user_limit"`               // Number of times a single user can use this coupon
	IsActive           bool           `json:"is_active" gorm:"default:true"`
	AppliesTo          string         `json:"applies_to"`          // all, categories, products (JSON array of IDs)
	ExcludedProducts   string         `json:"excluded_products"`   // JSON array of product IDs
	ExcludedCategories string         `json:"excluded_categories"` // JSON array of category IDs
	CreatedAt          time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}

// IsValid checks if the coupon is currently valid
func (c *Coupon) IsValid() bool {
	now := time.Now()

	if !c.IsActive {
		return false
	}

	if c.StartDate != nil && now.Before(*c.StartDate) {
		return false
	}

	if c.EndDate != nil && now.After(*c.EndDate) {
		return false
	}

	if c.UsageLimit != nil && c.UsageCount >= *c.UsageLimit {
		return false
	}

	return true
}
