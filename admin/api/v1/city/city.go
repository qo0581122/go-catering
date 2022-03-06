package city

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize   int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum    int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	CityName   string `uri:"city_name"  form:"city_name" json:"city_name"`
	Status     int    `uri:"status" json:"status" form:"status"`
	ProvinceId uint64 `uri:"province_id"  form:"province_id" json:"province_id"`
}

type ResponseCity struct {
	*model.City
	Province_name string `json:"province_name" `
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
	city := &model.City{
		ProvinceId: params.ProvinceId,
		CityName:   params.CityName,
		Status:     params.Status,
	}
	res := cityService.ListPage(params.PageNum, params.PageSize, city)
	response.OkWithData(res, c)
}

type AddCityForm struct {
	ProvinceId uint64 `form:"province_id" json:"province_id" valid:"Required"`
	Status     int    `form:"status" json:"status" valid:"Range(1,2)"`
	CityName   string `json:"city_name" form:"city_name" valid:"Required"`
}

func Add(c *gin.Context) {
	var (
		form = AddCityForm{}
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

type UpdateCityForm struct {
	Id         uint64 `form:"id" json:"id" valid:"Required"`
	ProvinceId uint64 `form:"province_id" json:"province_id" valid:"Required"`
	Status     int    `form:"status" json:"status" valid:"Range(1,2)"`
	CityName   string `json:"city_name" form:"city_name" valid:"Required"`
}

func Update(c *gin.Context) {
	var (
		form = UpdateCityForm{}
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
	// id, _ := strconv.Atoi(c.DefaultQuery("id", ""))
	err := cityService.Delete(uint64(id))
	if err != nil {
		response.Fail(c)
		return
	}
	response.Ok(c)
}

func ListCityByProvinceId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// id, _ := strconv.Atoi(c.DefaultQuery("province_id", ""))
	city := &model.City{
		ProvinceId: uint64(id),
	}
	res := cityService.List(city)
	response.OkWithData(response.NewApiResponse(res), c)
}
