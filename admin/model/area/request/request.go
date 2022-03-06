package request

import "catering/model/common/request"

type DistinctListParams struct {
	CityId uint64 `form:"city_id" json:"city_id"`
}

type CityListParams struct {
	ProvinceId uint64 `form:"province_id" json:"province_id"`
}

type CityQueryParams struct {
	request.PageParams
	CityName   string `uri:"city_name"  form:"city_name" json:"city_name"`
	Status     int    `uri:"status" json:"status" form:"status"`
	ProvinceId uint64 `uri:"province_id"  form:"province_id" json:"province_id"`
}

type CityAddForm struct {
	ProvinceId uint64 `form:"province_id" json:"province_id" valid:"Required"`
	Status     int    `form:"status" json:"status" valid:"Range(1,2)"`
	CityName   string `json:"city_name" form:"city_name" valid:"Required"`
}

type CityUpdateForm struct {
	Id         uint64 `form:"id" json:"id" valid:"Required"`
	ProvinceId uint64 `form:"province_id" json:"province_id" valid:"Required"`
	Status     int    `form:"status" json:"status" valid:"Range(1,2)"`
	CityName   string `json:"city_name" form:"city_name" valid:"Required"`
}

type DistrictQueryParams struct {
	PageSize     int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum      int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	DistrictName string `uri:"district_name"  form:"district_name" json:"district_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
	CityId       uint64 `uri:"city_id"  form:"city_id" json:"city_id"`
}

type DistrictAddForm struct {
	CityId       uint64 `form:"city_id" json:"city_id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	DistrictName string `json:"district_name" form:"district_name" valid:"Required"`
}

type DistrictUpdateForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	CityId       uint64 `form:"city_id" json:"city_id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	DistrictName string `json:"district_name" form:"district_name" valid:"Required"`
}

type ProvinceUpdateForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	ProvinceName string `form:"province_name" json:"province_name" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
}

type ProvinceAddForm struct {
	ProvinceName string `form:"province_name" json:"province_name" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
}

type ProvinceQueryParams struct {
	PageSize     int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum      int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	ProvinceName string `uri:"province_name"  form:"province_name" json:"province_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
}
