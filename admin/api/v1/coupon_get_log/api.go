package coupon_get_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/coupon/request"
	"catering/pkg/valid"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		params = request.CouponGetLogQueryParams{}
	)
	msg, err := valid.BindAndValid(c, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.CouponGetLog{
		Uid:       params.Uid,
		CouponId:  params.CouponId,
		UseStatus: params.UseStatus,
		GetType:   params.GetType,
	}
	res := couponGetLogService.ListPage(params.PageNum, params.PageSize, model)
	response.OkWithData(res, c)
}
