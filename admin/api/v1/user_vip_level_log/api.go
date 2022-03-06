package user_vip_level_log

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

type UserVipLevelLogQueryParams struct {
	PageSize      int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum       int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	Uid           int    `uri:"uid" json:"uid" form:"uid"`
	BeforeLevelId uint64 `uri:"before_level_id" json:"before_level_id" form:"before_level_id"`
	AfterLevelId  uint64 `uri:"after_level_id" json:"after_level_id" form:"after_level_id"`
}

func ListUserVipLevelLog(c *gin.Context) {
	var (
		form = UserVipLevelLogQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
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
