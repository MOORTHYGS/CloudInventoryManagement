package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func WarehouseRouters(r *gin.Engine) {
	warehouses := r.Group("/tenants/:tenant_id/warehouses")
	{
		warehouse := warehouses.Group("/")
		warehouse.Use(middleware.UserAuthMiddleware())
		{
			warehouse.POST("/", controllers.CreateWarehouse)
			warehouse.GET("/", controllers.GetWarehousesByTenant)
			warehouse.GET("/:id", controllers.GetWarehouseByID)
			warehouse.PUT("/:id", controllers.UpdateWarehouseByID)
			warehouse.DELETE("/:id", controllers.DeleteWarehouseByID)
		}
	}
}
