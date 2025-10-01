package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// CreateStockTransfer
func CreateStockTransfer(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var input models.StockTransfer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfer := models.StockTransfer{
		ID:            uuid.New(),
		TenantId:      tenantUUID,
		InventoryId:   input.InventoryId,
		FromWarehouse: input.FromWarehouse,
		ToWarehouse:   input.ToWarehouse,
		Quantity:      input.Quantity,
		TransferDate:  time.Now(),
		CreatedBy:     input.CreatedBy,
		Status:        "initiated",
	}

	if err := db.DB.Create(&transfer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transfer)
}

// GetStockTransfersByTenant
func GetStockTransfersByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var transfers []models.StockTransfer
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&transfers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transfers)
}

// GetStockTransferByID
func GetStockTransferByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var transfer models.StockTransfer
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&transfer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock transfer not found"})
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// UpdateStockTransferStatus
func UpdateStockTransferStatus(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var transfer models.StockTransfer
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&transfer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock transfer not found"})
		return
	}

	var input struct {
		Status string `json:"status"` // initiated/completed/cancelled
	}

	if err := c.ShouldBindJSON(&input); err != nil || (input.Status != "initiated" && input.Status != "completed" && input.Status != "cancelled") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	transfer.Status = input.Status
	if err := db.DB.Save(&transfer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// DeleteStockTransfer
func DeleteStockTransfer(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.StockTransfer{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock transfer deleted successfully"})
}
