package shop_category

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/shop/request"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.ShopCategoryQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	category := &model.ShopCategory{
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	result := shopCategoryService.ListPage(form.PageNum, form.PageSize, category)
	response.OkWithData(result, c)
}

func Add(c *gin.Context) {
	var (
		form = request.ShopCategoryAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	category := &model.ShopCategory{
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	err = shopCategoryService.Add(category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.ShopCategoryUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	category := &model.ShopCategory{
		Model:        model.Model{ID: form.Id},
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	err = shopCategoryService.Update(category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := shopCategoryService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func List(c *gin.Context) {
	var (
		form = request.ShopCategoryListParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	data := shopCategoryService.List(&model.ShopCategory{})
	response.OkWithData(response.NewApiResponse(data), c)
}
