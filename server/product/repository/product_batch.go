package repository

import (
	pb "catering/proto/product"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProductBatchRepository interface {
	SelectByProductId(productId int64) []*pb.ProductBatch
}

type ProductBatchRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *ProductBatchRepositoryImpl) SelectByProductId(productId int64) []*pb.ProductBatch {
	SQL := "SELECT b.id, b.batch_name, b.sort FROM product_batch b left join product_batch_relation br on b.id = br.batch_id WHERE br.product_id = ? order by b.sort desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(productId)
	if err != nil {
		return nil
	}
	var result []*pb.ProductBatch
	var id sql.NullInt64
	var batch_name sql.NullString
	var sort sql.NullInt32
	for rows.Next() {
		rows.Scan(&id, &batch_name, &sort)
		result = append(result, &pb.ProductBatch{
			Id:        id.Int64,
			BatchName: batch_name.String,
			Sort:      sort.Int32,
		})
	}
	return result
}
