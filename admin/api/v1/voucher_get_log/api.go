package voucher_get_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/voucher/request"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.VoucherGetLogQueryParams{}
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
