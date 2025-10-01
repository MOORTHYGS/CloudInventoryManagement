package models

import (
	"time"

	"github.com/google/uuid"
)

type InventoryItem struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId    uuid.UUID  `gorm:"type:uuid;not null" json:"tenant_id"`
	CategoryId  *uuid.UUID `gorm:"type:uuid" json:"category_id,omitempty"`
	SkuID       uint16     `gorm:"not null" json:"sku_id"`
	Name        string     `gorm:"type:text;not null" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Unit        string     `gorm:"type:text;default:'pcs'" json:"unit"`
	Price       float64    `gorm:"type:numeric;default:0" json:"price"`
	CreatedAt   time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:now()" json:"updated_at"`
}

func (InventoryItem) TableName() string {
	return "inventory_items"
}
