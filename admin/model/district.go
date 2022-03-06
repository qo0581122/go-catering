package model

type District struct {
	CityId       uint64 `json:"city_id"`
	DistrictName string `json:"district_name"`
	Status       int    `json:"status"`
	City         *City  `gorm:"foreignKey:CityId" json:"city"`
	Model
}

func (*District) TableName() string {
	return "district"
}
