package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID string `validate:"required"`
	//OrderDate   time.Time
	AddressID string
	//Address     Address `gorm:"foriegnkey:AddressID;references:ID"`
	TotalAmount float64
	// OrderStatus string `gorm:"type:varchar(10); check(order_status IN ('pending', 'delivered', 'cancelled')) ;default:'pending'" json:"order_status" validate:"required"`
	PaymentMethod string `gorm:"type:varchar(10); check(order_status IN ('COD', 'RazorPay')) ;default:'COD'" json:"payment_method" validate:"required"`
	OrderStatus   string `gorm:"type:varchar(10);check:order_status IN ('pending','shipped', 'delivered', 'cancelled','failed');default:'pending'" json:"order_status" validate:"required,oneof=pending delivered shipped cancelled failed"`
	PaymentStatus string `gorm:"type:varchar(10);check:order_status IN ('paid','pending', 'failed','refunded');default:'pending'" json:"payment_status" validate:"required,oneof=paid pending failed refunded"`
	//OfferApplied   float64 `gorm:"default:0.00"`
	//CouponCode     string
	//DiscountAmount float64 `gorm:"type:decimal(10,2);default:0.00"`
	//FinalAmount    float64 `gorm:"type:decimal(10,2);not null"`
}
