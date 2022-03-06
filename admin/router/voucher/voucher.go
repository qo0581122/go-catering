package voucher

import (
	"catering/api/v1/voucher"
	"catering/api/v1/voucher_get_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type VoucherRouter struct{}

func (d *VoucherRouter) InitVoucherRouter(Router *gin.RouterGroup) {
	r := Router.Group("vouchers").Use()
	{
		r.GET("page", voucher.ListPage)

		r.GET("getLogs", voucher_get_log.ListPage)
	}
	r = Router.Group("voucher").Use(middleware.OperationRecord())
	{
		r.PUT("", voucher.Add)
		r.POST("", voucher.Update)
		r.DELETE(":id", voucher.Delete)
	}
}
