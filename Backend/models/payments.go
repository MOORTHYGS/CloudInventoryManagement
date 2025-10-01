package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId       uuid.UUID  `gorm:"type:uuid;not null" json:"tenant_id"`
	SubscriptionId *uuid.UUID `gorm:"type:uuid" json:"subscription_id"`
	Amount         float64    `gorm:"not null" json:"amount"`
	Method         string     `gorm:"type:text;default:'razorpay'" json:"method"` // razorpay/stripe/paypal/cash
	Status         string     `gorm:"type:text;default:'pending'" json:"status"`  // success/failed/pending
	TransactionID  string     `gorm:"type:text" json:"transaction_id"`
	PaidAt         time.Time  `gorm:"default:now()" json:"paid_at"`
}

func (Payment) TableName() string {
	return "payments"
}
