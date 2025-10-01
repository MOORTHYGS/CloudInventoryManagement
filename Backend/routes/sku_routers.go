package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func SkuRouters(r *gin.Engine) {
	sku := r.Group("/tenants/:tenant_id/sku")
	sku.Use(middleware.UserAuthMiddleware())
	{
		// CRUD routes for SKU
		sku.GET("/", controllers.GetAllSkus)      // Get all SKUs for a tenant
		sku.GET("/:id", controllers.GetSkuByID)   // Get a single SKU by ID
		sku.POST("/", controllers.AddSku)         // Create new SKU
		sku.PUT("/:id", controllers.UpdateSku)    // Update SKU by ID
		sku.DELETE("/:id", controllers.DeleteSku) // Delete SKU by ID
	}
}
