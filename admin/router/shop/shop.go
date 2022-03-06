package shop

import (
	"catering/api/v1/shop"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ShopRouter struct{}

func (d *ShopRouter) InitShopRouter(Router *gin.RouterGroup) {
	r := Router.Group("shops").Use()
	{
		r.GET("page", shop.ListPage)
	}
	r = Router.Group("shop").Use(middleware.OperationRecord())
	{
		r.PUT("", shop.Add)
		r.POST("", shop.Update)
		r.DELETE(":id", shop.Delete)
	}
}
