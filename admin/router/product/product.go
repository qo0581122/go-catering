package product

import (
	"catering/api/v1/product"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (d *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	r := Router.Group("products").Use()
	{
		r.GET("", product.List)
		r.GET("specis/:specis", product.ListBySpecis)
	}

	r = Router.Group("product").Use(middleware.OperationRecord())
	{
		r.POST("add", product.Add)
		r.POST("update", product.Update)
		r.DELETE("delete/:id", product.Delete)
	}
}
