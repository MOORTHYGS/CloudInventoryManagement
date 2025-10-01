package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// CreateSubscription
func CreateSubscription(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var input models.Subscription
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscription := models.Subscription{
		ID:        uuid.New(),
		TenantId:  tenantUUID,
		PlanId:    input.PlanId,
		Status:    input.Status,
		StartDate: time.Now(),
		EndDate:   input.EndDate,
		Renewal:   input.Renewal,
	}

	if err := db.DB.Create(&subscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subscription)
}

// GetSubscriptionsByTenant
func GetSubscriptionsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var subscriptions []models.Subscription
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptions)
}

// GetSubscriptionByID
func GetSubscriptionByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var subscription models.Subscription
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&subscription).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// UpdateSubscription
func UpdateSubscription(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var subscription models.Subscription
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&subscription).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	var input struct {
		PlanId  uuid.UUID `json:"plan_id"`
		Status  string    `json:"status"`
		EndDate time.Time `json:"end_date"`
		Renewal time.Time `json:"renewal_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.PlanId != uuid.Nil {
		subscription.PlanId = input.PlanId
	}
	if input.Status != "" {
		subscription.Status = input.Status
	}
	if !input.EndDate.IsZero() {
		subscription.EndDate = input.EndDate
	}
	if !input.Renewal.IsZero() {
		subscription.Renewal = input.Renewal
	}

	if err := db.DB.Save(&subscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// DeleteSubscription
func DeleteSubscription(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.Subscription{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription deleted successfully"})
}
