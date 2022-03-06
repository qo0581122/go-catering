package shop

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
	ShopName     string `uri:"shop_name"  form:"shop_name" json:"shop_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
	CityId       uint64 `uri:"city_id"  form:"city_id" json:"city_id"`
	ProvinceId   uint64 `uri:"province_id"  form:"province_id" json:"province_id"`
	DistrictId   uint64 `uri:"district_id"  form:"district_id" json:"district_id"`
	CategoryId   uint64 `uri:"category_id"  form:"category_id" json:"category_id"`
	BusinessFlag int    `uri:"business_flag"  form:"business_flag" json:"business_flag"`
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

type AddForm struct {
	CategoryId    uint64 `json:"category_id" form:"category_id"  valid:"Required"`
	ShopName      string `json:"shop_name" form:"shop_name" valid:"Required"`
	ShopAddress   string `json:"shop_address" form:"shop_address"  valid:"Required"`
	Longitude     string `json:"logintude" form:"logintude"  valid:"Required"`
	Latitude      string `json:"latitude" form:"latitude" valid:"Required"`
	BusinessFlag  int    `json:"business_flag" form:"business_flag"  valid:"Range(1,2)"`
	ContactNumber string `json:"contact_number" form:"contact_number"  valid:"Required"`
	DistrictId    uint64 `json:"district_id" form:"district_id"  valid:"Required"`
	ProvinceId    uint64 `json:"province_id" form:"province_id"  valid:"Required"`
	CityId        uint64 `json:"city_id" form:"city_id"  valid:"Required"`
	Status        int    `form:"status" json:"status" valid:"Range(1,2)"`
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

type UpdateForm struct {
	Id            uint64 `form:"id" json:"id" valid:"Required"`
	CategoryId    uint64 `json:"category_id" form:"category_id"  valid:"Required"`
	ShopName      string `json:"shop_name" form:"shop_name" valid:"Required"`
	ShopAddress   string `json:"shop_address" form:"shop_address"  valid:"Required"`
	BusinessFlag  int    `json:"business_flag" form:"business_flag"  valid:"Range(1,2)"`
	ContactNumber string `json:"contact_number" form:"contact_number"  valid:"Required"`
	DistrictId    uint64 `json:"district_id" form:"district_id"  valid:"Required"`
	ProvinceId    uint64 `json:"province_id" form:"province_id"  valid:"Required"`
	CityId        uint64 `json:"city_id" form:"city_id"  valid:"Required"`
	Status        int    `form:"status" json:"status" valid:"Range(1,2)"`
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
	model := &model.Shop{
		Model: model.Model{
			Id: form.Id,
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

// func ListDistrictByCityId(c *gin.Context) {
// 	var (
// 		appG = app.Gin{C: c}
// 	)
// 	id, _ := strconv.Atoi(c.DefaultQuery("city_id", ""))
// 	district := &model.District{
// 		CityId: id,
// 	}
// 	res := district.ListByCityId()
// 	if res != nil {
// 		appG.Response(http.StatusOK, e.ERROR, "查找失败")
// 		return
// 	}
// 	appG.Response(http.StatusOK, e.SUCCESS, res)
// }
