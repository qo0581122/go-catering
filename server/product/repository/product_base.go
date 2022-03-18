package repository

import (
	pb "catering/proto/product"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProductBaseRepository interface {
	SelectAll() []*pb.ProductBase
	SelectByProductIds(product_id int64) *pb.ProductBase
}

type ProductBaseRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *ProductBaseRepositoryImpl) SelectAll() []*pb.ProductBase {
	SQL := "SELECT id, product_name, description, specis, old_money, money, category_id, pic_url, sort FROM product WHERE id = ? AND statu = 1 order by sort desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	var result []*pb.ProductBase
	var sort sql.NullInt32
	var id, specis, oldMoney, money, categoryID sql.NullInt64
	var productName, description, picUrl sql.NullString
	for rows.Next() {
		rows.Scan(&id, &productName, &description, &specis, &oldMoney, &money, &categoryID, &picUrl, &sort)
		result = append(result, &pb.ProductBase{
			Id:          id.Int64,
			ProductName: productName.String,
			Description: description.String,
			Specis:      specis.Int64,
			OldMoney:    oldMoney.Int64,
			Money:       money.Int64,
			CategoryId:  categoryID.Int64,
			PicUrl:      picUrl.String,
			Sort:        sort.Int32,
		})
	}
	return result
}

func (impl *ProductBaseRepositoryImpl) SelectByProductIds(id int64) *pb.ProductBase {
	SQL := "SELECT product_name, description, specis, old_money, money, category_id, pic_url, sort FROM product WHERE id = ? AND statu = 1 limit 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil
	}
	var sort sql.NullInt32
	var specis, oldMoney, money, categoryID sql.NullInt64
	var productName, description, picUrl sql.NullString
	for rows.Next() {
		rows.Scan(&productName, &description, &specis, &oldMoney, &money, &categoryID, &picUrl, &sort)
		return &pb.ProductBase{
			Id:          id,
			ProductName: productName.String,
			Description: description.String,
			Specis:      specis.Int64,
			OldMoney:    oldMoney.Int64,
			Money:       money.Int64,
			CategoryId:  categoryID.Int64,
			PicUrl:      picUrl.String,
			Sort:        sort.Int32,
		}
	}
	return nil
}
