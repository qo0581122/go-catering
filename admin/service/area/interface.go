package area

import (
	"catering/model"
	"catering/model/common/response"
)

type cityService interface {
	ListPage(pageNum, pageSize int, params *model.City) *response.ApiResponse
	GetById(id uint64) *model.City
	List(params *model.City) []*model.City
	Count() int
	Add(params *model.City) error
	Update(params *model.City) error
	Delete(id uint64) error
}

type provinceService interface {
	ListPage(pageNum, pageSize int, params *model.Province) *response.ApiResponse
	GetById(id uint64) *model.Province
	List(params *model.Province) []*model.Province
	Count() int
	Add(params *model.Province) error
	Update(params *model.Province) error
	Delete(id uint64) error
}

type districtService interface {
	ListPage(pageNum, pageSize int, params *model.District) *response.ApiResponse
	GetById(id uint64) *model.District
	List(params *model.District) []*model.District
	Count() int
	Add(params *model.District) error
	Update(params *model.District) error
	Delete(id uint64) error
}
