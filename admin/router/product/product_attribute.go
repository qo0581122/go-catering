package product

import (
	"catering/api/v1/product_attribute"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductAttributeRouter struct{}

func (d *ProductRouter) InitProductAttributeRouter(Router *gin.RouterGroup) {
	Router = Router.Group("product")
	r := Router.Group("attributes").Use()
	{
		r.GET("page", product_attribute.ListPage)
		r.GET("", product_attribute.List)
	}

	r = Router.Group("attribute").Use(middleware.OperationRecord())
	{
		r.PUT("", product_attribute.Add)
		r.POST("", product_attribute.Update)
		r.DELETE(":id", product_attribute.Delete)
	}
}
