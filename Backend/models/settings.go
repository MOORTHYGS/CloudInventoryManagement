package models

import (
	"time"

	"github.com/google/uuid"
)

type Setting struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId  uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	Key       string    `gorm:"type:text;not null" json:"key"` // e.g. "currency", "tax_rate"
	Value     string    `gorm:"type:jsonb" json:"value"`       // store JSON for flexibility
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`
}

func (Setting) TableName() string {
	return "settings"
}
