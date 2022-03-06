package user_integration

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum  int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	Uid      int    `uri:"uid"  form:"uid" json:"uid"`
	LevelId  uint64 `uri:"level_id" json:"level_id" form:"level_id"`
}

type QueryResponse struct {
	List  []*model.UserIntegration `json:"list,inline"`
	Total int                      `json:"total"`
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
	model := &model.UserIntegration{
		Uid:     form.Uid,
		LevelId: form.LevelId,
	}
	integrations := userIntegratonService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(integrations, c)
}

type ChangeForm struct {
	Id          uint64 `uri:"id"  form:"id" json:"id"`
	ChangeValue int    `uri:"change_value"  form:"change_value" json:"change_value"`
	ChangeType  int    `uri:"change_type"  form:"change_type" json:"change_type"`
	AddStatus   int    `uri:"add_status"  form:"add_status" json:"add_status"`
}

func Change(c *gin.Context) {
	var (
		form = ChangeForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	if form.AddStatus == 1 {
		err = userIntegratonService.IncreaseIntegration(form.Id, form.ChangeValue, form.ChangeType)
	} else if form.AddStatus == 2 {
		err = userIntegratonService.DecreaseIntegration(form.Id, form.ChangeValue, form.ChangeType)
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
