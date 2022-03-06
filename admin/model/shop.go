package model

type Shop struct {
	CategoryId        uint64        `json:"category_id"`
	ShopName          string        `json:"shop_name"`
	ShopAddress       string        `json:"shop_address"`
	Longitude         string        `json:"logintude"`
	Latitude          string        `json:"latitude"`
	BusinessFlag      int           `json:"business_flag"`
	ContactNumber     string        `json:"contact_number"`
	Status            int           `json:"status"`
	DistrictId        uint64        `json:"district_id"`
	ShopDetailAddress string        `json:"shop_detail_address"`
	ProvinceId        uint64        `json:"province_id"`
	CityId            uint64        `json:"city_id"`
	Category          *ShopCategory `gorm:"foreignKey:CategoryId" json:"category"`
	Province          *Province     `gorm:"foreign:ProvinceId" json:"province"`
	City              *City         `gorm:"foreignKey:CityId" json:"city"`
	District          *District     `gorm:"foreignKey:DistrictId" json:"district"`
	Model
}

func (*Shop) TableName() string {
	return "shop"
}
