package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// CreatePayment
func CreatePayment(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var input models.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := models.Payment{
		ID:             uuid.New(),
		TenantId:       tenantUUID,
		SubscriptionId: input.SubscriptionId,
		Amount:         input.Amount,
		Method:         input.Method,
		Status:         input.Status,
		TransactionID:  input.TransactionID,
		PaidAt:         time.Now(),
	}

	if err := db.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

// GetPaymentsByTenant
func GetPaymentsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var payments []models.Payment
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// GetPaymentByID
func GetPaymentByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var payment models.Payment
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// UpdatePayment
func UpdatePayment(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var payment models.Payment
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	var input struct {
		Amount        float64   `json:"amount"`
		Method        string    `json:"method"`
		Status        string    `json:"status"`
		TransactionID string    `json:"transaction_id"`
		PaidAt        time.Time `json:"paid_at"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Amount != 0 {
		payment.Amount = input.Amount
	}
	if input.Method != "" {
		payment.Method = input.Method
	}
	if input.Status != "" {
		payment.Status = input.Status
	}
	if input.TransactionID != "" {
		payment.TransactionID = input.TransactionID
	}
	if !input.PaidAt.IsZero() {
		payment.PaidAt = input.PaidAt
	}

	if err := db.DB.Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// DeletePayment
func DeletePayment(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.Payment{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
