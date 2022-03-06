package shop_category

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize     int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum      int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	CategoryName string `uri:"category_name"  form:"category_name" json:"category_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
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
	category := &model.ShopCategory{
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	result := shopCategoryService.ListPage(form.PageNum, form.PageSize, category)
	response.OkWithData(result, c)
}

type AddForm struct {
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryName string `json:"category_name" form:"category_name" valid:"Required"`
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

type UpdateForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryName string `json:"category_name" form:"category_name" valid:"Required"`
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
	category := &model.ShopCategory{
		Model:        model.Model{Id: form.Id},
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

func ListAll(c *gin.Context) {
	data := shopCategoryService.List(&model.ShopCategory{})
	response.OkWithData(response.NewApiResponse(data), c)
}
