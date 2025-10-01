package models

import (
	"time"

	"github.com/google/uuid"
)

// Sku represents the sku table in Postgres
type Sku struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"` // internal ID
	TenantID   uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	SkuID      int       `gorm:"unique;autoIncrement" json:"sku_id"` // business SKU number
	SkuDesc    string    `gorm:"type:text" json:"sku_desc"`          // description of SKU
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	LastEditAt time.Time `gorm:"autoUpdateTime" json:"last_edit_at"`
	LastEditBy string    `gorm:"type:text" json:"last_edit_by"`
}

// TableName overrides default naming
func (Sku) TableName() string {
	return "sku"
}
