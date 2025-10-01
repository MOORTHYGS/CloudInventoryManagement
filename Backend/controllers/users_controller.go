package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/middleware"
	"github.com/moorthy/cloud_inventory/models"
)

// ========================= USER CONTROLLER ========================= //

// UserLogin - POST /tenants/:tenant_id/users/login
func UserLogin(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	var input struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.Where("user_name = ? AND tenant_id = ?", input.UserName, tenantID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or tenant"})
		return
	}

	// Use CheckPassword utility
	if !CheckPassword(user.PasswordHash, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT
	token, err := middleware.GenerateUserToken(user.ID.String(), tenantID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"user":         user,
	})
}

// CreateUser - POST /tenants/:tenant_id/users
func CreateUser(c *gin.Context) {
	tenantID := c.Param("tenant_id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password before storing
	hashedPassword, err := HashPassword(input.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		ID:           uuid.New(),
		TenantId:     tenantUUID,
		Name:         input.Name,
		Email:        input.Email,
		UserName:     input.UserName,
		PasswordHash: hashedPassword,
		Role:         input.Role,
		Status:       input.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// defaults
	if user.Role == "" {
		user.Role = "Staff"
	}
	if user.Status == "" {
		user.Status = "active"
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsersByTenant - GET /tenants/:tenant_id/users
func GetUsersByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var users []models.User
	if err := db.DB.Where("tenant_id = ?", tenantUUID).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByTenant - GET /tenants/:tenant_id/users/:username
func GetUserByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	username := c.Param("username")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var user models.User
	if err := db.DB.Where("tenant_id = ? AND user_name = ?", tenantUUID, username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found for this tenant"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsersByRole - GET /tenants/:tenant_id/users/roles/:role
func GetUsersByRole(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	role := c.Param("role")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var users []models.User
	if err := db.DB.Where("tenant_id = ? AND role = ?", tenantUUID, role).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found with this role"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUserByTenant - PUT /tenants/:tenant_id/users/:username
func UpdateUserByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	username := c.Param("username")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var user models.User
	if err := db.DB.Where("tenant_id = ? AND user_name = ?", tenantUUID, username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
		Role         string `json:"role"`
		Status       string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.PasswordHash != "" {
		// Hash new password before saving
		hashedPassword, err := HashPassword(input.PasswordHash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.PasswordHash = hashedPassword
	}
	if input.Role != "" {
		user.Role = input.Role
	}
	if input.Status != "" {
		user.Status = input.Status
	}

	user.UpdatedAt = time.Now()

	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserByTenant - DELETE /tenants/:tenant_id/users/:username
func DeleteUserByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	username := c.Param("username")

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	if err := db.DB.Where("tenant_id = ? AND user_name = ?", tenantUUID, username).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
