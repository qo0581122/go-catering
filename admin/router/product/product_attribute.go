package product

import (
	"catering/api/v1/product_attribute"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductAttributeRouter struct{}

func (d *ProductRouter) InitProductAttributeRouter(Router *gin.RouterGroup) {
	r := Router.Group("product-attributes").Use()
	{
		r.GET("", product_attribute.List)
		r.GET("all", product_attribute.ListAll)
	}

	r = Router.Group("product-attribute").Use(middleware.OperationRecord())
	{
		r.POST("add", product_attribute.Add)
		r.POST("update", product_attribute.Update)
		r.DELETE("delete/:id", product_attribute.Delete)
	}
}
