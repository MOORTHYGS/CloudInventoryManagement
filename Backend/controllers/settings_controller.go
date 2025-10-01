package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// Get all settings for a tenant
func GetSettingsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var settings []models.Setting
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, settings)
}

// Get a single setting by key
func GetSettingByKey(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	key := c.Param("key")
	tenantUUID, _ := uuid.Parse(tenantID)

	var setting models.Setting
	if err := db.DB.Where("tenant_id = ? AND key = ?", tenantUUID, key).First(&setting).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// Create or update a setting
func UpsertSetting(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var input struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"` // store JSON as string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var setting models.Setting
	err := db.DB.Where("tenant_id = ? AND key = ?", tenantUUID, input.Key).First(&setting).Error
	if err != nil { // create new
		setting = models.Setting{
			ID:        uuid.New(),
			TenantId:  tenantUUID,
			Key:       input.Key,
			Value:     input.Value,
			UpdatedAt: time.Now(),
		}
		if err := db.DB.Create(&setting).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else { // update existing
		setting.Value = input.Value
		setting.UpdatedAt = time.Now()
		if err := db.DB.Save(&setting).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, setting)
}

// Delete a setting by key
func DeleteSetting(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	key := c.Param("key")
	tenantUUID, _ := uuid.Parse(tenantID)

	if err := db.DB.Where("tenant_id = ? AND key = ?", tenantUUID, key).Delete(&models.Setting{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Setting deleted successfully"})
}
