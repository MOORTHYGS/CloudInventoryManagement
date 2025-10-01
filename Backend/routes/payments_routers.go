package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func PaymentRouters(r *gin.Engine) {
	payments := r.Group("/tenants/:tenant_id/payments")
	{
		payment := payments.Group("/")
		payment.Use(middleware.UserAuthMiddleware())
		{
			payment.POST("/", controllers.CreatePayment)
			payment.GET("/", controllers.GetPaymentsByTenant)
			payment.GET("/:id", controllers.GetPaymentByID)
			payment.PUT("/:id", controllers.UpdatePayment)
			payment.DELETE("/:id", controllers.DeletePayment)
		}
	}
}
