package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// ========================= TENANT CONTROLLER ========================= //

// CreateTenant - POST /customers/:customer_id/tenants
func CreateTenant(c *gin.Context) {
	customerID := c.Param("customer_id")

	var input struct {
		Name   string `json:"name" binding:"required"`
		Domain string `json:"domain"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate customer ID
	custUUID, err := uuid.Parse(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	// Enforce max 3 tenants per customer
	var count int64
	db.DB.Model(&models.Tenant{}).Where("customer_id = ?", custUUID).Count(&count)
	if count >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tenant limit reached (max 3 per customer)"})
		return
	}

	tenant := models.Tenant{
		ID:         uuid.New(),
		CustomerID: custUUID,
		Name:       input.Name,
		Domain:     input.Domain,
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := db.DB.Create(&tenant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tenant)
}

// GetTenantsByCustomer - GET /customers/:customer_id/tenants
func GetTenantsByCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	custUUID, err := uuid.Parse(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var tenants []models.Tenant
	if err := db.DB.Where("customer_id = ?", custUUID).Find(&tenants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenants)
}

// GetTenantByCustomerAndTenant - GET /customers/:customer_id/tenants/:tenant_id
func GetTenantByCustomerAndTenant(c *gin.Context) {
	customerID := c.Param("customer_id")
	tenantID := c.Param("tenant_id")

	custUUID, err := uuid.Parse(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var tenant models.Tenant
	if err := db.DB.Where("id = ? AND customer_id = ?", tenantUUID, custUUID).First(&tenant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found for this customer"})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

// UpdateTenantByCustomer - PUT /customers/:customer_id/tenants/:tenant_id
func UpdateTenantByCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	tenantID := c.Param("tenant_id")

	custUUID, err := uuid.Parse(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var tenant models.Tenant
	if err := db.DB.Where("id = ? AND customer_id = ?", tenantUUID, custUUID).First(&tenant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}

	var input struct {
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		tenant.Name = input.Name
	}
	if input.Domain != "" {
		tenant.Domain = input.Domain
	}
	if input.Status != "" {
		tenant.Status = input.Status
	}

	tenant.UpdatedAt = time.Now()

	if err := db.DB.Save(&tenant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

// DeleteTenantByCustomer - DELETE /customers/:customer_id/tenants/:tenant_id
func DeleteTenantByCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	tenantID := c.Param("tenant_id")

	custUUID, err := uuid.Parse(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	if err := db.DB.Where("id = ? AND customer_id = ?", tenantUUID, custUUID).Delete(&models.Tenant{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tenant deleted successfully"})
}
