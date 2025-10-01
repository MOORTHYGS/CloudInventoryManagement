package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func Customer(r *gin.Engine) {
	customer := r.Group("/customers")
	{
		// Public route to create a customer
		customer.POST("/", controllers.CreateCustomer)

		// Login route to generate JWT
		customer.POST("/login", controllers.LoginCustomer)

		// Protected routes (require token)
		protected := customer.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Get customer info
			protected.GET("/", controllers.GetCustomer)

			// Update customer info
			protected.PUT("/update/:id", controllers.UpdateCustomer)

			// Delete customer
			protected.DELETE("/delete/:id", controllers.DeleteCustomer)
		}
	}
}
