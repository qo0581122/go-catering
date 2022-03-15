package product_attribute_value

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
		params = request.AttributeValueQueryParams{}
	)
	msg, err := app.BindAndValid(c, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.ProductAttributeValue{
		AttributeValue: params.AttributeValue,
		Status:         params.Status,
		AttributeId:    params.AttributeId,
	}
	res := productAttributeValueService.ListPage(params.PageNum, params.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.AttributeValueAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	productattributevalue := &model.ProductAttributeValue{
		AttributeValue: form.AttributeValue,
		Status:         form.Status,
	}
	err = productAttributeValueService.Add(productattributevalue)
	if err == nil {
		response.Ok(c)
		return
	}
	response.Fail(c)
}

func Update(c *gin.Context) {
	var (
		form = request.AttributeValueUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	productattributevalue := &model.ProductAttributeValue{
		Model:          model.Model{ID: form.Id},
		AttributeValue: form.AttributeValue,
		Status:         form.Status,
	}
	err = productAttributeValueService.Update(productattributevalue)
	if err != nil {
		response.Fail(c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := productAttributeValueService.Delete(uint64(id))
	if err != nil {
		response.Fail(c)
		return
	}
	response.Ok(c)
}

func List(c *gin.Context) {
	model := &model.ProductAttributeValue{}
	data := productAttributeValueService.List(model)
	response.OkWithData(response.NewApiResponse(data), c)
}
