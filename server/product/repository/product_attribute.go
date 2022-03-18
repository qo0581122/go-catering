package repository

import (
	pb "catering/proto/product"
	"database/sql"
	"sync"

	"github.com/jmoiron/sqlx"
)

type ProductAttributeRepository interface {
	SelectByProductId(productId int64) []*pb.ProductAttribute
}

type ProductAttributeRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *ProductAttributeRepositoryImpl) SelectByProductId(productId int64) []*pb.ProductAttribute {
	SQL := "SELECT a.id, a.attribute_name, a.sort FROM product_attribute a left join product_attribute_relation ar on a.id = ar.attribute_id WHERE ar.product_id = ? and a.status = 1 order by a.sort desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(productId)
	if err != nil {
		return nil
	}
	var result []*pb.ProductAttribute
	var id sql.NullInt64
	var attribute_name sql.NullString
	var sort sql.NullInt32
	for rows.Next() {
		rows.Scan(&id, &attribute_name, &sort)
		result = append(result, &pb.ProductAttribute{
			Id:            id.Int64,
			AttributeName: attribute_name.String,
			Sort:          sort.Int32,
		})
	}
	var wg sync.WaitGroup
	wg.Add(len(result))
	for index, item := range result {
		go func(index int, item *pb.ProductAttribute) {
			defer wg.Done()
			SQL = "SELECT id, value FROM product_attribute_value WHERE attribute_id = ? and status = 1"
			rows, err := stmt.Query(item.Id)
			if err != nil {
				return
			}
			defer stmt.Close()
			var values []*pb.ProductAttributeValue
			for rows.Next() {
				var id sql.NullInt64
				var value_ sql.NullString
				rows.Scan(&id, &value_)
				values = append(values, &pb.ProductAttributeValue{
					Id:    id.Int64,
					Value: value_.String,
				})
			}
			result[index].AttributeValues = values
		}(index, item)
	}
	wg.Wait()
	return result
}
