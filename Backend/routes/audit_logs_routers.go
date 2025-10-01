package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func AuditLogRoutes(r *gin.Engine) {
	logs := r.Group("/tenants/:tenant_id/audit_logs")
	{
		log := logs.Group("/")
		log.Use(middleware.UserAuthMiddleware())
		{
			log.POST("/", controllers.CreateAuditLog) // manual create
			log.GET("/", controllers.GetAuditLogsByTenant)
			log.GET("/:id", controllers.GetAuditLogByID)
		}
	}
}
