package voucher

import (
	"catering/api/v1/voucher"
	"catering/api/v1/voucher_get_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type VoucherRouter struct{}

func (d *VoucherRouter) InitVoucherRouter(Router *gin.RouterGroup) {
	r := Router.Group("vouchers").Use(middleware.OperationRecord())
	{
		r.GET("", voucher.List)

		r.GET("get-logs", voucher_get_log.List)
	}
	r = Router.Group("voucher").Use(middleware.OperationRecord())
	{
		r.POST("add", voucher.Add)
		r.POST("update", voucher.Update)
		r.DELETE("delete/:id", voucher.Delete)
	}
}
