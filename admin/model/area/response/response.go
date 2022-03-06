package response

import "catering/model"

type CityDetail struct {
	*model.City
	Distincts []model.District `gorm:"foreignKey:Id" json:"distincts"`
}

type DistrictDetail struct {
	*model.District
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	ProvinceId   uint64 `json:"province_id"`
}
