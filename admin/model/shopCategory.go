package model

type ShopCategory struct {
	CategoryName string `json:"category_name"`
	Status       int    `json:"status"`
	Model
}

func (*ShopCategory) TableName() string {
	return "shop_category"
}
