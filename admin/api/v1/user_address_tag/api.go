package user_address_tag

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
		form = request.UserAddressTagQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.UserAddressTag{
		TagName: form.AddressTagName,
		Status:  form.Status,
	}
	tags := userAddressTagService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(tags, c)
}

func Add(c *gin.Context) {
	var (
		form = request.UserAddressTagAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	tag := &model.UserAddressTag{
		TagName:         form.TagName,
		TextColor:       form.TextColor,
		BorderColor:     form.BorderColor,
		BackgroundColor: form.BackgroundColor,
		Sort:            form.Sort,
		Status:          form.Status,
	}
	err = userAddressTagService.Add(tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.UserAddressTagUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	tag := &model.UserAddressTag{
		TagName:         form.TagName,
		TextColor:       form.TextColor,
		BorderColor:     form.BorderColor,
		BackgroundColor: form.BackgroundColor,
		Sort:            form.Sort,
		Status:          form.Status,
	}
	err = userAddressTagService.Update(tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := userAddressTagService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
