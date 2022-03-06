package user_integration_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize   int `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum    int `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	Uid        int `uri:"uid" json:"uid" form:"uid"`
	ChangeType int `uri:"change_type" json:"change_type" form:"change_type"`
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
	model := &model.UserIntegrationLog{
		Uid:        form.Uid,
		ChangeType: form.ChangeType,
	}
	res := userIntegrationLogService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}
