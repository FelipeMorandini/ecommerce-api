package database

import (
	"ecommerce-api/internal/models"
	"log"

	"gorm.io/gorm"
)

// AutoMigrate automatically migrates the database schema
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Add all your models here
	err := db.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductVariant{},
		&models.ProductAttribute{},
		&models.Option{},
		&models.OptionValue{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
		&models.Review{},
		&models.Coupon{},
		&models.ShippingMethod{},
	)

	if err != nil {
		log.Printf("Error during migration: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
