package user_integration

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/user/request"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.UserIntegrationQueryParams{}
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

func Change(c *gin.Context) {
	var (
		form = request.UserIntegrationChangeForm{}
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
