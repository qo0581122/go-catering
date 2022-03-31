package model

type UserAddressTagRelation struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	AddressId uint64
	TagId     uint64
}

func (UserAddressTagRelation) TableName() string {
	return "user_address_tag_relation"
}
