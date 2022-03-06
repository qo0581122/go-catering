package model

type ProductAttributeRelation struct {
	Id          uint64 `gorm:"primary_key" json:"id"`
	ProductId   uint64 `json:"product_id"`
	AttributeId uint64 `json:"attribute_id"`
}

func (ProductAttributeRelation) TableName() string {
	return "product_attribute_relation"
}
