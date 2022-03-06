package district

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
	DistrictName string `uri:"district_name"  form:"district_name" json:"district_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
	CityId       uint64 `uri:"city_id"  form:"city_id" json:"city_id"`
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
	district := &model.District{
		CityId:       params.CityId,
		DistrictName: params.DistrictName,
		Status:       params.Status,
	}
	res := districtService.ListPage(params.PageNum, params.PageSize, district)
	response.OkWithData(res, c)
}

type AddForm struct {
	CityId       uint64 `form:"city_id" json:"city_id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	DistrictName string `json:"district_name" form:"district_name" valid:"Required"`
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

type UpdateForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	CityId       uint64 `form:"city_id" json:"city_id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	DistrictName string `json:"district_name" form:"district_name" valid:"Required"`
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

func ListDistrictByCityId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	district := &model.District{
		CityId: uint64(id),
	}
	data := districtService.List(district)
	if data == nil {
		response.Fail(c)
		return
	}
	response.OkWithData(response.NewApiResponse(data), c)
}
