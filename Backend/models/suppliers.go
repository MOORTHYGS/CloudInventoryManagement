package models

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID          uuid.UUID              `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId    uuid.UUID              `gorm:"type:uuid;not null" json:"tenant_id"`
	Name        string                 `gorm:"type:text;not null" json:"name"`
	ContactInfo map[string]interface{} `gorm:"type:jsonb" json:"contact_info"`
	CreatedAt   time.Time              `gorm:"default:now()" json:"created_at"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
