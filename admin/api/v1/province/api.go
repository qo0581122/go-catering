package province

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
	ProvinceName string `uri:"province_name"  form:"province_name" json:"province_name"`
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
	model := &model.Province{
		ProvinceName: form.ProvinceName,
		Status:       form.Status,
	}
	res := provinceService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}

type AddProvinceForm struct {
	ProvinceName string `form:"province_name" json:"province_name" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
}

func Add(c *gin.Context) {
	var (
		form = AddProvinceForm{}
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

type UpdateProvinceForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	ProvinceName string `form:"province_name" json:"province_name" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
}

func Update(c *gin.Context) {
	var (
		form = UpdateProvinceForm{}
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

func ListAll(c *gin.Context) {
	res := provinceService.List(&model.Province{})
	response.OkWithData(response.NewApiResponse(res), c)
}
