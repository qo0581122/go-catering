package model

type ProductBatchRelation struct {
	Id        uint64 `gorm:"primary_key" json:"id"`
	ProductId uint64 `json:"product_id"`
	BatchId   uint64 `json:"batch_id"`
}

func (ProductBatchRelation) TableName() string {
	return "product_batch_relation"
}
