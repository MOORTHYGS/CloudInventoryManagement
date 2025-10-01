package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/models"
)

// CreateSupplier
func CreateSupplier(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var input models.Supplier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := models.Supplier{
		ID:          uuid.New(),
		TenantId:    tenantUUID,
		Name:        input.Name,
		ContactInfo: input.ContactInfo,
		CreatedAt:   time.Now(),
	}

	if err := db.DB.Create(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

// GetSuppliersByTenant
func GetSuppliersByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	tenantUUID, _ := uuid.Parse(tenantID)

	var suppliers []models.Supplier
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&suppliers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

// GetSupplierByID
func GetSupplierByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var supplier models.Supplier
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// UpdateSupplierByID
func UpdateSupplierByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	var supplier models.Supplier
	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	var input struct {
		Name        string                 `json:"name"`
		ContactInfo map[string]interface{} `json:"contact_info"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		supplier.Name = input.Name
	}
	if input.ContactInfo != nil {
		supplier.ContactInfo = input.ContactInfo
	}

	if err := db.DB.Save(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// DeleteSupplierByID
func DeleteSupplierByID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	tenantUUID, _ := uuid.Parse(tenantID)
	idUUID, _ := uuid.Parse(id)

	if err := db.DB.Where("tenant_id = ? AND id = ?", tenantUUID, idUUID).Delete(&models.Supplier{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}
