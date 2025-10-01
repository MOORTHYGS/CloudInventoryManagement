package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moorthy/cloud_inventory/controllers"
)

func SubscriptionRouters(r *gin.Engine) {
	subscriptions := r.Group("/tenants/:tenant_id/subscriptions")
	{
		subscription := subscriptions.Group("/")
		{
			subscription.POST("/", controllers.CreateSubscription)
			subscription.GET("/", controllers.GetSubscriptionsByTenant)
			subscription.GET("/:id", controllers.GetSubscriptionByID)
			subscription.PUT("/:id", controllers.UpdateSubscription)
			subscription.DELETE("/:id", controllers.DeleteSubscription)
		}
	}
}
