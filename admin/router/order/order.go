package order

import (
	"catering/api/v1/order"

	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (d *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	r := Router.Group("orders").Use()
	{
		r.GET("", order.List)
		r.GET("ids", order.GetOrderProduct)
	}
}
