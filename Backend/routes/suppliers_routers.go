package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func SupplierRouters(r *gin.Engine) {
	suppliers := r.Group("/tenants/:tenant_id/suppliers")
	{
		supplier := suppliers.Group("/")
		supplier.Use(middleware.UserAuthMiddleware())
		{
			supplier.POST("/", controllers.CreateSupplier)
			supplier.GET("/", controllers.GetSuppliersByTenant)
			supplier.GET("/:id", controllers.GetSupplierByID)
			supplier.PUT("/:id", controllers.UpdateSupplierByID)
			supplier.DELETE("/:id", controllers.DeleteSupplierByID)
		}
	}
}
