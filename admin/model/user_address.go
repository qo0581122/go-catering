package model

type UserAddress struct {
	Model
	Uid           uint64           `json:"uid"`
	ReceiveName   string           `json:"receive_name"`
	ReceiveSex    uint32           `json:"receive_sex"`
	ReceivePhone  string           `json:"receive_phone"`
	ProvinceId    uint64           `json:"province_id"`
	CityId        uint64           `json:"city_id"`
	DistinctId    uint64           `json:"distinct_id"`
	DetailAddress string           `json:"detail_address"`
	IsDefault     uint32           `json:"is_default"`
	Sort          uint32           `json:"sort"`
	Province      *Province        `json:"province" gorm:"foreignKey:ProvinceId"`
	City          *City            `json:"city" gorm:"foreignKey:CityId"`
	District      *District        `json:"district" gorm:"foreignKey:DistinctId"`
	AddressTag    []UserAddressTag `json:"tags" gorm:"many2many:user_address_tag_relation;joinForeignKey:address_id;joinReferences:tag_id"`
}

func (UserAddress) TableName() string {
	return "user_address"
}
