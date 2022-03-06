package product

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize    int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum     int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	ProductName string `uri:"product_name"  form:"product_name" json:"product_name"`
	Status      int    `uri:"status" json:"status" form:"status"`
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
	model := &model.Product{
		ProductName: params.ProductName,
		Status:      params.Status,
	}
	res := productService.ListPage(params.PageNum, params.PageSize, model)
	response.OkWithData(res, c)
}

type AddForm struct {
	ProductName   string   `json:"product_name" form:"product_name" valid:"Required"`
	Description   string   `json:"description" form:"description" valid:"Required"`
	Specis        int      `json:"specis" form:"specis" valid:"Range(1,2)"`
	Url           string   `json:"url" form:"url" valid:"Required"`
	Status        int      `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryId    uint64   `json:"category_id" form:"category_id"`
	AttributeIds  []uint64 `json:"attribute_ids" form:"attribute_ids"`
	BatchIds      []uint64 `json:"batch_ids" form:"batch_ids"`
	ProductIds    []uint64 `json:"product_ids" form:"product_ids" `
	Discount      int      `json:"discount" form:"discount"`
	OriginalPrice int      `json:"original_price" form:"original_price"`
	PayPrice      int      `json:"pay_price" form:"pay_price"`
}

func Add(c *gin.Context) {
	var (
		form = AddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
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

type UpdateForm struct {
	Id           uint64   `form:"id" json:"id" valid:"Required"`
	ProductName  string   `json:"product_name" form:"product_name" valid:"Required"`
	Description  string   `json:"description" form:"description" valid:"Required"`
	Specis       int      `json:"specis" form:"specis" valid:"Range(1,2)"`
	Url          string   `json:"url" form:"url" valid:"Required"`
	Status       int      `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryId   uint64   `json:"category_id" form:"category_id"`
	AttributeIds []uint64 `json:"attribute_ids" form:"attribute_ids"`
	BatchIds     []uint64 `json:"batch_ids" form:"batch_ids"`
	ProductIds   []uint64 `json:"product_ids" form:"product_ids" `
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
	model := &model.Product{
		Model: model.Model{
			Id: form.Id,
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
