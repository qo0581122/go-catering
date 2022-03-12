package product

import (
	"catering/api/v1/product_attribute_value"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductAttributeValueRouter struct{}

func (d *ProductRouter) InitProductAttributeValueRouter(Router *gin.RouterGroup) {
	Router = Router.Group("product").Group("attribute")
	r := Router.Group("values").Use()
	{
		r.GET("page", product_attribute_value.ListPage)
		r.GET("", product_attribute_value.List)
	}
	r = Router.Group("value").Use(middleware.OperationRecord())
	{
		r.PUT("", product_attribute_value.Add)
		r.POST("", product_attribute_value.Update)
		r.DELETE(":id", product_attribute_value.Delete)
	}
}
