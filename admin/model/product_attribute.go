package model

type ProductAttribute struct {
	AttributeName string                  `json:"attribute_name"`
	Status        int                     `json:"status"`
	Values        []ProductAttributeValue `json:"values" gorm:"foreignKey:AttributeId"`
	Model
}

func (ProductAttribute) TableName() string {
	return "product_attribute"
}
