package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// Customer represents the customers table in Supabase
type Customer struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string         `gorm:"type:text;not null" json:"name"`
	UserName    string         `gorm:"type:text;not null;unique" json:"user_name"`
	Password    string         `gorm:"type:text;not null" json:"password"`
	ContactInfo datatypes.JSON `gorm:"type:jsonb" json:"contact_info"` // <-- updated
	CreatedAt   time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}
