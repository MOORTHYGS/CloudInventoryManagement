package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId  uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	PlanId    uuid.UUID `gorm:"type:uuid;not null" json:"plan_id"`
	Status    string    `gorm:"type:text;default:'trial'" json:"status"` // active/expired/trial
	StartDate time.Time `gorm:"default:now()" json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Renewal   time.Time `json:"renewal_date"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
