package model

type ShopProduct struct {
	ID        uint64   `gorm:"primary_key" json:"id"`
	ShopId    uint64   `json:"shop_id"`
	ProductId uint64   `json:"product_id"`
	Status    uint32   `json:"status"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductId"`
}
