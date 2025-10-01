package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId  uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	UserId    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Action    string    `gorm:"type:text;not null" json:"action"` // e.g. CREATE, UPDATE, DELETE
	Table     string    `gorm:"type:text" json:"table"`           // e.g. users, suppliers
	RecordId  uuid.UUID `gorm:"type:uuid" json:"record_id"`
	Timestamp time.Time `gorm:"default:now()" json:"timestamp"`
	Details   string    `gorm:"type:jsonb" json:"details"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
