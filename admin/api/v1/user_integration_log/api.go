package user_integration_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/user/request"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.UserIntegrationLogQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.UserIntegrationLog{
		Uid:        form.Uid,
		ChangeType: form.ChangeType,
	}
	res := userIntegrationLogService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}
