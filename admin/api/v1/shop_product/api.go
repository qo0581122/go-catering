package shop_product

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/shop/request"
	"catering/pkg/valid"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.ShopProductQueryParams{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.ShopProduct{
		ProductId: form.ProductId,
		ShopId:    form.ShopId,
		Status:    form.Status,
	}
	res := shopProductService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.ShopProductAddForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.ShopProduct{
		ProductId: form.ProductId,
		ShopId:    form.ShopId,
		Status:    form.Status,
	}
	err = shopProductService.Add(model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.ShopProductUpdateForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.ShopProduct{
		ProductId: form.ProductId,
		ShopId:    form.ShopId,
		Status:    form.Status,
	}
	err = shopProductService.Update(model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := shopProductService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
