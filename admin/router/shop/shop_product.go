package shop

import (
	"catering/api/v1/shop_product"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ShopProductRouter struct{}

func (d *ShopProductRouter) InitShopProductRouter(Router *gin.RouterGroup) {
	r := Router.Group("shop").Group("products").Use()
	{
		r.GET("page", shop_product.ListPage)
	}
	r = Router.Group("shop").Use(middleware.OperationRecord())
	{
		r.PUT("", shop_product.Add)
		r.POST("", shop_product.Update)
		r.DELETE(":id", shop_product.Delete)
	}
}
