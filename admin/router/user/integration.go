package user

import (
	"catering/api/v1/user_integration"
	"catering/api/v1/user_integration_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type IntegrationRouter struct{}

func (d *IntegrationRouter) InitIntegrationRouter(Router *gin.RouterGroup) {
	r := Router.Group("user-integrations").Use(middleware.OperationRecord())
	{
		r.GET("", user_integration.List)
		r.GET("logs", user_integration_log.List)
	}
	r = Router.Group("user-integration").Use(middleware.OperationRecord())
	{
		r.POST("change", user_integration.Change)
	}
}
