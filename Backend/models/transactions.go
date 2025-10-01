package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId    uuid.UUID  `gorm:"type:uuid;not null" json:"tenant_id"`
	InventoryId uuid.UUID  `gorm:"type:uuid;not null" json:"inventory_id"`
	WarehouseId uuid.UUID  `gorm:"type:uuid;not null" json:"warehouse_id"`
	Type        string     `gorm:"type:text;not null" json:"type"` // IN, OUT, TRANSFER, ADJUST
	Quantity    int        `gorm:"not null" json:"quantity"`
	UnitPrice   float64    `gorm:"default:0" json:"unit_price"`
	TotalPrice  float64    `gorm:"default:0" json:"total_price"`
	SupplierId  *uuid.UUID `gorm:"type:uuid" json:"supplier_id"`
	CustomerId  *uuid.UUID `gorm:"type:uuid" json:"customer_id"`
	Date        time.Time  `gorm:"default:now()" json:"date"`
	Note        string     `gorm:"type:text" json:"note"`
	CreatedBy   uuid.UUID  `gorm:"type:uuid" json:"created_by"`
}

func (Transaction) TableName() string {
	return "transactions"
}
