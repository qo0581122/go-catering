package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProductRelationRepository interface {
	SelectIdByParentId(productId int64) []int64
}

type ProductRelationRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *ProductRelationRepositoryImpl) SelectIdByParentId(productId int64) []int64 {
	SQL := "SELECT child_product_id FROM product_relation WHERE parent_product_id = ?"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(productId)
	if err != nil {
		return nil
	}
	var id sql.NullInt64
	var result []int64
	for rows.Next() {
		rows.Scan(&id)
		result = append(result, id.Int64)
	}
	return result
}
