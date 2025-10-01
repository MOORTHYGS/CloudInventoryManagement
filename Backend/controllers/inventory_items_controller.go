package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// ========================= INVENTORY ITEM CONTROLLER ========================= //

// CreateInventoryItem - POST /tenants/:tenant_id/inventory_items
func CreateInventoryItem(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var input models.InventoryItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.InventoryItem{
		ID:          uuid.New(),
		TenantId:    tenantUUID,
		CategoryId:  input.CategoryId,
		SkuID:       input.SkuID,
		Name:        input.Name,
		Description: input.Description,
		Unit:        input.Unit,
		Price:       input.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if item.Unit == "" {
		item.Unit = "pcs"
	}
	if item.Price == 0 {
		item.Price = 0
	}

	if err := db.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetInventoryItemsByTenant - GET /tenants/:tenant_id/inventory_items
func GetInventoryItemsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var items []models.InventoryItem
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetInventoryItemByID - GET /tenants/:tenant_id/inventory_items/:id
func GetInventoryItemByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	itemID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(itemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory item ID"})
		return
	}

	var item models.InventoryItem
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateInventoryItemByID - PUT /tenants/:tenant_id/inventory_items/:id
func UpdateInventoryItemByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	itemID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(itemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory item ID"})
		return
	}

	var item models.InventoryItem
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	var input struct {
		CategoryId  *uuid.UUID `json:"category_id"`
		SkuID       *uint16    `json:"sku_id"`
		Name        string     `json:"name"`
		Description string     `json:"description"`
		Unit        string     `json:"unit"`
		Price       *float64   `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.CategoryId != nil {
		item.CategoryId = input.CategoryId
	}
	if input.SkuID != nil {
		item.SkuID = *input.SkuID
	}
	if input.Name != "" {
		item.Name = input.Name
	}
	if input.Description != "" {
		item.Description = input.Description
	}
	if input.Unit != "" {
		item.Unit = input.Unit
	}
	if input.Price != nil {
		item.Price = *input.Price
	}

	item.UpdatedAt = time.Now()

	if err := db.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteInventoryItemByID - DELETE /tenants/:tenant_id/inventory_items/:id
func DeleteInventoryItemByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	itemID := c.Param("id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	idUUID, err := uuid.Parse(itemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory item ID"})
		return
	}

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.InventoryItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item deleted successfully"})
}
