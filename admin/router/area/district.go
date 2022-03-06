package area

import (
	"catering/api/v1/district"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type DistrictRouter struct{}

func (d *DistrictRouter) InitDistrictRouter(Router *gin.RouterGroup) {
	r := Router.Group("districts").Use()
	{
		r.GET("page", district.ListPage)
		r.GET("", district.List)
	}
	r = Router.Group("district").Use(middleware.OperationRecord())
	{
		r.PUT("", district.Add)
		r.POST("", district.Update)
		r.DELETE(":id", district.Delete)
	}
}
