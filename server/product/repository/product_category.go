package repository

import (
	pb "catering/proto/product"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProductCategoryRepository interface {
	SelectByProductId(productId int64) *pb.ProductCategory
	SelectAll() []*pb.ProductCategory
}

type ProductCategoryRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *ProductCategoryRepositoryImpl) SelectByProductId(productId int64) *pb.ProductCategory {
	SQL := "SELECT b.id, b.category_name, b.sort FROM product_category b left join product p on b.id = p.category_id WHERE p.id = ? order by b.sort desc"
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
	var sort sql.NullInt32
	var categoryName sql.NullString
	for rows.Next() {
		rows.Scan(&id, &categoryName, &sort)
		return &pb.ProductCategory{
			Id:           id.Int64,
			CategoryName: categoryName.String,
			Sort:         sort.Int32,
		}
	}
	return nil
}

func (impl *ProductCategoryRepositoryImpl) SelectAll() []*pb.ProductCategory {
	SQL := "SELECT id, category_name, sort FROM product_category WHERE status = 1 ORDER BY sort desc, id desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	var id sql.NullInt64
	var sort sql.NullInt32
	var categoryName sql.NullString
	var result []*pb.ProductCategory
	for rows.Next() {
		rows.Scan(&id, &categoryName, &sort)
		result = append(result, &pb.ProductCategory{
			Id:           id.Int64,
			CategoryName: categoryName.String,
			Sort:         sort.Int32,
		})
	}
	return result
}
