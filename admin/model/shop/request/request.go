package request

import "catering/model/common/request"

type ShopCategoryListParams struct {
}

type ShopQueryParams struct {
	request.PageParams
	ShopName     string `uri:"shop_name"  form:"shop_name" json:"shop_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
	CityId       uint64 `uri:"city_id"  form:"city_id" json:"city_id"`
	ProvinceId   uint64 `uri:"province_id"  form:"province_id" json:"province_id"`
	DistrictId   uint64 `uri:"district_id"  form:"district_id" json:"district_id"`
	CategoryId   uint64 `uri:"category_id"  form:"category_id" json:"category_id"`
	BusinessFlag int    `uri:"business_flag"  form:"business_flag" json:"business_flag"`
}

type ShopAddForm struct {
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

type ShopUpdateForm struct {
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

type ShopCategoryQueryParams struct {
	request.PageParams
	CategoryName string `uri:"category_name"  form:"category_name" json:"category_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
}

type ShopCategoryAddForm struct {
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryName string `json:"category_name" form:"category_name" valid:"Required"`
}

type ShopCategoryUpdateForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryName string `json:"category_name" form:"category_name" valid:"Required"`
}

type ShopProductQueryParams struct {
	request.PageParams
	ProductId uint64 `uri:"product_id" json:"product_id" form:"product_id"`
	ShopId    uint64 `uri:"shop_id" json:"shop_id" form:"shop_id"`
	Status    uint32 `form:"status" json:"status" valid:"Range(1,2)"`
}

type ShopProductAddForm struct {
	ProductId uint64 `json:"product_id" form:"product_id"  valid:"Required"`
	ShopId    uint64 `json:"shop_id" form:"shop_id"  valid:"Required"`
	Status    uint32 `form:"status" json:"status" valid:"Range(1,2)"`
}

type ShopProductUpdateForm struct {
	ProductId uint64 `json:"product_id" form:"product_id"  valid:"Required"`
	ShopId    uint64 `json:"shop_id" form:"shop_id"  valid:"Required"`
	Status    uint32 `form:"status" json:"status" valid:"Range(1,2)"`
}
