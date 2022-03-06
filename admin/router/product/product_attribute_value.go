package product

import (
	"catering/api/v1/product_attribute_value"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductAttributeValueRouter struct{}

func (d *ProductRouter) InitProductAttributeValueRouter(Router *gin.RouterGroup) {
	r := Router.Group("product-attribute-values").Use()
	{
		r.GET("", product_attribute_value.List)
		r.GET("all", product_attribute_value.ListAll)
	}
	r = Router.Group("product-attribute-value").Use(middleware.OperationRecord())
	{
		r.POST("add", product_attribute_value.Add)
		r.POST("update", product_attribute_value.Update)
		r.DELETE("delete/:id", product_attribute_value.Delete)
	}
}
