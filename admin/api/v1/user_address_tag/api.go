package user_address_tag

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize       int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum        int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	AddressTagName string `uri:"tag_name"  form:"tag_name" json:"tag_name"`
	Status         uint32 `uri:"status" json:"status" form:"status"`
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
	model := &model.AddressTag{
		TagName: form.AddressTagName,
		Status:  form.Status,
	}
	tags := userAddressTagService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(tags, c)
}

type AddForm struct {
	TagName         string `json:"tag_name" form:"tag_name" valid:"Required"`
	TextColor       string `json:"text_color" form:"text_color" valid:"Required"`
	BorderColor     string `json:"border_color" form:"border_color" `
	BackgroundColor string `json:"background_color" form:"background_color" `
	Sort            uint32 `json:"sort" form:"sort" valid:"Required"`
	Status          uint32 `json:"status" form:"status" valid:"Required"`
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
	tag := &model.AddressTag{
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

type UpdateForm struct {
	Id              uint64 `form:"id" json:"id" valid:"Required"`
	TagName         string `json:"tag_name" form:"tag_name" valid:"Required"`
	TextColor       string `json:"text_color" form:"text_color" valid:"Required"`
	BorderColor     string `json:"border_color" form:"border_color" `
	BackgroundColor string `json:"background_color" form:"background_color" `
	Sort            uint32 `json:"sort" form:"sort" valid:"Required"`
	Status          uint32 `json:"status" form:"status" valid:"Required"`
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
	tag := &model.AddressTag{
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
