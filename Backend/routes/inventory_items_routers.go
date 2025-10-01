package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func InventoryItemRouters(r *gin.Engine) {
	items := r.Group("/tenants/:tenant_id/inventory_items")
	{
		InventoryItem := items.Group("/")
		InventoryItem.Use(middleware.UserAuthMiddleware())
		{
			InventoryItem.POST("/", controllers.CreateInventoryItem)
			InventoryItem.GET("/", controllers.GetInventoryItemsByTenant)
			InventoryItem.GET("/:id", controllers.GetInventoryItemByID)
			InventoryItem.PUT("/:id", controllers.UpdateInventoryItemByID)
			InventoryItem.DELETE("/:id", controllers.DeleteInventoryItemByID)
		}
	}
}
