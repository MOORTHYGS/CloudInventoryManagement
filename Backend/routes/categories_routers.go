package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func CategoryRouters(r *gin.Engine) {
	categories := r.Group("/tenants/:tenant_id/categories")
	{
		Category := categories.Group("/")
		Category.Use(middleware.UserAuthMiddleware())
		{
			Category.POST("/", controllers.CreateCategory)
			Category.GET("/", controllers.GetCategoriesByTenant)
			Category.GET("/:id", controllers.GetCategoryByID)
			Category.PUT("/:id", controllers.UpdateCategoryByID)
			Category.DELETE("/:id", controllers.DeleteCategoryByID)
		}
	}
}
