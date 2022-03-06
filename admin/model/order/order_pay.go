package order

import (
	"catering/model"
)

type OrderPay struct {
	Id            uint64     `gorm:"primary_key" json:"id"`
	OrderId       uint64     `json:"order_id"`
	PayNumber     string     `json:"pay_number"`
	PayMoney      uint64     `json:"pay_money"`
	TotalMoney    uint64     `json:"total_money"`
	DiscountMoney uint64     `json:"discount_money"`
	PayType       uint32     `json:"pay_type"`
	PayTime       model.Time `json:"pay_time"`
}

func (OrderPay) TableName() string {
	return "order_pay"
}
