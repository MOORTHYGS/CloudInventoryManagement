package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Warehouse struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId    uuid.UUID      `gorm:"type:uuid;not null" json:"tenant_id"`
	Name        string         `gorm:"type:text;not null" json:"name"`
	Code        string         `gorm:"type:text;not null" json:"code"`
	Address     datatypes.JSON `json:"address,omitempty"`
	ContactInfo datatypes.JSON `json:"contact_info,omitempty"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:now()" json:"updated_at"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}
