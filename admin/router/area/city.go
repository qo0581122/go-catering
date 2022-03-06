package area

import (
	"catering/api/v1/city"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type CityRouter struct{}

func (d *CityRouter) InitCityRouter(Router *gin.RouterGroup) {
	r := Router.Group("citys").Use()
	{
		r.GET("page", city.ListPage)
		r.GET("", city.List)
	}
	r = Router.Group("city").Use(middleware.OperationRecord())
	{
		r.PUT("", city.Add)
		r.POST("", city.Update)
		r.DELETE(":id", city.Delete)
	}
}
