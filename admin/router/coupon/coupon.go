package coupon

import (
	"catering/api/v1/coupon"
	"catering/api/v1/coupon_get_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type CouponRouter struct{}

func (d *CouponRouter) InitCouponRouter(Router *gin.RouterGroup) {
	r := Router.Group("coupons").Use()
	{
		r.GET("page", coupon.ListPage)
		r.GET("getLogs", coupon_get_log.ListPage)
	}
	r = Router.Group("coupon").Use(middleware.OperationRecord())
	{
		r.PUT("", coupon.Add)
		r.POST("", coupon.Update)
		r.DELETE(":id", coupon.Delete)
	}
}
