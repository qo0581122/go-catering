package product_attribute

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize      int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum       int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	AttributeName string `uri:"attribute_name"  form:"attribute_name" json:"attribute_name"`
	Status        int    `uri:"status" json:"status" form:"status"`
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
	productattribute := &model.ProductAttribute{
		AttributeName: form.AttributeName,
		Status:        form.Status,
	}
	result := productAttributeService.ListPage(form.PageNum, form.PageSize, productattribute)
	response.OkWithData(result, c)
}

type AddForm struct {
	Status          int      `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeName   string   `json:"attribute_name" form:"attribute_name" valid:"Required"`
	AttributeValues []string `json:"values" form:"values"`
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
	productattribute := &model.ProductAttribute{
		AttributeName: form.AttributeName,
		Status:        form.Status,
	}

	err = productAttributeService.Add(productattribute, form.AttributeValues)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

type UpdateForm struct {
	Id              uint64   `form:"id" json:"id" valid:"Required"`
	Status          int      `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeName   string   `json:"attribute_name" form:"attribute_name" valid:"Required"`
	AttributeValues []string `json:"values" form:"values"`
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
	productattribute := &model.ProductAttribute{
		Model:         model.Model{Id: form.Id},
		AttributeName: form.AttributeName,
		Status:        form.Status,
	}
	err = productAttributeService.Update(productattribute, form.AttributeValues)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := productAttributeService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func ListAll(c *gin.Context) {
	productattribute := &model.ProductAttribute{}
	result := productAttributeService.List(productattribute)
	response.OkWithData(response.NewApiResponse(result), c)
}
