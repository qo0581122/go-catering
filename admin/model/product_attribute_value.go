package model

type ProductAttributeValue struct {
	AttributeId    uint64            `json:"attribute_id"`
	AttributeValue string            `json:"attribute_value"`
	Status         int               `json:"status"`
	Attribute      *ProductAttribute `json:"attribute" gorm:"foreignKey:AttributeId"`
	Model
}

func (ProductAttributeValue) TableName() string {
	return "product_attribute_value"
}
