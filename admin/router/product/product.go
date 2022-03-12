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
		r.GET("page", product.List)
		r.GET("specis/:specis", product.ListBySpecis)
	}

	r = Router.Group("product").Use(middleware.OperationRecord())
	{
		r.PUT("", product.Add)
		r.POST("", product.Update)
		r.DELETE(":id", product.Delete)
	}
}
