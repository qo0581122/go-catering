package province

import (
	"catering/model"
	"catering/model/area/request"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.ProvinceQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Province{
		ProvinceName: form.ProvinceName,
		Status:       form.Status,
	}
	res := provinceService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.ProvinceAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	province := &model.Province{
		ProvinceName: form.ProvinceName,
		Status:       form.Status,
	}
	err = provinceService.Add(province)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.ProvinceUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	province := &model.Province{
		Model: model.Model{
			Id: form.Id,
		},
		ProvinceName: form.ProvinceName,
		Status:       form.Status,
	}
	err = provinceService.Update(province)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := provinceService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func List(c *gin.Context) {
	res := provinceService.List(&model.Province{})
	response.OkWithData(response.NewApiResponse(res), c)
}
