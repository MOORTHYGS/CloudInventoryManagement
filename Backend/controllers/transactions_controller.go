package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// CreateTransaction
func CreateTransaction(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = uuid.New()
	input.TenantId = tenantUUID
	input.TotalPrice = float64(input.Quantity) * input.UnitPrice
	if input.Date.IsZero() {
		input.Date = time.Now()
	}

	if err := db.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GetTransactionsByTenant
func GetTransactionsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var transactions []models.Transaction
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransactionByID
func GetTransactionByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var transaction models.Transaction
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction
func UpdateTransaction(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var transaction models.Transaction
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	var input struct {
		Type       string     `json:"type"`
		Quantity   int        `json:"quantity"`
		UnitPrice  float64    `json:"unit_price"`
		Note       string     `json:"note"`
		SupplierId *uuid.UUID `json:"supplier_id"`
		CustomerId *uuid.UUID `json:"customer_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Type != "" {
		transaction.Type = input.Type
	}
	if input.Quantity != 0 {
		transaction.Quantity = input.Quantity
		transaction.TotalPrice = float64(transaction.Quantity) * transaction.UnitPrice
	}
	if input.UnitPrice != 0 {
		transaction.UnitPrice = input.UnitPrice
		transaction.TotalPrice = float64(transaction.Quantity) * transaction.UnitPrice
	}
	if input.Note != "" {
		transaction.Note = input.Note
	}
	if input.SupplierId != nil {
		transaction.SupplierId = input.SupplierId
	}
	if input.CustomerId != nil {
		transaction.CustomerId = input.CustomerId
	}

	if err := db.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction
func DeleteTransaction(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.Transaction{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
