package district

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
		params = request.DistrictQueryParams{}
	)
	msg, err := app.BindAndValid(c, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	district := &model.District{
		CityId:       params.CityId,
		DistrictName: params.DistrictName,
		Status:       params.Status,
	}
	res := districtService.ListPage(params.PageNum, params.PageSize, district)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.DistrictAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	district := &model.District{
		CityId:       form.CityId,
		DistrictName: form.DistrictName,
		Status:       form.Status,
	}
	err = districtService.Add(district)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.DistrictUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	district := &model.District{
		Model:        model.Model{Id: form.Id},
		CityId:       form.CityId,
		DistrictName: form.DistrictName,
		Status:       form.Status,
	}
	err = districtService.Update(district)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := districtService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func List(c *gin.Context) {
	var (
		form = request.DistinctListParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}

	district := &model.District{
		CityId: form.CityId,
	}
	data := districtService.List(district)
	if data == nil {
		response.Fail(c)
		return
	}
	response.OkWithData(response.NewApiResponse(data), c)
}
