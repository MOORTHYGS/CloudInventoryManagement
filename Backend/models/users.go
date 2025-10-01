package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TenantId     uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	Name         string    `gorm:"type:text;not null" json:"name"`
	Email        string    `gorm:"type:text;not null" json:"email"`
	UserName     string    `gorm:"type:text;not null" json:"user_name"`
	PasswordHash string    `gorm:"type:text;not null" json:"password_hash"`
	Role         string    `gorm:"type:text;default:'Staff'" json:"role"`
	Status       string    `gorm:"type:text;default:'active'" json:"status"`
	CreatedAt    time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:now()" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
