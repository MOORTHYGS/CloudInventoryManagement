package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/db"
	"github.com/moorthy/cloud_inventory/routes"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.TenantRoutes(r)
	routes.Customer(r)
	routes.UserRouters(r)
	routes.SkuRouters(r)
	routes.CategoryRouters(r)
	routes.InventoryItemRouters(r)
	routes.WarehouseRouters(r)
	routes.InventoryStockRouters(r)
	routes.SupplierRouters(r)
	routes.TransactionRouters(r)
	routes.StockTransferRouters(r)
	routes.SubscriptionRouters(r)
	routes.SettingsRoutes(r)
	routes.AuditLogRoutes(r)

	r.Run(":8080")
}
