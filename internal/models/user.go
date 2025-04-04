package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
	RoleSeller   Role = "seller"
)

type User struct {
	ID              int64          `json:"id" gorm:"primaryKey"`
	Email           string         `json:"email" gorm:"uniqueIndex;not null"`
	Password        string         `json:"-" gorm:"not null"` // Never expose in JSON
	FirstName       string         `json:"first_name" gorm:"not null"`
	LastName        string         `json:"last_name" gorm:"not null"`
	Role            Role           `json:"role" gorm:"type:varchar(20);default:'customer'"`
	Phone           string         `json:"phone" gorm:"size:20"`
	ProfileImageURL string         `json:"profile_image_url"`
	Verified        bool           `json:"verified" gorm:"default:false"`
	LastLogin       *time.Time     `json:"last_login"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Addresses []Address  `json:"addresses,omitempty" gorm:"foreignKey:UserID"`
	Orders    []Order    `json:"orders,omitempty" gorm:"foreignKey:UserID"`
	Reviews   []Review   `json:"reviews,omitempty" gorm:"foreignKey:UserID"`
	Wishlist  []Product  `json:"wishlist,omitempty" gorm:"many2many:user_wishlist_items"`
	Cart      []CartItem `json:"cart,omitempty" gorm:"foreignKey:UserID"`
}

// BeforeSave is a GORM hook that hashes the password before saving
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Only hash the password if it has been changed
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// ComparePassword checks if the provided password matches the stored hash
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
