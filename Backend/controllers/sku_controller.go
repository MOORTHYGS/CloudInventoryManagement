package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// ========================= SKU CONTROLLER ========================= //

// GetAllSkus - GET /tenants/:tenant_id/users/:user_id/sku
func GetAllSkus(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var skus []models.Sku
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&skus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skus)
}

// GetSkuByID - GET /tenants/:tenant_id/sku/:id
func GetSkuByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	skuID := c.Param("id") // this is bigint / int, not UUID

	var sku models.Sku
	if err := db.DB.Where("tenant_id = ? AND sku_id = ?", tenantUUID, skuID).First(&sku).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}

	c.JSON(http.StatusOK, sku)
}

// AddSku - POST /tenants/:tenant_id/users/:user_id/sku
func AddSku(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var sku models.Sku
	if err := c.ShouldBindJSON(&sku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// enforce tenant_id
	sku.TenantID = tenantUUID

	if err := db.DB.Create(&sku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sku)
}

// UpdateSku - PUT /tenants/:tenant_id/users/:user_id/sku/:id
func UpdateSku(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	skuID := c.Param("id")

	var sku models.Sku
	if err := db.DB.Where("tenant_id = ? AND sku_id = ?", tenantUUID, skuID).First(&sku).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}

	var input models.Sku
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update allowed fields
	sku.SkuDesc = input.SkuDesc
	sku.LastEditBy = input.LastEditBy

	if err := db.DB.Save(&sku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sku)
}

// DeleteSku - DELETE /tenants/:tenant_id/users/:user_id/sku/:id
func DeleteSku(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	skuID := c.Param("id")

	if err := db.DB.Where("tenant_id = ? AND sku_id = ?", tenantUUID, skuID).Delete(&models.Sku{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SKU deleted successfully"})
}
