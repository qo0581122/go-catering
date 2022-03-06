package product

import (
	"catering/api/v1/product_category"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductCategoryRouter struct{}

func (d *ProductCategoryRouter) InitProductCategoryRouter(Router *gin.RouterGroup) {
	r := Router.Group("product-categorys").Use()
	{
		r.GET("", product_category.List)
		r.GET("all", product_category.ListAll)
	}
	r = Router.Group("product-category").Use(middleware.OperationRecord())
	{
		r.POST("add", product_category.Add)
		r.POST("update", product_category.Update)
		r.DELETE("delete/:id", product_category.Delete)
	}
}
