package voucher_get_log

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
	VoucherId uint64 `uri:"voucher_id" json:"voucher_id" form:"voucher_id"`
	UseStatus uint32 `uri:"use_status" json:"use_status" form:"use_status"`
	GetType   uint32 `uri:"get_type" json:"get_type" form:"get_type"`
}

func List(c *gin.Context) {
	var (
		form = QueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.VoucherGetLog{
		Uid:       form.Uid,
		VoucherId: form.VoucherId,
		UseStatus: form.UseStatus,
		GetType:   form.GetType,
	}
	res := voucherGetLogService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}
