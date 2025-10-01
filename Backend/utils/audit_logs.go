package utils

import (
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

func CreateAuditLog(tenantId uuid.UUID, userId uuid.UUID, action, table string, recordId uuid.UUID, details string) {
	log := models.AuditLog{
		ID:       uuid.New(),
		TenantId: tenantId,
		UserId:   userId,
		Action:   action,
		Table:    table, // renamed
		RecordId: recordId,
		Details:  details,
	}
	db.DB.Create(&log)
}
