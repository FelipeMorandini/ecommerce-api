package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
)

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard     PaymentMethod = "credit_card"
	PaymentMethodPayPal         PaymentMethod = "paypal"
	PaymentMethodBankTransfer   PaymentMethod = "bank_transfer"
	PaymentMethodCashOnDelivery PaymentMethod = "cash_on_delivery"
)

type Order struct {
	ID                 int64          `json:"id" gorm:"primaryKey"`
	UserID             int64          `json:"user_id" gorm:"index;not null"`
	OrderNumber        string         `json:"order_number" gorm:"uniqueIndex;not null"`
	Status             OrderStatus    `json:"status" gorm:"type:varchar(20);default:'pending'"`
	PaymentStatus      PaymentStatus  `json:"payment_status" gorm:"type:varchar(20);default:'pending'"`
	PaymentMethod      PaymentMethod  `json:"payment_method" gorm:"type:varchar(20)"`
	PaymentID          string         `json:"payment_id"` // External payment ID (e.g., from PayPal)
	Currency           string         `json:"currency" gorm:"size:3;default:'USD'"`
	Subtotal           float64        `json:"subtotal" gorm:"not null"`
	ShippingCost       float64        `json:"shipping_cost" gorm:"not null;default:0"`
	TaxAmount          float64        `json:"tax_amount" gorm:"not null;default:0"`
	DiscountAmount     float64        `json:"discount_amount" gorm:"not null;default:0"`
	CouponCode         string         `json:"coupon_code"`
	TotalAmount        float64        `json:"total_amount" gorm:"not null"`
	Notes              string         `json:"notes"`
	ShippingAddressID  int64          `json:"shipping_address_id" gorm:"index"`
	BillingAddressID   int64          `json:"billing_address_id" gorm:"index"`
	ShippingMethod     string         `json:"shipping_method"`
	TrackingNumber     string         `json:"tracking_number"`
	EstimatedDelivery  *time.Time     `json:"estimated_delivery"`
	ShippedAt          *time.Time     `json:"shipped_at"`
	DeliveredAt        *time.Time     `json:"delivered_at"`
	CancelledAt        *time.Time     `json:"cancelled_at"`
	CancellationReason string         `json:"cancellation_reason"`
	RefundedAt         *time.Time     `json:"refunded_at"`
	RefundAmount       *float64       `json:"refund_amount"`
	IPAddress          string         `json:"ip_address"`
	UserAgent          string         `json:"user_agent"`
	CreatedAt          time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User            User          `json:"-" gorm:"foreignKey:UserID"`
	Items           []OrderItem   `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	ShippingAddress Address       `json:"shipping_address" gorm:"foreignKey:ShippingAddressID"`
	BillingAddress  Address       `json:"billing_address" gorm:"foreignKey:BillingAddressID"`
	Transactions    []Transaction `json:"transactions,omitempty" gorm:"foreignKey:OrderID"`
}

// CalculateTotal calculates the total amount for the order
func (o *Order) CalculateTotal() float64 {
	return o.Subtotal + o.ShippingCost + o.TaxAmount - o.DiscountAmount
}

// CalculateSubtotal calculates the subtotal from order items
func (o *Order) CalculateSubtotal() float64 {
	var subtotal float64
	for _, item := range o.Items {
		subtotal += item.Price * float64(item.Quantity)
	}
	return subtotal
}
