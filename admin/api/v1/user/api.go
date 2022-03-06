package user

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize  int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum   int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	Mobie     string `uri:"mobie" json:"mobie" form:"mobie"`
	ChannelId uint64 `uri:"channel_id" json:"channel_id" form:"channel_id"`
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
	model := &model.User{
		Mobie:     form.Mobie,
		ChannelId: form.ChannelId,
	}
	res := userService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}
