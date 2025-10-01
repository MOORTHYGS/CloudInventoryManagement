package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func SettingsRoutes(r *gin.Engine) {
	settings := r.Group("/tenants/:tenant_id/settings")
	{
		protected := settings.Group("/")
		protected.Use(middleware.UserAuthMiddleware())
		{
			protected.GET("/", controllers.GetSettingsByTenant)
			protected.GET("/:key", controllers.GetSettingByKey)
			protected.POST("/", controllers.UpsertSetting)
			protected.PUT("/", controllers.UpsertSetting)
			protected.DELETE("/:key", controllers.DeleteSetting)
		}
	}
}
