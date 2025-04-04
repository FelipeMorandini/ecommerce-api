package models

import (
	"time"

	"gorm.io/gorm"
)

type AddressType string

const (
	AddressTypeBilling  AddressType = "billing"
	AddressTypeShipping AddressType = "shipping"
)

type Address struct {
	ID           int64          `json:"id" gorm:"primaryKey"`
	UserID       int64          `json:"user_id" gorm:"index;not null"`
	Type         AddressType    `json:"type" gorm:"type:varchar(20);not null"`
	IsDefault    bool           `json:"is_default" gorm:"default:false"`
	FirstName    string         `json:"first_name" gorm:"not null"`
	LastName     string         `json:"last_name" gorm:"not null"`
	AddressLine1 string         `json:"address_line1" gorm:"not null"`
	AddressLine2 string         `json:"address_line2"`
	City         string         `json:"city" gorm:"not null"`
	State        string         `json:"state" gorm:"not null"`
	PostalCode   string         `json:"postal_code" gorm:"not null"`
	Country      string         `json:"country" gorm:"not null"`
	Phone        string         `json:"phone" gorm:"size:20"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User User `json:"-" gorm:"foreignKey:UserID"`
}
