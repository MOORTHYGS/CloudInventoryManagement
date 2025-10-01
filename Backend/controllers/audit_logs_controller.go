package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// ✅ Create log manually (optional)
func CreateAuditLog(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var input models.AuditLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = uuid.New()
	input.TenantId = tenantUUID

	if err := db.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// ✅ Get all logs by tenant
func GetAuditLogsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var logs []models.AuditLog
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Order("timestamp desc").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

// ✅ Get single log by ID
func GetAuditLogByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid log ID"})
		return
	}

	var log models.AuditLog
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&log).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}
