package product_attribute

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/product/request"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.AttributeQueryParams{}
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

func Add(c *gin.Context) {
	var (
		form = request.AttributeAddForm{}
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

func Update(c *gin.Context) {
	var (
		form = request.AttributeUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	productattribute := &model.ProductAttribute{
		Model:         model.Model{ID: form.Id},
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

func List(c *gin.Context) {
	productattribute := &model.ProductAttribute{}
	result := productAttributeService.List(productattribute)
	response.OkWithData(response.NewApiResponse(result), c)
}
