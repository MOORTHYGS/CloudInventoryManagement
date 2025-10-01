package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/middleware"
	"github.com/moorthy/cloud_inventory/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

// ========================= CUSTOMER CONTROLLER ========================= //

// LoginCustomer - POST /customers/login
func LoginCustomer(c *gin.Context) {
	var input struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	if err := db.DB.Where("user_name = ?", input.UserName).First(&customer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	if !CheckPassword(customer.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT
	token, err := middleware.GenerateToken(customer.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"user":         customer,
	})
}

// CreateCustomer - POST /customers
func CreateCustomer(c *gin.Context) {
	var input struct {
		Name        string                 `json:"name" binding:"required"`
		UserName    string                 `json:"user_name" binding:"required"`
		Password    string                 `json:"password" binding:"required"`
		ContactInfo map[string]interface{} `json:"contact_info"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username already exists
	var existing models.Customer
	if err := db.DB.Where("user_name = ?", input.UserName).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Convert ContactInfo to datatypes.JSON
	contactJSON, _ := json.Marshal(input.ContactInfo)

	customer := models.Customer{
		ID:          uuid.New(),
		Name:        input.Name,
		UserName:    input.UserName,
		Password:    hashedPassword,
		ContactInfo: datatypes.JSON(contactJSON),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// GetCustomer - GET /customers?user_name=&password=
func GetCustomer(c *gin.Context) {
	userName := c.Query("user_name")
	password := c.Query("password")

	if userName == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_name and password required"})
		return
	}

	var customer models.Customer
	if err := db.DB.Where("user_name ILIKE ?", userName).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found with this username"})
		return
	}

	// Check password
	if !CheckPassword(customer.Password, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user creditionals"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer - PUT /customers/:id
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := db.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var input struct {
		Name        string                 `json:"name"`
		Password    string                 `json:"password"`
		OldPassword string                 `json:"old_password"`
		ContactInfo map[string]interface{} `json:"contact_info"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != "" {
		if input.OldPassword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Old password required"})
			return
		}
		if !CheckPassword(customer.Password, input.OldPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password does not match"})
			return
		}
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
			return
		}
		customer.Password = hashedPassword
	}

	if input.Name != "" {
		customer.Name = input.Name
	}

	if input.ContactInfo != nil {
		contactJSON, _ := json.Marshal(input.ContactInfo)
		customer.ContactInfo = datatypes.JSON(contactJSON)
	}

	customer.UpdatedAt = time.Now()

	if err := db.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// DeleteCustomer - DELETE /customers/:id
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Delete(&models.Customer{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// ========================= PASSWORD UTILS ========================= //

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
