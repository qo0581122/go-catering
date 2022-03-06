package area

import (
	"catering/api/v1/city"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type CityRouter struct{}

func (d *CityRouter) InitCityRouter(Router *gin.RouterGroup) {
	r := Router.Group("citys").Use(middleware.OperationRecord())
	{
		r.GET("", city.List)
		r.GET("province/:id", city.ListCityByProvinceId)
	}
	r = Router.Group("city").Use(middleware.OperationRecord())
	{
		r.POST("create", city.Add)
		r.POST("update", city.Update)
		r.DELETE("delete/:id", city.Delete)
	}
}
