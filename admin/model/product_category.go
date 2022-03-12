package model

type ProductCategory struct {
	CategoryName string `json:"category_name"`
	Status       int    `json:"status"`
	Model
}

func (ProductCategory) TableName() string {
	return "product_category"
}
