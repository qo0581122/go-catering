package order

import (
	"catering/model"
)

type Order struct {
	Id            uint64         `gorm:"primary_key" json:"id"`
	Uid           uint64         `json:"uid"`
	OrderType     uint32         `json:"order_type"`
	GetType       uint32         `json:"get_type"`
	GetTime       model.Time     `json:"get_time"`
	OrderNumber   string         `json:"order_number"`
	OrderStatus   uint32         `json:"order_status"`
	Integration   uint64         `json:"integration"`
	Remark        string         `json:"remark"`
	ProductCount  uint32         `json:"product_count"`
	CreatedAt     model.Time     `gorm:"column:created_time" json:"created_time"`
	UpdateAt      model.Time     `gorm:"column:updated_time" json:"updated_time"`
	OrderShop     OrderShop      `json:"order_shop" gorm:"foreignKey:order_id"`
	OrderPay      OrderPay       `json:"order_pay" gorm:"foreignKey:order_id"`
	OrderProducts []OrderProduct `json:"order_products" gorm:"foreignKey:order_id"`
}

func (Order) TableName() string {
	return "order"
}
