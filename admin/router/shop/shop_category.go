package shop

import (
	"catering/api/v1/shop_category"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ShopCategoryRouter struct{}

func (d *ShopCategoryRouter) InitShopCategoryRouter(Router *gin.RouterGroup) {
	r := Router.Group("shop-categorys").Use(middleware.OperationRecord())
	{
		r.GET("", shop_category.List)
		r.GET("all", shop_category.ListAll)
	}
	r = Router.Group("shop-category").Use(middleware.OperationRecord())
	{
		r.POST("add", shop_category.Add)
		r.POST("update", shop_category.Update)
		r.DELETE("delete", shop_category.Delete)
	}
}
