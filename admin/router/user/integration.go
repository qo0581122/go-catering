package user

import (
	"catering/api/v1/user_integration"
	"catering/api/v1/user_integration_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type IntegrationRouter struct{}

func (d *IntegrationRouter) InitIntegrationRouter(Router *gin.RouterGroup) {
	router := Router.Group("user")
	r := router.Group("integrations").Use()
	{
		r.GET("page", user_integration.ListPage)
		r.GET("logs", user_integration_log.ListPage)
	}
	r = router.Group("integration").Use(middleware.OperationRecord())
	{
		r.POST("change", user_integration.Change)
	}
}
