package user_vip_level

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize int `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum  int `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
}

type QueryResponse struct {
	List  []*model.UserVipLevel `json:"list,inline"`
	Total int                   `json:"total"`
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
	levels := userVipLevelService.ListPage(form.PageNum, form.PageSize, &model.UserVipLevel{})
	response.OkWithData(levels, c)
}

type AddForm struct {
	PreLevelId        uint64 `json:"pre_level_id" form:"pre_level_id"`
	LevelName         string `json:"level_name" form:"level_name" valid:"Required"`
	UpNeedIntegration int    `json:"up_need_integration" form:"up_need_integration" valid:"Required"`
	LevelDiscount     int    `json:"level_discount" form:"level_discount" `
	NextLevelId       uint64 `json:"next_level_id" form:"next_level_id"`
	Level             int    `json:"level" form:"level"`
}

func Add(c *gin.Context) {
	var (
		form = AddForm{}
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

type AddForms struct {
	Forms []struct {
		LevelName         string `json:"level_name" form:"level_name" valid:"Required"`
		UpNeedIntegration int    `json:"up_need_integration" form:"up_need_integration" valid:"Required"`
		LevelDiscount     int    `json:"level_discount" form:"level_discount" `
		Level             int    `json:"level" form:"level"`
	} `json:"forms"`
}

func Adds(c *gin.Context) {
	var (
		form = AddForms{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	var levels []*model.UserVipLevel

	for _, item := range form.Forms {
		levels = append(levels, &model.UserVipLevel{
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

type UpdateForm struct {
	Id                uint64 `form:"id" json:"id" valid:"Required"`
	LevelName         string `json:"level_name" form:"level_name" valid:"Required"`
	UpNeedIntegration int    `json:"up_need_integration" form:"up_need_integration" valid:"Required"`
	LevelDiscount     int    `json:"level_discount" form:"level_discount" `
	NextLevelId       uint64 `json:"next_level_id" form:"next_level_id"`
}

func Update(c *gin.Context) {
	var (
		form = UpdateForm{}
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
