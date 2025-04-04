package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductStatus string

const (
	ProductStatusDraft      ProductStatus = "draft"
	ProductStatusActive     ProductStatus = "active"
	ProductStatusInactive   ProductStatus = "inactive"
	ProductStatusOutOfStock ProductStatus = "out_of_stock"
)

type Product struct {
	ID               int64          `json:"id" gorm:"primaryKey"`
	Name             string         `json:"name" gorm:"not null"`
	Description      string         `json:"description"`
	ShortDescription string         `json:"short_description"`
	SKU              string         `json:"sku" gorm:"uniqueIndex;not null"`
	Slug             string         `json:"slug" gorm:"uniqueIndex;not null"`
	Price            float64        `json:"price" gorm:"not null"`
	CompareAtPrice   *float64       `json:"compare_at_price"` // Original price for showing discounts
	CostPrice        *float64       `json:"cost_price"`
	Stock            int            `json:"stock" gorm:"not null;default:0"`
	Weight           *float64       `json:"weight"`
	WeightUnit       string         `json:"weight_unit" gorm:"default:'kg'"`
	Length           *float64       `json:"length"`
	Width            *float64       `json:"width"`
	Height           *float64       `json:"height"`
	DimensionUnit    string         `json:"dimension_unit" gorm:"default:'cm'"`
	Status           ProductStatus  `json:"status" gorm:"type:varchar(20);default:'draft'"`
	FeaturedImageURL string         `json:"featured_image_url"`
	IsDigital        bool           `json:"is_digital" gorm:"default:false"`
	DigitalFileURL   string         `json:"digital_file_url"`
	IsFeatured       bool           `json:"is_featured" gorm:"default:false"`
	IsNew            bool           `json:"is_new" gorm:"default:false"`
	IsBestseller     bool           `json:"is_bestseller" gorm:"default:false"`
	TaxClass         string         `json:"tax_class"`
	MetaTitle        string         `json:"meta_title"`
	MetaDescription  string         `json:"meta_description"`
	MetaKeywords     string         `json:"meta_keywords"`
	CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Categories      []Category         `json:"categories,omitempty" gorm:"many2many:product_categories"`
	Images          []ProductImage     `json:"images,omitempty" gorm:"foreignKey:ProductID"`
	Variants        []ProductVariant   `json:"variants,omitempty" gorm:"foreignKey:ProductID"`
	Attributes      []ProductAttribute `json:"attributes,omitempty" gorm:"foreignKey:ProductID"`
	Reviews         []Review           `json:"reviews,omitempty" gorm:"foreignKey:ProductID"`
	RelatedProducts []Product          `json:"related_products,omitempty" gorm:"many2many:product_related"`
}

// CalculateAverageRating calculates the average rating for the product
func (p *Product) CalculateAverageRating() float64 {
	if len(p.Reviews) == 0 {
		return 0
	}

	var sum int
	for _, review := range p.Reviews {
		sum += review.Rating
	}

	return float64(sum) / float64(len(p.Reviews))
}

// IsOnSale checks if the product is on sale
func (p *Product) IsOnSale() bool {
	return p.CompareAtPrice != nil && *p.CompareAtPrice > p.Price
}

// DiscountPercentage calculates the discount percentage if the product is on sale
func (p *Product) DiscountPercentage() int {
	if !p.IsOnSale() {
		return 0
	}

	discount := (*p.CompareAtPrice - p.Price) / *p.CompareAtPrice * 100
	return int(discount)
}
