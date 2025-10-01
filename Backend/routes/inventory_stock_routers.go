package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func InventoryStockRouters(r *gin.Engine) {
	stocks := r.Group("/tenants/:tenant_id/inventory_stock")
	{
		stock := stocks.Group("/")
		stock.Use(middleware.UserAuthMiddleware())
		{
			stock.POST("/", controllers.CreateInventoryStock)
			stock.GET("/", controllers.GetInventoryStockByTenant)
			stock.GET("/:id", controllers.GetInventoryStockByID)
			stock.PUT("/:id", controllers.UpdateInventoryStockByID)
			stock.DELETE("/:id", controllers.DeleteInventoryStockByID)
		}
	}
}
