package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func TransactionRouters(r *gin.Engine) {
	transactions := r.Group("/tenants/:tenant_id/transactions")
	{
		transaction := transactions.Group("/")
		transaction.Use(middleware.UserAuthMiddleware())
		{
			transaction.POST("/", controllers.CreateTransaction)
			transaction.GET("/", controllers.GetTransactionsByTenant)
			transaction.GET("/:id", controllers.GetTransactionByID)
			transaction.PUT("/:id", controllers.UpdateTransaction)
			transaction.DELETE("/:id", controllers.DeleteTransaction)
		}
	}
}
