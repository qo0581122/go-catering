package product

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/product/request"
	"catering/pkg/valid"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var (
		params = request.ProductQueryParams{}
	)
	msg, err := valid.BindAndValid(c, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Product{
		ProductName: params.ProductName,
		Status:      params.Status,
	}
	res := productService.ListPage(params.PageNum, params.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.ProductAddForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		fmt.Println(err.Error())
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Product{
		ProductName:   form.ProductName,
		Description:   form.Description,
		Specis:        form.Specis,
		Url:           form.Url,
		Status:        form.Status,
		CategoryId:    form.CategoryId,
		Discount:      form.Discount,
		OriginalPrice: form.OriginalPrice,
		PayPrice:      form.PayPrice,
	}
	err = productService.Add(model, form.AttributeIds, form.BatchIds, form.ProductIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.ProductUpdateForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Product{
		Model: model.Model{
			ID: form.Id,
		},
		ProductName: form.ProductName,
		Description: form.Description,
		Specis:      form.Specis,
		Url:         form.Url,
		Status:      form.Status,
		CategoryId:  form.CategoryId,
	}
	err = productService.Update(model, form.AttributeIds, form.BatchIds, form.ProductIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := productService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func ListBySpecis(c *gin.Context) {
	specis, _ := strconv.Atoi(c.Param("specis"))
	product := &model.Product{
		Specis: specis,
	}
	res := productService.List(product)
	if res == nil {
		response.Fail(c)
		return
	}
	response.OkWithData(response.NewApiResponse(res), c)
}
