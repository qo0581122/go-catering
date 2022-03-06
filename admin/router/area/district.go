package area

import (
	"catering/api/v1/district"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type DistrictRouter struct{}

func (d *DistrictRouter) InitDistrictRouter(Router *gin.RouterGroup) {
	r := Router.Group("districts").Use(middleware.OperationRecord())
	{
		r.GET("", district.List)
		r.GET("city/:id", district.ListDistrictByCityId)
	}
	r = Router.Group("district").Use(middleware.OperationRecord())
	{
		r.POST("create", district.Add)
		r.POST("update", district.Update)
		r.DELETE("delete/:id", district.Delete)
	}
}
