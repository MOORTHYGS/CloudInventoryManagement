package models

import (
	"time"

	"github.com/google/uuid"
)

type StockTransfer struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId      uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	InventoryId   uuid.UUID `gorm:"type:uuid;not null" json:"inventory_id"`
	FromWarehouse uuid.UUID `gorm:"type:uuid" json:"from_warehouse"`
	ToWarehouse   uuid.UUID `gorm:"type:uuid" json:"to_warehouse"`
	Quantity      int       `gorm:"not null" json:"quantity"`
	TransferDate  time.Time `gorm:"default:now()" json:"transfer_date"`
	CreatedBy     uuid.UUID `gorm:"type:uuid" json:"created_by"`
	Status        string    `gorm:"type:text;default:'initiated'" json:"status"` // initiated/completed/cancelled
}

func (StockTransfer) TableName() string {
	return "stock_transfers"
}
