package models

import (
	"time"
)

type TransactionType string

const (
	TransactionTypePayment TransactionType = "payment"
	TransactionTypeRefund  TransactionType = "refund"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID              int64             `json:"id" gorm:"primaryKey"`
	OrderID         int64             `json:"order_id" gorm:"index;not null"`
	Type            TransactionType   `json:"type" gorm:"type:varchar(20);not null"`
	Status          TransactionStatus `json:"status" gorm:"type:varchar(20);not null"`
	Amount          float64           `json:"amount" gorm:"not null"`
	Currency        string            `json:"currency" gorm:"size:3;default:'USD'"`
	PaymentMethod   PaymentMethod     `json:"payment_method" gorm:"type:varchar(20)"`
	TransactionID   string            `json:"transaction_id"` // External transaction ID
	PaymentGateway  string            `json:"payment_gateway"`
	GatewayResponse string            `json:"gateway_response"` // JSON response from payment gateway
	ErrorMessage    string            `json:"error_message"`
	CreatedAt       time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time         `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Order Order `json:"-" gorm:"foreignKey:OrderID"`
}
