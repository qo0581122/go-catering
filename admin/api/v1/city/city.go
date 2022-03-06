package city

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
		params = request.CityQueryParams{}
	)
	msg, err := app.BindAndValid(c, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	city := &model.City{
		ProvinceId: params.ProvinceId,
		CityName:   params.CityName,
		Status:     params.Status,
	}
	res := cityService.ListPage(params.PageNum, params.PageSize, city)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.CityAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	city := &model.City{
		ProvinceId: form.ProvinceId,
		CityName:   form.CityName,
		Status:     form.Status,
	}
	err = cityService.Add(city)
	if err == nil {
		response.Ok(c)
		return
	}
	response.Fail(c)
}

func Update(c *gin.Context) {
	var (
		form = request.CityUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	city := &model.City{
		Model:      model.Model{Id: form.Id},
		ProvinceId: form.ProvinceId,
		CityName:   form.CityName,
		Status:     form.Status,
	}
	err = cityService.Update(city)
	if err != nil {
		response.Fail(c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := cityService.Delete(uint64(id))
	if err != nil {
		response.Fail(c)
		return
	}
	response.Ok(c)
}

func List(c *gin.Context) {
	var (
		form = request.CityListParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	city := &model.City{
		ProvinceId: form.ProvinceId,
	}
	res := cityService.List(city)
	response.OkWithData(response.NewApiResponse(res), c)
}
