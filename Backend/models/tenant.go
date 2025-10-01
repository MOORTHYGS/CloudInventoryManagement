package models

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name       string    `gorm:"type:text;not null" json:"name"`
	Domain     string    `gorm:"type:text" json:"domain"`
	Status     string    `gorm:"type:text;default:'active'" json:"status"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	CreatedAt  time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:now()" json:"updated_at"`
}

func (Tenant) TableName() string {
	return "tenants"
}
