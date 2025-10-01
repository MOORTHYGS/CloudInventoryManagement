package models

import (
	"time"

	"github.com/google/uuid"
)

type InventoryStock struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId     uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	InventoryId  uuid.UUID `gorm:"type:uuid;not null" json:"inventory_id"`
	WarehouseId  uuid.UUID `gorm:"type:uuid;not null" json:"warehouse_id"`
	Quantity     int       `gorm:"default:0" json:"quantity"`
	ReorderLevel int       `gorm:"default:0" json:"reorder_level"`
	UpdatedAt    time.Time `gorm:"default:now()" json:"updated_at"`
}

func (InventoryStock) TableName() string {
	return "inventory_stock"
}
