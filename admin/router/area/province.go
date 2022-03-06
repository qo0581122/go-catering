package area

import (
	"catering/api/v1/province"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProvinceRouter struct{}

func (d *ProvinceRouter) InitProvinceRouter(Router *gin.RouterGroup) {
	r := Router.Group("provinces").Use()
	{
		r.GET("page", province.ListPage)
		r.GET("", province.List)
	}
	r = Router.Group("province").Use(middleware.OperationRecord())
	{
		r.PUT("", province.Add)
		r.POST("", province.Update)
		r.DELETE(":id", province.Delete)
	}
}
