package model

type Province struct {
	ProvinceName string `json:"province_name" xorm:"not null VARCHAR(255)"`
	Status       int    `json:"status" xorm:"not null default 1 INT(1)"`
	Model
}

func (*Province) TableName() string {
	return "province"
}
