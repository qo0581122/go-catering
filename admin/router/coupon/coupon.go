package coupon

import (
	"catering/api/v1/coupon"
	"catering/api/v1/coupon_get_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type CouponRouter struct{}

func (d *CouponRouter) InitCouponRouter(Router *gin.RouterGroup) {
	r := Router.Group("coupons").Use(middleware.OperationRecord())
	{
		r.GET("", coupon.List)
		r.GET("get-logs", coupon_get_log.List)
	}
	r = Router.Group("coupon").Use(middleware.OperationRecord())
	{
		r.POST("add", coupon.Add)
		r.POST("update", coupon.Update)
		r.DELETE("delete/:id", coupon.Delete)
	}
}
