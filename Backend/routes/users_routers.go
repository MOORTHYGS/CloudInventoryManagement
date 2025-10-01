package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func UserRouters(r *gin.Engine) {
	users := r.Group("/tenants/:tenant_id/users")
	{
		// Public login route for users
		users.POST("/login", controllers.UserLogin)

		// Customer-protected routes (customers manage users)
		protected := users.Group("/")
		protected.Use(middleware.AuthMiddleware()) // <- only customers
		{
			protected.POST("/", controllers.CreateUser)
			protected.GET("/", controllers.GetUsersByTenant)
			protected.GET("/:username", controllers.GetUserByTenant)
			protected.GET("/roles/:role", controllers.GetUsersByRole)
			protected.PUT("/:username", controllers.UpdateUserByTenant)
			protected.DELETE("/:username", controllers.DeleteUserByTenant)
		}
	}
}
