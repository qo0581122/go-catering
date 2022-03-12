package model

type Product struct {
	ProductName   string              `json:"product_name"`
	Description   string              `json:"description"`
	CategoryId    uint64              `json:"category_id"`
	Specis        int                 `json:"specis"`
	Url           string              `json:"url"`
	Status        int                 `json:"status"`
	OriginalPrice int                 `json:"original_price"`
	PayPrice      int                 `json:"pay_price"`
	Discount      int                 `json:"discount"`
	Category      *ProductCategory    `gorm:"foreignKey:CategoryId" json:"category"`
	Attributes    []*ProductAttribute `json:"attributes" gorm:"many2many:product_attribute_relation;joinForeignKey:product_id;joinReferences:attribute_id"`
	Batchs        []*ProductBatch     `json:"batchs" gorm:"many2many:product_batch_relation;joinForeignKey:product_id;joinReferences:batch_id"`
	Model
}

func (Product) TableName() string {
	return "product"
}
