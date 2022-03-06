package product_attribute_value

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
	AttributeValue string `uri:"attribute_value"  form:"attribute_value" json:"attribute_value"`
	AttributeId    uint64 `uri:"attribute_id" form:"attribute_id" json:"attribute_id"`
	Status         int    `uri:"status" json:"status" form:"status"`
}

func List(c *gin.Context) {
	var (
		params = QueryParams{}
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

type AddForm struct {
	Status         int    `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeValue string `json:"attribute_value" form:"attribute_value" valid:"Required"`
	AttributeId    uint64 `form:"attribute_id" json:"attribute_id" valid:"Required"`
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

type UpdateForm struct {
	Id             uint64 `form:"id" json:"id" valid:"Required"`
	Status         int    `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeValue string `json:"attribute_value" form:"attribute_value" valid:"Required"`
	AttributeId    uint64 `form:"attribute_id" json:"attribute_id" valid:"Required"`
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
	productattributevalue := &model.ProductAttributeValue{
		Model:          model.Model{Id: form.Id},
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

func ListAll(c *gin.Context) {
	model := &model.ProductAttributeValue{}
	data := productAttributeValueService.List(model)
	response.OkWithData(response.NewApiResponse(data), c)
}
