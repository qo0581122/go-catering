package model

type City struct {
	ProvinceId uint64    `json:"province_id"`
	CityName   string    `json:"city_name"`
	Status     int       `json:"status"`
	Province   *Province `gorm:"foreignKey:ProvinceId" json:"province"`
	Model
}

func (c *City) TableName() string {
	return "city"
}
