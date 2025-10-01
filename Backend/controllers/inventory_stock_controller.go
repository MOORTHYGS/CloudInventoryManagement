package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// ========================= INVENTORY STOCK CONTROLLER ========================= //

// CreateInventoryStock - POST /tenants/:tenant_id/inventory_stock
func CreateInventoryStock(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var input models.InventoryStock
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stock := models.InventoryStock{
		ID:           uuid.New(),
		TenantId:     tenantUUID,
		InventoryId:  input.InventoryId,
		WarehouseId:  input.WarehouseId,
		Quantity:     input.Quantity,
		ReorderLevel: input.ReorderLevel,
		UpdatedAt:    time.Now(),
	}

	if err := db.DB.Create(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

// GetInventoryStockByTenant - GET /tenants/:tenant_id/inventory_stock
func GetInventoryStockByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var stocks []models.InventoryStock
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

// GetInventoryStockByID - GET /tenants/:tenant_id/inventory_stock/:id
func GetInventoryStockByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	stockID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(stockID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	var stock models.InventoryStock
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

// UpdateInventoryStockByID - PUT /tenants/:tenant_id/inventory_stock/:id
func UpdateInventoryStockByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	stockID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(stockID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	var stock models.InventoryStock
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory stock not found"})
		return
	}

	var input struct {
		Quantity     *int `json:"quantity"`
		ReorderLevel *int `json:"reorder_level"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Quantity != nil {
		stock.Quantity = *input.Quantity
	}
	if input.ReorderLevel != nil {
		stock.ReorderLevel = *input.ReorderLevel
	}

	stock.UpdatedAt = time.Now()

	if err := db.DB.Save(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stock)
}

// DeleteInventoryStockByID - DELETE /tenants/:tenant_id/inventory_stock/:id
func DeleteInventoryStockByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	stockID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(stockID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.InventoryStock{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory stock deleted successfully"})
}
