package product

import (
	"catering/api/v1/product_category"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductCategoryRouter struct{}

func (d *ProductCategoryRouter) InitProductCategoryRouter(Router *gin.RouterGroup) {
	r := Router.Group("product").Group("categorys").Use()
	{
		r.GET("page", product_category.ListPage)
		r.GET("", product_category.List)
	}
	r = Router.Group("product").Group("category").Use(middleware.OperationRecord())
	{
		r.PUT("", product_category.Add)
		r.POST("", product_category.Update)
		r.DELETE(":id", product_category.Delete)
	}
}
