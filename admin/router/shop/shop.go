package shop

import (
	"catering/api/v1/shop"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ShopRouter struct{}

func (d *ShopRouter) InitShopRouter(Router *gin.RouterGroup) {
	r := Router.Group("shops").Use(middleware.OperationRecord())
	{
		r.GET("", shop.List)
	}
	r = Router.Group("shop").Use(middleware.OperationRecord())
	{
		r.POST("add", shop.Add)
		r.POST("update", shop.Update)
		r.DELETE("delete/:id", shop.Delete)
	}
}
