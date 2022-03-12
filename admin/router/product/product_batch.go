package product

import (
	"catering/api/v1/product_batch"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductBatchRouter struct{}

func (d *ProductRouter) InitProductBatchRouter(Router *gin.RouterGroup) {
	r := Router.Group("product").Group("batchs").Use()
	{
		r.GET("page", product_batch.ListPage)
		r.GET("", product_batch.List)
	}
	r = Router.Group("product").Group("batch").Use(middleware.OperationRecord())
	{
		r.PUT("", product_batch.Add)
		r.POST("", product_batch.Update)
		r.DELETE(":id", product_batch.Delete)
	}
}
