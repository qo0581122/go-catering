package user_vip_level

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/user/request"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.UserVipLevelQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	levels := userVipLevelService.ListPage(form.PageNum, form.PageSize, &model.UserVipLevel{})
	response.OkWithData(levels, c)
}

func Add(c *gin.Context) {
	var (
		form = request.UserVipLevelAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	level := &model.UserVipLevel{
		LevelName:         form.LevelName,
		UpNeedIntegration: form.UpNeedIntegration,
		LevelDiscount:     form.LevelDiscount,
		NextLevelId:       form.NextLevelId,
		Level:             form.Level,
	}
	err = userVipLevelService.Add(level, form.PreLevelId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Adds(c *gin.Context) {
	var (
		form = request.UserVipLevelAddForms{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	var levels []model.UserVipLevel

	for _, item := range form.Forms {
		levels = append(levels, model.UserVipLevel{
			LevelName:         item.LevelName,
			UpNeedIntegration: item.UpNeedIntegration,
			LevelDiscount:     item.LevelDiscount,
			Level:             item.Level,
		})
	}
	err = userVipLevelService.Adds(levels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.UserVipLevelUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	level := &model.UserVipLevel{
		Id:                form.Id,
		LevelName:         form.LevelName,
		UpNeedIntegration: form.UpNeedIntegration,
		LevelDiscount:     form.LevelDiscount,
		NextLevelId:       form.NextLevelId,
	}
	err = userVipLevelService.Update(level)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := userVipLevelService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
