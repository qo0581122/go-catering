package shop

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
		form = request.ShopQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Shop{
		ShopName:     form.ShopName,
		Status:       form.Status,
		CityId:       form.CityId,
		ProvinceId:   form.ProvinceId,
		DistrictId:   form.DistrictId,
		BusinessFlag: form.BusinessFlag,
	}
	res := shopService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.ShopAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Shop{
		ShopName:      form.ShopName,
		CategoryId:    form.CategoryId,
		ShopAddress:   form.ShopAddress,
		Longitude:     form.Longitude,
		BusinessFlag:  form.BusinessFlag,
		ContactNumber: form.ContactNumber,
		DistrictId:    form.DistrictId,
		ProvinceId:    form.ProvinceId,
		CityId:        form.CityId,
		Status:        form.Status,
	}
	err = shopService.Add(model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.ShopUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Shop{
		Model: model.Model{
			ID: form.Id,
		},
		ShopName:      form.ShopName,
		CategoryId:    form.CategoryId,
		ShopAddress:   form.ShopAddress,
		BusinessFlag:  form.BusinessFlag,
		ContactNumber: form.ContactNumber,
		DistrictId:    form.DistrictId,
		ProvinceId:    form.ProvinceId,
		CityId:        form.CityId,
		Status:        form.Status,
	}
	err = shopService.Update(model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := shopService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
