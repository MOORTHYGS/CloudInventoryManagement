package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func TenantRoutes(r *gin.Engine) {
	tenants := r.Group("/customers/:customer_id/tenants")
	tenants.Use(middleware.AuthMiddleware()) // Protect all tenant routes
	{
		tenants.POST("/", controllers.CreateTenant)
		tenants.GET("/", controllers.GetTenantsByCustomer)
		tenants.GET("/:tenant_id", controllers.GetTenantByCustomerAndTenant)
		tenants.PUT("/:tenant_id", controllers.UpdateTenantByCustomer)
		tenants.DELETE("/:tenant_id", controllers.DeleteTenantByCustomer)
	}
}
