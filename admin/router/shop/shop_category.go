package shop

import (
	"catering/api/v1/shop_category"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type ShopCategoryRouter struct{}

func (d *ShopCategoryRouter) InitShopCategoryRouter(Router *gin.RouterGroup) {
	router := Router.Group("shop")
	r := router.Group("categorys").Use()
	{
		r.GET("page", shop_category.ListPage)
		r.GET("", shop_category.List)
	}
	r = router.Group("category").Use(middleware.OperationRecord())
	{
		r.PUT("", shop_category.Add)
		r.POST("", shop_category.Update)
		r.DELETE(":id", shop_category.Delete)
	}
}
