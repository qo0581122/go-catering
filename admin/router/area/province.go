package area

import (
	"catering/api/v1/province"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProvinceRouter struct{}

func (d *ProvinceRouter) InitProvinceRouter(Router *gin.RouterGroup) {
	r := Router.Group("provinces").Use(middleware.OperationRecord())
	{
		r.GET("", province.List)
		r.GET("all", province.ListAll)
	}
	r = Router.Group("province").Use(middleware.OperationRecord())
	{
		r.POST("create", province.Add)
		r.POST("update", province.Update)
		r.DELETE("delete/:id", province.Delete)
	}
}
