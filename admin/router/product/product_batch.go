package product

import (
	"catering/api/v1/product_batch"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ProductBatchRouter struct{}

func (d *ProductRouter) InitProductBatchRouter(Router *gin.RouterGroup) {
	r := Router.Group("product-batchs").Use()
	{
		r.GET("", product_batch.List)
		r.GET("all", product_batch.ListAll)
	}
	r = Router.Group("product-batch").Use(middleware.OperationRecord())
	{
		r.POST("add", product_batch.Add)
		r.POST("update", product_batch.Update)
		r.DELETE("delete/:id", product_batch.Delete)
	}
}
