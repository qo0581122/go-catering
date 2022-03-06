package model

type ProductRelation struct {
	Id              uint64 `gorm:"primary_key" json:"id"`
	ParentProductId uint64 `json:"parent_product_id"`
	ChildProductId  uint64 `json:"child_product_id"`
}
