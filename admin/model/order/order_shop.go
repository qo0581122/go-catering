package order

import "catering/model"

type OrderShop struct {
	Id         uint64     `gorm:"primary_key" json:"id"`
	OrderId    uint64     `json:"order_id"`
	ShopId     uint64     `json:"shop_id"`
	ShopNumber string     `json:"shop_number"`
	PayMoney   uint64     `json:"pay_money"`
	PayTime    model.Time `json:"pay_time"`
	GetCode    string     `json:"get_code"`
	CreatedAt  model.Time `gorm:"column:created_time" json:"created_time"`
	Shop       model.Shop `json:"shop" gorm:"foreignKey:ShopId"`
}

func (OrderShop) TableName() string {
	return "order_shop"
}
