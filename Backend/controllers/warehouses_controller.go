package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
	"gorm.io/datatypes"
)

// ========================= WAREHOUSE CONTROLLER ========================= //

// CreateWarehouse - POST /tenants/:tenant_id/warehouses
func CreateWarehouse(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var input models.Warehouse
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	warehouse := models.Warehouse{
		ID:          uuid.New(),
		TenantId:    tenantUUID,
		Name:        input.Name,
		Code:        input.Code,
		Address:     input.Address,
		ContactInfo: input.ContactInfo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.DB.Create(&warehouse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, warehouse)
}

// GetWarehousesByTenant - GET /tenants/:tenant_id/warehouses
func GetWarehousesByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var warehouses []models.Warehouse
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&warehouses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, warehouses)
}

// GetWarehouseByID - GET /tenants/:tenant_id/warehouses/:id
func GetWarehouseByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	warehouseID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(warehouseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warehouse ID"})
		return
	}

	var warehouse models.Warehouse
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&warehouse).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}

	c.JSON(http.StatusOK, warehouse)
}

// UpdateWarehouseByID - PUT /tenants/:tenant_id/warehouses/:id
func UpdateWarehouseByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	warehouseID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(warehouseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warehouse ID"})
		return
	}

	var warehouse models.Warehouse
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&warehouse).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}

	var input struct {
		Name        string         `json:"name"`
		Code        string         `json:"code"`
		Address     datatypes.JSON `json:"address"`
		ContactInfo datatypes.JSON `json:"contact_info"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		warehouse.Name = input.Name
	}
	if input.Code != "" {
		warehouse.Code = input.Code
	}
	if input.Address != nil {
		warehouse.Address = input.Address
	}
	if input.ContactInfo != nil {
		warehouse.ContactInfo = input.ContactInfo
	}

	warehouse.UpdatedAt = time.Now()

	if err := db.DB.Save(&warehouse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, warehouse)
}

// DeleteWarehouseByID - DELETE /tenants/:tenant_id/warehouses/:id
func DeleteWarehouseByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	warehouseID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(warehouseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warehouse ID"})
		return
	}

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.Warehouse{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted successfully"})
}
