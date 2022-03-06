package order

import "catering/model"

type OrderProduct struct {
	Id         uint64     `gorm:"primary_key" json:"id"`
	OrderId    uint64     `json:"order_id"`
	ProductId  uint64     `json:"product_id"`
	PayMoney   uint64     `json:"pay_money"`
	Count      uint32     `json:"count"`
	TotalMoney uint64     `json:"total_money"`
	CreatedAt  model.Time `gorm:"column:created_time" json:"created_time"`
}

func (OrderProduct) TableName() string {
	return "order_product"
}
