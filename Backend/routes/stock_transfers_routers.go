package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
	"github.com/moorthy/cloud_inventory/middleware"
)

func StockTransferRouters(r *gin.Engine) {
	transfers := r.Group("/tenants/:tenant_id/stock_transfers")
	{
		stock_transfers := transfers.Group("/")
		stock_transfers.Use(middleware.UserAuthMiddleware())
		{
			stock_transfers.POST("/", controllers.CreateStockTransfer)
			stock_transfers.GET("/", controllers.GetStockTransfersByTenant)
			stock_transfers.GET("/:id", controllers.GetStockTransferByID)
			stock_transfers.PUT("/:id", controllers.UpdateStockTransferStatus)
			stock_transfers.DELETE("/:id", controllers.DeleteStockTransfer)
		}
	}
}
