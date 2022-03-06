package coupon_get_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize  int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum   int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	Uid       int    `uri:"uid" json:"uid" form:"uid"`
	CouponId  uint64 `uri:"coupon_id" json:"coupon_id" form:"coupon_id"`
	UseStatus uint32 `uri:"use_status" json:"use_status" form:"use_status"`
	GetType   uint32 `uri:"get_type" json:"get_type" form:"get_type"`
}

func List(c *gin.Context) {
	var (
		params = QueryParams{}
	)
	msg, err := app.BindAndValid(c, &params)
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
