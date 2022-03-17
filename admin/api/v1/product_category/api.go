package product_category

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/product/request"
	"catering/pkg/valid"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.CategoryQueryParams{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	productcategory := &model.ProductCategory{
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	result := productCategoryService.ListPage(form.PageNum, form.PageSize, productcategory)
	response.OkWithData(result, c)
}

func Add(c *gin.Context) {
	var (
		form = request.CategoryAddForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	productcategory := &model.ProductCategory{
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	err = productCategoryService.Add(productcategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.CategoryUpdateForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	productcategory := &model.ProductCategory{
		Model:        model.Model{ID: form.Id},
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	err = productCategoryService.Update(productcategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := productCategoryService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func List(c *gin.Context) {
	productcategory := &model.ProductCategory{}
	data := productCategoryService.List(productcategory)
	response.OkWithData(response.NewApiResponse(data), c)
}
