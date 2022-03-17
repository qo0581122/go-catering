package user_vip_level_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/user/request"
	"catering/pkg/valid"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.UserVipLevelLogQueryParams{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.UserVipLevelLog{
		Uid:           form.Uid,
		BeforeLevelId: form.BeforeLevelId,
		AfterLevelId:  form.AfterLevelId,
	}
	res := userVipLevelLogService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}
